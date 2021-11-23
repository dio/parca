// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: parca/profilestore/v1alpha1/profilestore.proto

package profilestorev1alpha1

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// WriteRawRequest writes a pprof profile for a given tenant
type WriteRawRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// tenant is the given tenant to store the pprof profile under
	Tenant string `protobuf:"bytes,1,opt,name=tenant,proto3" json:"tenant,omitempty"`
	// series is a set raw pprof profiles and accompanying labels
	Series []*RawProfileSeries `protobuf:"bytes,2,rep,name=series,proto3" json:"series,omitempty"`
}

func (x *WriteRawRequest) Reset() {
	*x = WriteRawRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteRawRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRawRequest) ProtoMessage() {}

func (x *WriteRawRequest) ProtoReflect() protoreflect.Message {
	mi := &file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRawRequest.ProtoReflect.Descriptor instead.
func (*WriteRawRequest) Descriptor() ([]byte, []int) {
	return file_parca_profilestore_v1alpha1_profilestore_proto_rawDescGZIP(), []int{0}
}

func (x *WriteRawRequest) GetTenant() string {
	if x != nil {
		return x.Tenant
	}
	return ""
}

func (x *WriteRawRequest) GetSeries() []*RawProfileSeries {
	if x != nil {
		return x.Series
	}
	return nil
}

// WriteRawResponse is the empty response
type WriteRawResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *WriteRawResponse) Reset() {
	*x = WriteRawResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WriteRawResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRawResponse) ProtoMessage() {}

func (x *WriteRawResponse) ProtoReflect() protoreflect.Message {
	mi := &file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRawResponse.ProtoReflect.Descriptor instead.
func (*WriteRawResponse) Descriptor() ([]byte, []int) {
	return file_parca_profilestore_v1alpha1_profilestore_proto_rawDescGZIP(), []int{1}
}

// RawProfileSeries represents the pprof profile and its associated labels
type RawProfileSeries struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// LabelSet is the key value pairs to identify the corresponding profile
	Labels *LabelSet `protobuf:"bytes,1,opt,name=labels,proto3" json:"labels,omitempty"`
	// samples are the set of profile bytes
	Samples []*RawSample `protobuf:"bytes,2,rep,name=samples,proto3" json:"samples,omitempty"`
}

func (x *RawProfileSeries) Reset() {
	*x = RawProfileSeries{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawProfileSeries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawProfileSeries) ProtoMessage() {}

func (x *RawProfileSeries) ProtoReflect() protoreflect.Message {
	mi := &file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawProfileSeries.ProtoReflect.Descriptor instead.
func (*RawProfileSeries) Descriptor() ([]byte, []int) {
	return file_parca_profilestore_v1alpha1_profilestore_proto_rawDescGZIP(), []int{2}
}

func (x *RawProfileSeries) GetLabels() *LabelSet {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *RawProfileSeries) GetSamples() []*RawSample {
	if x != nil {
		return x.Samples
	}
	return nil
}

// Label is a key value pair of identifiers
type Label struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is the label name
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// value is the value for the label name
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *Label) Reset() {
	*x = Label{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Label) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Label) ProtoMessage() {}

func (x *Label) ProtoReflect() protoreflect.Message {
	mi := &file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Label.ProtoReflect.Descriptor instead.
func (*Label) Descriptor() ([]byte, []int) {
	return file_parca_profilestore_v1alpha1_profilestore_proto_rawDescGZIP(), []int{3}
}

func (x *Label) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Label) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

// LabelSet is a group of labels
type LabelSet struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// labels are the grouping of labels
	Labels []*Label `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty"`
}

