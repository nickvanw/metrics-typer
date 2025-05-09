// Copyright 2017 Prometheus Team
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Source: https://github.com/prometheus/prometheus/blob/master/prompb/types.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: internal/proto/prometheus-types.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	descriptorpb "google.golang.org/protobuf/types/descriptorpb"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type MetricMetadata_MetricType int32

const (
	MetricMetadata_UNKNOWN        MetricMetadata_MetricType = 0
	MetricMetadata_COUNTER        MetricMetadata_MetricType = 1
	MetricMetadata_GAUGE          MetricMetadata_MetricType = 2
	MetricMetadata_HISTOGRAM      MetricMetadata_MetricType = 3
	MetricMetadata_GAUGEHISTOGRAM MetricMetadata_MetricType = 4
	MetricMetadata_SUMMARY        MetricMetadata_MetricType = 5
	MetricMetadata_INFO           MetricMetadata_MetricType = 6
	MetricMetadata_STATESET       MetricMetadata_MetricType = 7
)

// Enum value maps for MetricMetadata_MetricType.
var (
	MetricMetadata_MetricType_name = map[int32]string{
		0: "UNKNOWN",
		1: "COUNTER",
		2: "GAUGE",
		3: "HISTOGRAM",
		4: "GAUGEHISTOGRAM",
		5: "SUMMARY",
		6: "INFO",
		7: "STATESET",
	}
	MetricMetadata_MetricType_value = map[string]int32{
		"UNKNOWN":        0,
		"COUNTER":        1,
		"GAUGE":          2,
		"HISTOGRAM":      3,
		"GAUGEHISTOGRAM": 4,
		"SUMMARY":        5,
		"INFO":           6,
		"STATESET":       7,
	}
)

func (x MetricMetadata_MetricType) Enum() *MetricMetadata_MetricType {
	p := new(MetricMetadata_MetricType)
	*p = x
	return p
}

func (x MetricMetadata_MetricType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetricMetadata_MetricType) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_proto_prometheus_types_proto_enumTypes[0].Descriptor()
}

func (MetricMetadata_MetricType) Type() protoreflect.EnumType {
	return &file_internal_proto_prometheus_types_proto_enumTypes[0]
}

