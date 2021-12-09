package arrowdb

import (
	"bytes"
	"io/ioutil"
	"math"
	"testing"

	"github.com/go-kit/kit/log"
	"github.com/google/pprof/profile"
	"github.com/parca-dev/parca/pkg/storage"
	"github.com/parca-dev/parca/pkg/storage/metastore"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/prometheus/pkg/labels"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/trace"
	"golang.org/x/net/context"
)

func TestAppendProfile(t *testing.T) {
	ctx := context.Background()
	logger := log.NewNopLogger()
	registry := prometheus.NewRegistry()
	tracer := trace.NewNoopTracerProvider().Tracer("")

	b, err := ioutil.ReadFile("../storage/testdata/profile1.pb.gz")
	require.NoError(t, err)

	p, err := profile.Parse(bytes.NewBuffer(b))
	require.NoError(t, err)

	ms := metastore.NewBadgerMetastore(logger, registry, tracer, metastore.NewRandomUUIDGenerator())

	fp, err := storage.FlatProfileFromPprof(ctx, logger, ms, p, 0)
	require.NoError(t, err)

	db := NewArrowDB()
	appender, _ := db.Appender(ctx, labels.Labels{
		{
			Name:  "__name__",
			Value: "allocs",
		},
	})

	for i := 0; i < 3; i++ {
		err := appender.AppendFlat(ctx, fp)
		require.NoError(t, err)
	}

	q := db.Querier(context.Background(), math.MinInt64, math.MaxInt64, false)
	_ = q.Select(nil, &labels.Matcher{
		Type:  labels.MatchEqual,
		Name:  "__name__",
		Value: "allocs",
	}) // select all - for now
}