func (x *LabelSet) Reset() {
	*x = LabelSet{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LabelSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelSet) ProtoMessage() {}

func (x *LabelSet) ProtoReflect() protoreflect.Message {
	mi := &file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelSet.ProtoReflect.Descriptor instead.
func (*LabelSet) Descriptor() ([]byte, []int) {
	return file_parca_profilestore_v1alpha1_profilestore_proto_rawDescGZIP(), []int{4}
}

func (x *LabelSet) GetLabels() []*Label {
	if x != nil {
		return x.Labels
	}
	return nil
}

// RawSample is the set of bytes that correspond to a pprof profile
type RawSample struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// raw_profile is the set of bytes of the pprof profile
	RawProfile []byte `protobuf:"bytes,1,opt,name=raw_profile,json=rawProfile,proto3" json:"raw_profile,omitempty"`
}

func (x *RawSample) Reset() {
	*x = RawSample{}
	if protoimpl.UnsafeEnabled {
		mi := &file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RawSample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RawSample) ProtoMessage() {}

func (x *RawSample) ProtoReflect() protoreflect.Message {
	mi := &file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RawSample.ProtoReflect.Descriptor instead.
func (*RawSample) Descriptor() ([]byte, []int) {
	return file_parca_profilestore_v1alpha1_profilestore_proto_rawDescGZIP(), []int{5}
}

func (x *RawSample) GetRawProfile() []byte {
	if x != nil {
		return x.RawProfile
	}
	return nil
}

var File_parca_profilestore_v1alpha1_profilestore_proto protoreflect.FileDescriptor

var file_parca_profilestore_v1alpha1_profilestore_proto_rawDesc = []byte{
	0x0a, 0x2e, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x1b, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1c, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x70, 0x0a, 0x0f, 0x57,
	0x72, 0x69, 0x74, 0x65, 0x52, 0x61, 0x77, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x12, 0x45, 0x0a, 0x06, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x61, 0x77, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53,
	0x65, 0x72, 0x69, 0x65, 0x73, 0x52, 0x06, 0x73, 0x65, 0x72, 0x69, 0x65, 0x73, 0x22, 0x12, 0x0a,
	0x10, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x61, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x93, 0x01, 0x0a, 0x10, 0x52, 0x61, 0x77, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x53, 0x65, 0x72, 0x69, 0x65, 0x73, 0x12, 0x3d, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x53, 0x65, 0x74, 0x52, 0x06, 0x6c,
	0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x40, 0x0a, 0x07, 0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x61, 0x77, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x52, 0x07,
	0x73, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x73, 0x22, 0x31, 0x0a, 0x05, 0x4c, 0x61, 0x62, 0x65, 0x6c,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x46, 0x0a, 0x08, 0x4c, 0x61,
	0x62, 0x65, 0x6c, 0x53, 0x65, 0x74, 0x12, 0x3a, 0x0a, 0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2e, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x52, 0x06, 0x6c, 0x61, 0x62, 0x65,
	0x6c, 0x73, 0x22, 0x2c, 0x0a, 0x09, 0x52, 0x61, 0x77, 0x53, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x12,
	0x1f, 0x0a, 0x0b, 0x72, 0x61, 0x77, 0x5f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x72, 0x61, 0x77, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65,
	0x32, 0x9e, 0x01, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x86, 0x01, 0x0a, 0x08, 0x57, 0x72, 0x69,
	0x74, 0x65, 0x52, 0x61, 0x77, 0x12, 0x2c, 0x2e, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2e, 0x70, 0x72,
	0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x61, 0x77, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2d, 0x2e, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x66,
	0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x57, 0x72, 0x69, 0x74, 0x65, 0x52, 0x61, 0x77, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x1d, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x17, 0x22, 0x12, 0x2f, 0x70, 0x72, 0x6f,
	0x66, 0x69, 0x6c, 0x65, 0x73, 0x2f, 0x77, 0x72, 0x69, 0x74, 0x65, 0x72, 0x61, 0x77, 0x3a, 0x01,
	0x2a, 0x42, 0x9c, 0x02, 0x0a, 0x1f, 0x63, 0x6f, 0x6d, 0x2e, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2e,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x11, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74,
	0x6f, 0x72, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x58, 0x67, 0x69, 0x74, 0x68,
	0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2d, 0x64, 0x65, 0x76,
	0x2f, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x61, 0x72, 0x63, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b,
	0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x50, 0x58, 0xaa, 0x02, 0x1b, 0x50, 0x61, 0x72,
	0x63, 0x61, 0x2e, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x1b, 0x50, 0x61, 0x72, 0x63, 0x61,
	0x5c, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5c, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x27, 0x50, 0x61, 0x72, 0x63, 0x61, 0x5c, 0x50,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x5c, 0x56, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0xea, 0x02, 0x1d, 0x50, 0x61, 0x72, 0x63, 0x61, 0x3a, 0x3a, 0x50, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_parca_profilestore_v1alpha1_profilestore_proto_rawDescOnce sync.Once
	file_parca_profilestore_v1alpha1_profilestore_proto_rawDescData = file_parca_profilestore_v1alpha1_profilestore_proto_rawDesc
)

func file_parca_profilestore_v1alpha1_profilestore_proto_rawDescGZIP() []byte {
	file_parca_profilestore_v1alpha1_profilestore_proto_rawDescOnce.Do(func() {
		file_parca_profilestore_v1alpha1_profilestore_proto_rawDescData = protoimpl.X.CompressGZIP(file_parca_profilestore_v1alpha1_profilestore_proto_rawDescData)
	})
	return file_parca_profilestore_v1alpha1_profilestore_proto_rawDescData
}

var file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_parca_profilestore_v1alpha1_profilestore_proto_goTypes = []interface{}{
	(*WriteRawRequest)(nil),  // 0: parca.profilestore.v1alpha1.WriteRawRequest
	(*WriteRawResponse)(nil), // 1: parca.profilestore.v1alpha1.WriteRawResponse
	(*RawProfileSeries)(nil), // 2: parca.profilestore.v1alpha1.RawProfileSeries
	(*Label)(nil),            // 3: parca.profilestore.v1alpha1.Label
	(*LabelSet)(nil),         // 4: parca.profilestore.v1alpha1.LabelSet
	(*RawSample)(nil),        // 5: parca.profilestore.v1alpha1.RawSample
}
var file_parca_profilestore_v1alpha1_profilestore_proto_depIdxs = []int32{
	2, // 0: parca.profilestore.v1alpha1.WriteRawRequest.series:type_name -> parca.profilestore.v1alpha1.RawProfileSeries
	4, // 1: parca.profilestore.v1alpha1.RawProfileSeries.labels:type_name -> parca.profilestore.v1alpha1.LabelSet
	5, // 2: parca.profilestore.v1alpha1.RawProfileSeries.samples:type_name -> parca.profilestore.v1alpha1.RawSample
	3, // 3: parca.profilestore.v1alpha1.LabelSet.labels:type_name -> parca.profilestore.v1alpha1.Label
	0, // 4: parca.profilestore.v1alpha1.ProfileStoreService.WriteRaw:input_type -> parca.profilestore.v1alpha1.WriteRawRequest
	1, // 5: parca.profilestore.v1alpha1.ProfileStoreService.WriteRaw:output_type -> parca.profilestore.v1alpha1.WriteRawResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_parca_profilestore_v1alpha1_profilestore_proto_init() }
func file_parca_profilestore_v1alpha1_profilestore_proto_init() {
	if File_parca_profilestore_v1alpha1_profilestore_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteRawRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WriteRawResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RawProfileSeries); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Label); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LabelSet); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RawSample); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_parca_profilestore_v1alpha1_profilestore_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_parca_profilestore_v1alpha1_profilestore_proto_goTypes,
		DependencyIndexes: file_parca_profilestore_v1alpha1_profilestore_proto_depIdxs,
		MessageInfos:      file_parca_profilestore_v1alpha1_profilestore_proto_msgTypes,
	}.Build()
	File_parca_profilestore_v1alpha1_profilestore_proto = out.File
	file_parca_profilestore_v1alpha1_profilestore_proto_rawDesc = nil
	file_parca_profilestore_v1alpha1_profilestore_proto_goTypes = nil
	file_parca_profilestore_v1alpha1_profilestore_proto_depIdxs = nil
}