func (x MetricMetadata_MetricType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetricMetadata_MetricType.Descriptor instead.
func (MetricMetadata_MetricType) EnumDescriptor() ([]byte, []int) {
	return file_internal_proto_prometheus_types_proto_rawDescGZIP(), []int{0, 0}
}

type LabelMatcher_Type int32

const (
	LabelMatcher_EQ  LabelMatcher_Type = 0
	LabelMatcher_NEQ LabelMatcher_Type = 1
	LabelMatcher_RE  LabelMatcher_Type = 2
	LabelMatcher_NRE LabelMatcher_Type = 3
)

// Enum value maps for LabelMatcher_Type.
var (
	LabelMatcher_Type_name = map[int32]string{
		0: "EQ",
		1: "NEQ",
		2: "RE",
		3: "NRE",
	}
	LabelMatcher_Type_value = map[string]int32{
		"EQ":  0,
		"NEQ": 1,
		"RE":  2,
		"NRE": 3,
	}
)

func (x LabelMatcher_Type) Enum() *LabelMatcher_Type {
	p := new(LabelMatcher_Type)
	*p = x
	return p
}

func (x LabelMatcher_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (LabelMatcher_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_proto_prometheus_types_proto_enumTypes[1].Descriptor()
}

func (LabelMatcher_Type) Type() protoreflect.EnumType {
	return &file_internal_proto_prometheus_types_proto_enumTypes[1]
}

func (x LabelMatcher_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use LabelMatcher_Type.Descriptor instead.
func (LabelMatcher_Type) EnumDescriptor() ([]byte, []int) {
	return file_internal_proto_prometheus_types_proto_rawDescGZIP(), []int{5, 0}
}

// We require this to match chunkenc.Encoding.
type Chunk_Encoding int32

const (
	Chunk_UNKNOWN Chunk_Encoding = 0
	Chunk_XOR     Chunk_Encoding = 1
)

// Enum value maps for Chunk_Encoding.
var (
	Chunk_Encoding_name = map[int32]string{
		0: "UNKNOWN",
		1: "XOR",
	}
	Chunk_Encoding_value = map[string]int32{
		"UNKNOWN": 0,
		"XOR":     1,
	}
)

func (x Chunk_Encoding) Enum() *Chunk_Encoding {
	p := new(Chunk_Encoding)
	*p = x
	return p
}

func (x Chunk_Encoding) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Chunk_Encoding) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_proto_prometheus_types_proto_enumTypes[2].Descriptor()
}

func (Chunk_Encoding) Type() protoreflect.EnumType {
	return &file_internal_proto_prometheus_types_proto_enumTypes[2]
}

func (x Chunk_Encoding) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Chunk_Encoding.Descriptor instead.
func (Chunk_Encoding) EnumDescriptor() ([]byte, []int) {
	return file_internal_proto_prometheus_types_proto_rawDescGZIP(), []int{7, 0}
}

type MetricMetadata struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Represents the metric type, these match the set from Prometheus.
	// Refer to pkg/textparse/interface.go for details.
	// https://github.com/prometheus/prometheus/blob/master/pkg/textparse/interface.go
	Type             MetricMetadata_MetricType `protobuf:"varint,1,opt,name=type,proto3,enum=prometheus.MetricMetadata_MetricType" json:"type,omitempty"`
	MetricFamilyName string                    `protobuf:"bytes,2,opt,name=metric_family_name,json=metricFamilyName,proto3" json:"metric_family_name,omitempty"`
	Help             string                    `protobuf:"bytes,4,opt,name=help,proto3" json:"help,omitempty"`
	Unit             string                    `protobuf:"bytes,5,opt,name=unit,proto3" json:"unit,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *MetricMetadata) Reset() {
	*x = MetricMetadata{}
	mi := &file_internal_proto_prometheus_types_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *MetricMetadata) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricMetadata) ProtoMessage() {}

func (x *MetricMetadata) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_prometheus_types_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricMetadata.ProtoReflect.Descriptor instead.
func (*MetricMetadata) Descriptor() ([]byte, []int) {
	return file_internal_proto_prometheus_types_proto_rawDescGZIP(), []int{0}
}

func (x *MetricMetadata) GetType() MetricMetadata_MetricType {
	if x != nil {
		return x.Type
	}
	return MetricMetadata_UNKNOWN
}

func (x *MetricMetadata) GetMetricFamilyName() string {
	if x != nil {
		return x.MetricFamilyName
	}
	return ""
}

func (x *MetricMetadata) GetHelp() string {
	if x != nil {
		return x.Help
	}
	return ""
}

func (x *MetricMetadata) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

type Sample struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Value         float64                `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	Timestamp     int64                  `protobuf:"varint,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Sample) Reset() {
	*x = Sample{}
	mi := &file_internal_proto_prometheus_types_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Sample) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Sample) ProtoMessage() {}

func (x *Sample) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_prometheus_types_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Sample.ProtoReflect.Descriptor instead.
func (*Sample) Descriptor() ([]byte, []int) {
	return file_internal_proto_prometheus_types_proto_rawDescGZIP(), []int{1}
}

func (x *Sample) GetValue() float64 {
	if x != nil {
		return x.Value
	}
	return 0
}

func (x *Sample) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

// TimeSeries represents samples and labels for a single time series.
type TimeSeries struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Labels        []*Label               `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty"`
	Samples       []*Sample              `protobuf:"bytes,2,rep,name=samples,proto3" json:"samples,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TimeSeries) Reset() {
	*x = TimeSeries{}
	mi := &file_internal_proto_prometheus_types_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TimeSeries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimeSeries) ProtoMessage() {}

func (x *TimeSeries) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_prometheus_types_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimeSeries.ProtoReflect.Descriptor instead.
func (*TimeSeries) Descriptor() ([]byte, []int) {
	return file_internal_proto_prometheus_types_proto_rawDescGZIP(), []int{2}
}

func (x *TimeSeries) GetLabels() []*Label {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *TimeSeries) GetSamples() []*Sample {
	if x != nil {
		return x.Samples
	}
	return nil
}

type Label struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Value         string                 `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Label) Reset() {
	*x = Label{}
	mi := &file_internal_proto_prometheus_types_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Label) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Label) ProtoMessage() {}

