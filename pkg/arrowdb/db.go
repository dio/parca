package arrowdb

import (
	"context"
	"fmt"
	"math"
	"sort"
	"strconv"

	"github.com/apache/arrow/go/arrow"
	"github.com/apache/arrow/go/arrow/array"
	"github.com/apache/arrow/go/arrow/memory"
	"github.com/dgraph-io/sroar"
	"github.com/parca-dev/parca/pkg/storage"
	"github.com/parca-dev/parca/pkg/storage/index"
	"github.com/prometheus/prometheus/pkg/labels"
)

const (
	//tenantCol = iota
	//labelSetIDCol
	stackTraceIDCol = iota
	//timeStampCol
	valueCol
)

const (
	timestampMeta = "ts"
	labelsetMeta  = "ls"
)

var schemaFields = []arrow.Field{
	//{Name: "tenant", Type: arrow.BinaryTypes.String},
	//{Name: "labelsetID", Type: arrow.PrimitiveTypes.Uint64},
	{Name: "stackTraceID", Type: arrow.BinaryTypes.String},
	//{Name: "timestamp", Type: arrow.FixedWidthTypes.Time64us},
	{Name: "value", Type: arrow.PrimitiveTypes.Int64},
}

// DB is an in memory arrow db for profile data
type DB struct {
	*memory.GoAllocator
	*arrow.Schema

	// recordList is the list of records present in the db
	// canonically a record is a profile (TODO?)
	recordList []array.Record

	idx *LabelIndex
}

// NewArrowDB returns a new arrow db
func NewArrowDB() *DB {
	return &DB{
		GoAllocator: memory.NewGoAllocator(),
		Schema:      arrow.NewSchema(schemaFields, nil), // no metadata (TODO)
		idx: &LabelIndex{
			postings: index.NewMemPostings(),
		},
	}
}

func (db *DB) String() string {
	tbl := array.NewTableFromRecords(db.Schema, db.recordList)
	defer tbl.Release()

	tr := array.NewTableReader(tbl, -1)
	defer tr.Release()

	var s string
	for tr.Next() {
		rec := tr.Record()
		for i, col := range rec.Columns() {
			s = fmt.Sprintf("%v%q: %v\n", s, rec.ColumnName(i), col)
		}
	}

	return s
}

// Appender implements the storage.Appender interface
func (db *DB) Appender(ctx context.Context, lset labels.Labels) (storage.Appender, error) {
	db.idx.postings.Add(lset.Hash(), lset) // TODO probably not safe to perform here; as there's no guarantee that anything is ever appended
	return &appender{
		lsetID: lset.Hash(),
		db:     db,
	}, nil
}

type appender struct {
	lsetID uint64
	db     *DB
}

func (a *appender) Append(ctx context.Context, p *storage.Profile) error {
	panic("unimplemented")
}

// AppendFlat implements the Appender interface
func (a *appender) AppendFlat(ctx context.Context, p *storage.FlatProfile) error {
	//tenant := "tenant-placeholder"

	// Create a record builder for the profile
	md := arrow.MetadataFrom(map[string]string{
		"PeriodType":  fmt.Sprintf("%v", p.Meta.PeriodType),
		"SampleType":  fmt.Sprintf("%v", p.Meta.SampleType),
		timestampMeta: fmt.Sprintf("%v", p.Meta.Timestamp),
		"Duration":    fmt.Sprintf("%v", p.Meta.Duration),
		"Period":      fmt.Sprintf("%v", p.Meta.Period),
		labelsetMeta:  fmt.Sprintf("%v", a.lsetID),
	})
	b := array.NewRecordBuilder(a.db, arrow.NewSchema(schemaFields, &md))
	defer b.Release()

	// Iterate over all samples adding them to the record
	for id, s := range p.Samples() {
		//b.Field(tenantCol).(*array.StringBuilder).Append(tenant)
		//b.Field(labelSetIDCol).(*array.Uint64Builder).Append(a.lsetID)
		b.Field(stackTraceIDCol).(*array.StringBuilder).Append(id)
		//b.Field(timeStampCol).(*array.Time64Builder).Append(arrow.Time64(p.Meta.Timestamp))
		b.Field(valueCol).(*array.Int64Builder).Append(s.Value)
	}

	// Create and store the record
	rec := b.NewRecord()
	a.db.recordList = append(a.db.recordList, rec)

	return nil
}

func (db *DB) Querier(ctx context.Context, mint, maxt int64, _ bool) storage.Querier {
	mints, maxts := fmt.Sprintf("%v", mint), fmt.Sprintf("%v", maxt)
	min := sort.Search(len(db.recordList), func(i int) bool {
		ts := db.recordList[i].Schema().Metadata().Values()[db.recordList[i].Schema().Metadata().FindKey(timestampMeta)]
		return ts >= mints
	})
	max := sort.Search(len(db.recordList), func(i int) bool {
		ts := db.recordList[i].Schema().Metadata().Values()[db.recordList[i].Schema().Metadata().FindKey(timestampMeta)]
		return ts >= maxts
	})

	return &querier{
		ctx:    ctx,
		db:     db,
		minIdx: min,
		maxIdx: max,
	}
}

type querier struct {
	ctx    context.Context
	db     *DB
	minIdx int
	maxIdx int
}

func (q *querier) LabelValues(name string, ms ...*labels.Matcher) ([]string, storage.Warnings, error) {
	//TODO implement me
	panic("implement me")
}

func (q *querier) LabelNames(ms ...*labels.Matcher) ([]string, storage.Warnings, error) {
	//TODO implement me
	panic("implement me")
}

// Select will obtain a set of postings from the label index based on the given label matchers.
// Using those postings it will select a set of stack traces from records that match those postings
func (q *querier) Select(hints *storage.SelectHints, ms ...*labels.Matcher) storage.SeriesSet {
	bm, err := storage.PostingsForMatchers(q.db.idx, ms...)
	if err != nil {
		return nil
	}

	var list []array.Record
	switch {
	case bm.Contains(math.MaxUint64): // @matthias does this check even make sense?
		list = q.db.recordList[q.minIdx:q.maxIdx]
	default:
		rl := q.db.recordList[q.minIdx:q.maxIdx]
		list = []array.Record{}
		for _, r := range rl {
			s := r.Schema().Metadata().Values()[r.Schema().Metadata().FindKey(labelsetMeta)]
			v, err := strconv.ParseUint(s, 10, 64)
			if err != nil {
				panic("busted!")
			}

			if bm.Contains(v) {
				list = append(list, r)
			}
		}
	}

	tbl := array.NewTableFromRecords(q.db.Schema, list)
	defer tbl.Release()

	tr := array.NewTableReader(tbl, -1)
	defer tr.Release()

	samples := map[string]*storage.Sample{}
	for tr.Next() {
		rec := tr.Record()

		s := array.NewStringData(rec.Column(0).Data())
		d := array.NewInt64Data(rec.Column(1).Data())
		for i := 0; i < rec.Column(0).Len(); i++ {
			samples[s.Value(i)] = &storage.Sample{
				Value: d.Value(i),
			}
		}
	}

	// printout partially reconstructed samples
	for k, v := range samples {
		fmt.Printf("%v: %v\n", k, v.Value)
	}

	return nil
}

// traceIDFromMatchers finds the set of trace IDs that satisfy the label matchers
func (q *querier) traceIDFromMatchers(ms ...*labels.Matcher) *sroar.Bitmap {
	return nil
}