func (x *Label) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_prometheus_types_proto_msgTypes[3]
	if x != nil {
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
	return file_internal_proto_prometheus_types_proto_rawDescGZIP(), []int{3}
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

type Labels struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Labels        []*Label               `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Labels) Reset() {
	*x = Labels{}
	mi := &file_internal_proto_prometheus_types_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Labels) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Labels) ProtoMessage() {}

func (x *Labels) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_prometheus_types_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Labels.ProtoReflect.Descriptor instead.
func (*Labels) Descriptor() ([]byte, []int) {
	return file_internal_proto_prometheus_types_proto_rawDescGZIP(), []int{4}
}

func (x *Labels) GetLabels() []*Label {
	if x != nil {
		return x.Labels
	}
	return nil
}

// Matcher specifies a rule, which can match or set of labels or not.
type LabelMatcher struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Type          LabelMatcher_Type      `protobuf:"varint,1,opt,name=type,proto3,enum=prometheus.LabelMatcher_Type" json:"type,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Value         string                 `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LabelMatcher) Reset() {
	*x = LabelMatcher{}
	mi := &file_internal_proto_prometheus_types_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LabelMatcher) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LabelMatcher) ProtoMessage() {}

func (x *LabelMatcher) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_prometheus_types_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LabelMatcher.ProtoReflect.Descriptor instead.
func (*LabelMatcher) Descriptor() ([]byte, []int) {
	return file_internal_proto_prometheus_types_proto_rawDescGZIP(), []int{5}
}

func (x *LabelMatcher) GetType() LabelMatcher_Type {
	if x != nil {
		return x.Type
	}
	return LabelMatcher_EQ
}

func (x *LabelMatcher) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *LabelMatcher) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type ReadHints struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StepMs        int64                  `protobuf:"varint,1,opt,name=step_ms,json=stepMs,proto3" json:"step_ms,omitempty"`    // Query step size in milliseconds.
	Func          string                 `protobuf:"bytes,2,opt,name=func,proto3" json:"func,omitempty"`                       // String representation of surrounding function or aggregation.
	StartMs       int64                  `protobuf:"varint,3,opt,name=start_ms,json=startMs,proto3" json:"start_ms,omitempty"` // Start time in milliseconds.
	EndMs         int64                  `protobuf:"varint,4,opt,name=end_ms,json=endMs,proto3" json:"end_ms,omitempty"`       // End time in milliseconds.
	Grouping      []string               `protobuf:"bytes,5,rep,name=grouping,proto3" json:"grouping,omitempty"`               // List of label names used in aggregation.
	By            bool                   `protobuf:"varint,6,opt,name=by,proto3" json:"by,omitempty"`                          // Indicate whether it is without or by.
	RangeMs       int64                  `protobuf:"varint,7,opt,name=range_ms,json=rangeMs,proto3" json:"range_ms,omitempty"` // Range vector selector range in milliseconds.
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReadHints) Reset() {
	*x = ReadHints{}
	mi := &file_internal_proto_prometheus_types_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadHints) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadHints) ProtoMessage() {}

func (x *ReadHints) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_prometheus_types_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadHints.ProtoReflect.Descriptor instead.
func (*ReadHints) Descriptor() ([]byte, []int) {
	return file_internal_proto_prometheus_types_proto_rawDescGZIP(), []int{6}
}

func (x *ReadHints) GetStepMs() int64 {
	if x != nil {
		return x.StepMs
	}
	return 0
}

func (x *ReadHints) GetFunc() string {
	if x != nil {
		return x.Func
	}
	return ""
}

func (x *ReadHints) GetStartMs() int64 {
	if x != nil {
		return x.StartMs
	}
	return 0
}

func (x *ReadHints) GetEndMs() int64 {
	if x != nil {
		return x.EndMs
	}
	return 0
}

func (x *ReadHints) GetGrouping() []string {
	if x != nil {
		return x.Grouping
	}
	return nil
}

func (x *ReadHints) GetBy() bool {
	if x != nil {
		return x.By
	}
	return false
}

func (x *ReadHints) GetRangeMs() int64 {
	if x != nil {
		return x.RangeMs
	}
	return 0
}

// Chunk represents a TSDB chunk.
// Time range [min, max] is inclusive.
type Chunk struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MinTimeMs     int64                  `protobuf:"varint,1,opt,name=min_time_ms,json=minTimeMs,proto3" json:"min_time_ms,omitempty"`
	MaxTimeMs     int64                  `protobuf:"varint,2,opt,name=max_time_ms,json=maxTimeMs,proto3" json:"max_time_ms,omitempty"`
	Type          Chunk_Encoding         `protobuf:"varint,3,opt,name=type,proto3,enum=prometheus.Chunk_Encoding" json:"type,omitempty"`
	Data          []byte                 `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	mi := &file_internal_proto_prometheus_types_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_prometheus_types_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_internal_proto_prometheus_types_proto_rawDescGZIP(), []int{7}
}

func (x *Chunk) GetMinTimeMs() int64 {
	if x != nil {
		return x.MinTimeMs
	}
	return 0
}

func (x *Chunk) GetMaxTimeMs() int64 {
	if x != nil {
		return x.MaxTimeMs
	}
	return 0
}

func (x *Chunk) GetType() Chunk_Encoding {
	if x != nil {
		return x.Type
	}
	return Chunk_UNKNOWN
}

func (x *Chunk) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

// ChunkedSeries represents single, encoded time series.
type ChunkedSeries struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Labels should be sorted.
	Labels []*Label `protobuf:"bytes,1,rep,name=labels,proto3" json:"labels,omitempty"`
	// Chunks will be in start time order and may overlap.
	Chunks        []*Chunk `protobuf:"bytes,2,rep,name=chunks,proto3" json:"chunks,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChunkedSeries) Reset() {
	*x = ChunkedSeries{}
	mi := &file_internal_proto_prometheus_types_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChunkedSeries) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkedSeries) ProtoMessage() {}

func (x *ChunkedSeries) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_prometheus_types_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkedSeries.ProtoReflect.Descriptor instead.
func (*ChunkedSeries) Descriptor() ([]byte, []int) {
	return file_internal_proto_prometheus_types_proto_rawDescGZIP(), []int{8}
}

func (x *ChunkedSeries) GetLabels() []*Label {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *ChunkedSeries) GetChunks() []*Chunk {
	if x != nil {
		return x.Chunks
	}
	return nil
}

var file_internal_proto_prometheus_types_proto_extTypes = []protoimpl.ExtensionInfo{
	{
		ExtendedType:  (*descriptorpb.FieldOptions)(nil),
		ExtensionType: (*bool)(nil),
		Field:         65001,
		Name:          "prometheus.nullable",
		Tag:           "varint,65001,opt,name=nullable",
		Filename:      "internal/proto/prometheus-types.proto",
	},
}

// Extension fields to descriptorpb.FieldOptions.
var (
	// optional bool nullable = 65001;
	E_Nullable = &file_internal_proto_prometheus_types_proto_extTypes[0]
)

var File_internal_proto_prometheus_types_proto protoreflect.FileDescriptor

const file_internal_proto_prometheus_types_proto_rawDesc = "" +
	"\n" +
	"%internal/proto/prometheus-types.proto\x12\n" +
	"prometheus\x1a google/protobuf/descriptor.proto\"\x9c\x02\n" +
	"\x0eMetricMetadata\x129\n" +
	"\x04type\x18\x01 \x01(\x0e2%.prometheus.MetricMetadata.MetricTypeR\x04type\x12,\n" +
	"\x12metric_family_name\x18\x02 \x01(\tR\x10metricFamilyName\x12\x12\n" +
	"\x04help\x18\x04 \x01(\tR\x04help\x12\x12\n" +
	"\x04unit\x18\x05 \x01(\tR\x04unit\"y\n" +
	"\n" +
	"MetricType\x12\v\n" +
	"\aUNKNOWN\x10\x00\x12\v\n" +
	"\aCOUNTER\x10\x01\x12\t\n" +
	"\x05GAUGE\x10\x02\x12\r\n" +
	"\tHISTOGRAM\x10\x03\x12\x12\n" +
	"\x0eGAUGEHISTOGRAM\x10\x04\x12\v\n" +
	"\aSUMMARY\x10\x05\x12\b\n" +
	"\x04INFO\x10\x06\x12\f\n" +
	"\bSTATESET\x10\a\"<\n" +
	"\x06Sample\x12\x14\n" +
	"\x05value\x18\x01 \x01(\x01R\x05value\x12\x1c\n" +
	"\ttimestamp\x18\x02 \x01(\x03R\ttimestamp\"q\n" +
	"\n" +
	"TimeSeries\x12/\n" +
	"\x06labels\x18\x01 \x03(\v2\x11.prometheus.LabelB\x04\xc8\xde\x1f\x00R\x06labels\x122\n" +
	"\asamples\x18\x02 \x03(\v2\x12.prometheus.SampleB\x04\xc8\xde\x1f\x00R\asamples\"1\n" +
	"\x05Label\x12\x12\n" +
	"\x04name\x18\x01 \x01(\tR\x04name\x12\x14\n" +
	"\x05value\x18\x02 \x01(\tR\x05value\"9\n" +
	"\x06Labels\x12/\n" +
	"\x06labels\x18\x01 \x03(\v2\x11.prometheus.LabelB\x04\xc8\xde\x1f\x00R\x06labels\"\x95\x01\n" +
	"\fLabelMatcher\x121\n" +
	"\x04type\x18\x01 \x01(\x0e2\x1d.prometheus.LabelMatcher.TypeR\x04type\x12\x12\n" +
	"\x04name\x18\x02 \x01(\tR\x04name\x12\x14\n" +
	"\x05value\x18\x03 \x01(\tR\x05value\"(\n" +
	"\x04Type\x12\x06\n" +
	"\x02EQ\x10\x00\x12\a\n" +
	"\x03NEQ\x10\x01\x12\x06\n" +
	"\x02RE\x10\x02\x12\a\n" +
	"\x03NRE\x10\x03\"\xb1\x01\n" +
	"\tReadHints\x12\x17\n" +
	"\astep_ms\x18\x01 \x01(\x03R\x06stepMs\x12\x12\n" +
	"\x04func\x18\x02 \x01(\tR\x04func\x12\x19\n" +
	"\bstart_ms\x18\x03 \x01(\x03R\astartMs\x12\x15\n" +
	"\x06end_ms\x18\x04 \x01(\x03R\x05endMs\x12\x1a\n" +
	"\bgrouping\x18\x05 \x03(\tR\bgrouping\x12\x0e\n" +
	"\x02by\x18\x06 \x01(\bR\x02by\x12\x19\n" +
	"\brange_ms\x18\a \x01(\x03R\arangeMs\"\xad\x01\n" +
	"\x05Chunk\x12\x1e\n" +
	"\vmin_time_ms\x18\x01 \x01(\x03R\tminTimeMs\x12\x1e\n" +
	"\vmax_time_ms\x18\x02 \x01(\x03R\tmaxTimeMs\x12.\n" +
	"\x04type\x18\x03 \x01(\x0e2\x1a.prometheus.Chunk.EncodingR\x04type\x12\x12\n" +
	"\x04data\x18\x04 \x01(\fR\x04data\" \n" +
	"\bEncoding\x12\v\n" +
	"\aUNKNOWN\x10\x00\x12\a\n" +
	"\x03XOR\x10\x01\"q\n" +
	"\rChunkedSeries\x12/\n" +
	"\x06labels\x18\x01 \x03(\v2\x11.prometheus.LabelB\x04\xc8\xde\x1f\x00R\x06labels\x12/\n" +
	"\x06chunks\x18\x02 \x03(\v2\x11.prometheus.ChunkB\x04\xc8\xde\x1f\x00R\x06chunks:;\n" +
	"\bnullable\x12\x1d.google.protobuf.FieldOptions\x18\xe9\xfb\x03 \x01(\bR\bnullableB5Z3github.com/nickvanw/metrics-typer/internal/protob\x06proto3"

var (
	file_internal_proto_prometheus_types_proto_rawDescOnce sync.Once
	file_internal_proto_prometheus_types_proto_rawDescData []byte
)

func file_internal_proto_prometheus_types_proto_rawDescGZIP() []byte {
	file_internal_proto_prometheus_types_proto_rawDescOnce.Do(func() {
		file_internal_proto_prometheus_types_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_internal_proto_prometheus_types_proto_rawDesc), len(file_internal_proto_prometheus_types_proto_rawDesc)))
	})
	return file_internal_proto_prometheus_types_proto_rawDescData
}

var file_internal_proto_prometheus_types_proto_enumTypes = make([]protoimpl.EnumInfo, 3)
var file_internal_proto_prometheus_types_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_internal_proto_prometheus_types_proto_goTypes = []any{
	(MetricMetadata_MetricType)(0),    // 0: prometheus.MetricMetadata.MetricType
	(LabelMatcher_Type)(0),            // 1: prometheus.LabelMatcher.Type
	(Chunk_Encoding)(0),               // 2: prometheus.Chunk.Encoding
	(*MetricMetadata)(nil),            // 3: prometheus.MetricMetadata
	(*Sample)(nil),                    // 4: prometheus.Sample
	(*TimeSeries)(nil),                // 5: prometheus.TimeSeries
	(*Label)(nil),                     // 6: prometheus.Label
	(*Labels)(nil),                    // 7: prometheus.Labels
	(*LabelMatcher)(nil),              // 8: prometheus.LabelMatcher
	(*ReadHints)(nil),                 // 9: prometheus.ReadHints
	(*Chunk)(nil),                     // 10: prometheus.Chunk
	(*ChunkedSeries)(nil),             // 11: prometheus.ChunkedSeries
	(*descriptorpb.FieldOptions)(nil), // 12: google.protobuf.FieldOptions
}
var file_internal_proto_prometheus_types_proto_depIdxs = []int32{
	0,  // 0: prometheus.MetricMetadata.type:type_name -> prometheus.MetricMetadata.MetricType
	6,  // 1: prometheus.TimeSeries.labels:type_name -> prometheus.Label
	4,  // 2: prometheus.TimeSeries.samples:type_name -> prometheus.Sample
	6,  // 3: prometheus.Labels.labels:type_name -> prometheus.Label
	1,  // 4: prometheus.LabelMatcher.type:type_name -> prometheus.LabelMatcher.Type
	2,  // 5: prometheus.Chunk.type:type_name -> prometheus.Chunk.Encoding
	6,  // 6: prometheus.ChunkedSeries.labels:type_name -> prometheus.Label
	10, // 7: prometheus.ChunkedSeries.chunks:type_name -> prometheus.Chunk
	12, // 8: prometheus.nullable:extendee -> google.protobuf.FieldOptions
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	8,  // [8:9] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_internal_proto_prometheus_types_proto_init() }
func file_internal_proto_prometheus_types_proto_init() {
	if File_internal_proto_prometheus_types_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internal_proto_prometheus_types_proto_rawDesc), len(file_internal_proto_prometheus_types_proto_rawDesc)),
			NumEnums:      3,
			NumMessages:   9,
			NumExtensions: 1,
			NumServices:   0,
		},
		GoTypes:           file_internal_proto_prometheus_types_proto_goTypes,
		DependencyIndexes: file_internal_proto_prometheus_types_proto_depIdxs,
		EnumInfos:         file_internal_proto_prometheus_types_proto_enumTypes,
		MessageInfos:      file_internal_proto_prometheus_types_proto_msgTypes,
		ExtensionInfos:    file_internal_proto_prometheus_types_proto_extTypes,
	}.Build()
	File_internal_proto_prometheus_types_proto = out.File
	file_internal_proto_prometheus_types_proto_goTypes = nil
	file_internal_proto_prometheus_types_proto_depIdxs = nil
}
