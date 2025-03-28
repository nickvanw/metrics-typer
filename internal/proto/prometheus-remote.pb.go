// Copyright 2016 Prometheus Team
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

// Source: https://github.com/prometheus/prometheus/blob/master/prompb/remote.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: internal/proto/prometheus-remote.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
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

type ReadRequest_ResponseType int32

const (
	// Server will return a single ReadResponse message with matched series that includes list of raw samples.
	// It's recommended to use streamed response types instead.
	//
	// Response headers:
	// Content-Type: "application/x-protobuf"
	// Content-Encoding: "snappy"
	ReadRequest_SAMPLES ReadRequest_ResponseType = 0
	// Server will stream a delimited ChunkedReadResponse message that contains XOR encoded chunks for a single series.
	// Each message is following varint size and fixed size bigendian uint32 for CRC32 Castagnoli checksum.
	//
	// Response headers:
	// Content-Type: "application/x-streamed-protobuf; proto=prometheus.ChunkedReadResponse"
	// Content-Encoding: ""
	ReadRequest_STREAMED_XOR_CHUNKS ReadRequest_ResponseType = 1
)

// Enum value maps for ReadRequest_ResponseType.
var (
	ReadRequest_ResponseType_name = map[int32]string{
		0: "SAMPLES",
		1: "STREAMED_XOR_CHUNKS",
	}
	ReadRequest_ResponseType_value = map[string]int32{
		"SAMPLES":             0,
		"STREAMED_XOR_CHUNKS": 1,
	}
)

func (x ReadRequest_ResponseType) Enum() *ReadRequest_ResponseType {
	p := new(ReadRequest_ResponseType)
	*p = x
	return p
}

func (x ReadRequest_ResponseType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ReadRequest_ResponseType) Descriptor() protoreflect.EnumDescriptor {
	return file_internal_proto_prometheus_remote_proto_enumTypes[0].Descriptor()
}

func (ReadRequest_ResponseType) Type() protoreflect.EnumType {
	return &file_internal_proto_prometheus_remote_proto_enumTypes[0]
}

func (x ReadRequest_ResponseType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ReadRequest_ResponseType.Descriptor instead.
func (ReadRequest_ResponseType) EnumDescriptor() ([]byte, []int) {
	return file_internal_proto_prometheus_remote_proto_rawDescGZIP(), []int{1, 0}
}

type WriteRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Timeseries    []*TimeSeries          `protobuf:"bytes,1,rep,name=timeseries,proto3" json:"timeseries,omitempty"`
	Metadata      []*MetricMetadata      `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WriteRequest) Reset() {
	*x = WriteRequest{}
	mi := &file_internal_proto_prometheus_remote_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WriteRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WriteRequest) ProtoMessage() {}

func (x *WriteRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_prometheus_remote_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WriteRequest.ProtoReflect.Descriptor instead.
func (*WriteRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_prometheus_remote_proto_rawDescGZIP(), []int{0}
}

func (x *WriteRequest) GetTimeseries() []*TimeSeries {
	if x != nil {
		return x.Timeseries
	}
	return nil
}

func (x *WriteRequest) GetMetadata() []*MetricMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

// ReadRequest represents a remote read request.
type ReadRequest struct {
	state   protoimpl.MessageState `protogen:"open.v1"`
	Queries []*Query               `protobuf:"bytes,1,rep,name=queries,proto3" json:"queries,omitempty"`
	// accepted_response_types allows negotiating the content type of the response.
	//
	// Response types are taken from the list in the FIFO order. If no response type in `accepted_response_types` is
	// implemented by server, error is returned.
	// For request that do not contain `accepted_response_types` field the SAMPLES response type will be used.
	AcceptedResponseTypes []ReadRequest_ResponseType `protobuf:"varint,2,rep,packed,name=accepted_response_types,json=acceptedResponseTypes,proto3,enum=prometheus.ReadRequest_ResponseType" json:"accepted_response_types,omitempty"`
	unknownFields         protoimpl.UnknownFields
	sizeCache             protoimpl.SizeCache
}

func (x *ReadRequest) Reset() {
	*x = ReadRequest{}
	mi := &file_internal_proto_prometheus_remote_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadRequest) ProtoMessage() {}

func (x *ReadRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_prometheus_remote_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadRequest.ProtoReflect.Descriptor instead.
func (*ReadRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_prometheus_remote_proto_rawDescGZIP(), []int{1}
}

func (x *ReadRequest) GetQueries() []*Query {
	if x != nil {
		return x.Queries
	}
	return nil
}

func (x *ReadRequest) GetAcceptedResponseTypes() []ReadRequest_ResponseType {
	if x != nil {
		return x.AcceptedResponseTypes
	}
	return nil
}

// ReadResponse is a response when response_type equals SAMPLES.
type ReadResponse struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// In same order as the request's queries.
	Results       []*QueryResult `protobuf:"bytes,1,rep,name=results,proto3" json:"results,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReadResponse) Reset() {
	*x = ReadResponse{}
	mi := &file_internal_proto_prometheus_remote_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadResponse) ProtoMessage() {}

func (x *ReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_prometheus_remote_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadResponse.ProtoReflect.Descriptor instead.
func (*ReadResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_prometheus_remote_proto_rawDescGZIP(), []int{2}
}

func (x *ReadResponse) GetResults() []*QueryResult {
	if x != nil {
		return x.Results
	}
	return nil
}

type Query struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	StartTimestampMs int64                  `protobuf:"varint,1,opt,name=start_timestamp_ms,json=startTimestampMs,proto3" json:"start_timestamp_ms,omitempty"`
	EndTimestampMs   int64                  `protobuf:"varint,2,opt,name=end_timestamp_ms,json=endTimestampMs,proto3" json:"end_timestamp_ms,omitempty"`
	Matchers         []*LabelMatcher        `protobuf:"bytes,3,rep,name=matchers,proto3" json:"matchers,omitempty"`
	Hints            *ReadHints             `protobuf:"bytes,4,opt,name=hints,proto3" json:"hints,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *Query) Reset() {
	*x = Query{}
	mi := &file_internal_proto_prometheus_remote_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Query) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Query) ProtoMessage() {}

func (x *Query) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_prometheus_remote_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Query.ProtoReflect.Descriptor instead.
func (*Query) Descriptor() ([]byte, []int) {
	return file_internal_proto_prometheus_remote_proto_rawDescGZIP(), []int{3}
}

func (x *Query) GetStartTimestampMs() int64 {
	if x != nil {
		return x.StartTimestampMs
	}
	return 0
}

func (x *Query) GetEndTimestampMs() int64 {
	if x != nil {
		return x.EndTimestampMs
	}
	return 0
}

func (x *Query) GetMatchers() []*LabelMatcher {
	if x != nil {
		return x.Matchers
	}
	return nil
}

func (x *Query) GetHints() *ReadHints {
	if x != nil {
		return x.Hints
	}
	return nil
}

type QueryResult struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// Samples within a time series must be ordered by time.
	Timeseries    []*TimeSeries `protobuf:"bytes,1,rep,name=timeseries,proto3" json:"timeseries,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *QueryResult) Reset() {
	*x = QueryResult{}
	mi := &file_internal_proto_prometheus_remote_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *QueryResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*QueryResult) ProtoMessage() {}

func (x *QueryResult) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_prometheus_remote_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use QueryResult.ProtoReflect.Descriptor instead.
func (*QueryResult) Descriptor() ([]byte, []int) {
	return file_internal_proto_prometheus_remote_proto_rawDescGZIP(), []int{4}
}

func (x *QueryResult) GetTimeseries() []*TimeSeries {
	if x != nil {
		return x.Timeseries
	}
	return nil
}

// ChunkedReadResponse is a response when response_type equals STREAMED_XOR_CHUNKS.
// We strictly stream full series after series, optionally split by time. This means that a single frame can contain
// partition of the single series, but once a new series is started to be streamed it means that no more chunks will
// be sent for previous one. Series are returned sorted in the same way TSDB block are internally.
type ChunkedReadResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ChunkedSeries []*ChunkedSeries       `protobuf:"bytes,1,rep,name=chunked_series,json=chunkedSeries,proto3" json:"chunked_series,omitempty"`
	// query_index represents an index of the query from ReadRequest.queries these chunks relates to.
	QueryIndex    int64 `protobuf:"varint,2,opt,name=query_index,json=queryIndex,proto3" json:"query_index,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ChunkedReadResponse) Reset() {
	*x = ChunkedReadResponse{}
	mi := &file_internal_proto_prometheus_remote_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ChunkedReadResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChunkedReadResponse) ProtoMessage() {}

func (x *ChunkedReadResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_prometheus_remote_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChunkedReadResponse.ProtoReflect.Descriptor instead.
func (*ChunkedReadResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_prometheus_remote_proto_rawDescGZIP(), []int{5}
}

func (x *ChunkedReadResponse) GetChunkedSeries() []*ChunkedSeries {
	if x != nil {
		return x.ChunkedSeries
	}
	return nil
}

func (x *ChunkedReadResponse) GetQueryIndex() int64 {
	if x != nil {
		return x.QueryIndex
	}
	return 0
}

var File_internal_proto_prometheus_remote_proto protoreflect.FileDescriptor

const file_internal_proto_prometheus_remote_proto_rawDesc = "" +
	"\n" +
	"&internal/proto/prometheus-remote.proto\x12\n" +
	"prometheus\x1a%internal/proto/prometheus-types.proto\"\x90\x01\n" +
	"\fWriteRequest\x12<\n" +
	"\n" +
	"timeseries\x18\x01 \x03(\v2\x16.prometheus.TimeSeriesB\x04\xc8\xde\x1f\x00R\n" +
	"timeseries\x12<\n" +
	"\bmetadata\x18\x03 \x03(\v2\x1a.prometheus.MetricMetadataB\x04\xc8\xde\x1f\x00R\bmetadataJ\x04\b\x02\x10\x03\"\xce\x01\n" +
	"\vReadRequest\x12+\n" +
	"\aqueries\x18\x01 \x03(\v2\x11.prometheus.QueryR\aqueries\x12\\\n" +
	"\x17accepted_response_types\x18\x02 \x03(\x0e2$.prometheus.ReadRequest.ResponseTypeR\x15acceptedResponseTypes\"4\n" +
	"\fResponseType\x12\v\n" +
	"\aSAMPLES\x10\x00\x12\x17\n" +
	"\x13STREAMED_XOR_CHUNKS\x10\x01\"A\n" +
	"\fReadResponse\x121\n" +
	"\aresults\x18\x01 \x03(\v2\x17.prometheus.QueryResultR\aresults\"\xc2\x01\n" +
	"\x05Query\x12,\n" +
	"\x12start_timestamp_ms\x18\x01 \x01(\x03R\x10startTimestampMs\x12(\n" +
	"\x10end_timestamp_ms\x18\x02 \x01(\x03R\x0eendTimestampMs\x124\n" +
	"\bmatchers\x18\x03 \x03(\v2\x18.prometheus.LabelMatcherR\bmatchers\x12+\n" +
	"\x05hints\x18\x04 \x01(\v2\x15.prometheus.ReadHintsR\x05hints\"E\n" +
	"\vQueryResult\x126\n" +
	"\n" +
	"timeseries\x18\x01 \x03(\v2\x16.prometheus.TimeSeriesR\n" +
	"timeseries\"x\n" +
	"\x13ChunkedReadResponse\x12@\n" +
	"\x0echunked_series\x18\x01 \x03(\v2\x19.prometheus.ChunkedSeriesR\rchunkedSeries\x12\x1f\n" +
	"\vquery_index\x18\x02 \x01(\x03R\n" +
	"queryIndexB5Z3github.com/nickvanw/metrics-typer/internal/protob\x06proto3"

var (
	file_internal_proto_prometheus_remote_proto_rawDescOnce sync.Once
	file_internal_proto_prometheus_remote_proto_rawDescData []byte
)

func file_internal_proto_prometheus_remote_proto_rawDescGZIP() []byte {
	file_internal_proto_prometheus_remote_proto_rawDescOnce.Do(func() {
		file_internal_proto_prometheus_remote_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_internal_proto_prometheus_remote_proto_rawDesc), len(file_internal_proto_prometheus_remote_proto_rawDesc)))
	})
	return file_internal_proto_prometheus_remote_proto_rawDescData
}

var file_internal_proto_prometheus_remote_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_internal_proto_prometheus_remote_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_internal_proto_prometheus_remote_proto_goTypes = []any{
	(ReadRequest_ResponseType)(0), // 0: prometheus.ReadRequest.ResponseType
	(*WriteRequest)(nil),          // 1: prometheus.WriteRequest
	(*ReadRequest)(nil),           // 2: prometheus.ReadRequest
	(*ReadResponse)(nil),          // 3: prometheus.ReadResponse
	(*Query)(nil),                 // 4: prometheus.Query
	(*QueryResult)(nil),           // 5: prometheus.QueryResult
	(*ChunkedReadResponse)(nil),   // 6: prometheus.ChunkedReadResponse
	(*TimeSeries)(nil),            // 7: prometheus.TimeSeries
	(*MetricMetadata)(nil),        // 8: prometheus.MetricMetadata
	(*LabelMatcher)(nil),          // 9: prometheus.LabelMatcher
	(*ReadHints)(nil),             // 10: prometheus.ReadHints
	(*ChunkedSeries)(nil),         // 11: prometheus.ChunkedSeries
}
var file_internal_proto_prometheus_remote_proto_depIdxs = []int32{
	7,  // 0: prometheus.WriteRequest.timeseries:type_name -> prometheus.TimeSeries
	8,  // 1: prometheus.WriteRequest.metadata:type_name -> prometheus.MetricMetadata
	4,  // 2: prometheus.ReadRequest.queries:type_name -> prometheus.Query
	0,  // 3: prometheus.ReadRequest.accepted_response_types:type_name -> prometheus.ReadRequest.ResponseType
	5,  // 4: prometheus.ReadResponse.results:type_name -> prometheus.QueryResult
	9,  // 5: prometheus.Query.matchers:type_name -> prometheus.LabelMatcher
	10, // 6: prometheus.Query.hints:type_name -> prometheus.ReadHints
	7,  // 7: prometheus.QueryResult.timeseries:type_name -> prometheus.TimeSeries
	11, // 8: prometheus.ChunkedReadResponse.chunked_series:type_name -> prometheus.ChunkedSeries
	9,  // [9:9] is the sub-list for method output_type
	9,  // [9:9] is the sub-list for method input_type
	9,  // [9:9] is the sub-list for extension type_name
	9,  // [9:9] is the sub-list for extension extendee
	0,  // [0:9] is the sub-list for field type_name
}

func init() { file_internal_proto_prometheus_remote_proto_init() }
func file_internal_proto_prometheus_remote_proto_init() {
	if File_internal_proto_prometheus_remote_proto != nil {
		return
	}
	file_internal_proto_prometheus_types_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_internal_proto_prometheus_remote_proto_rawDesc), len(file_internal_proto_prometheus_remote_proto_rawDesc)),
			NumEnums:      1,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_internal_proto_prometheus_remote_proto_goTypes,
		DependencyIndexes: file_internal_proto_prometheus_remote_proto_depIdxs,
		EnumInfos:         file_internal_proto_prometheus_remote_proto_enumTypes,
		MessageInfos:      file_internal_proto_prometheus_remote_proto_msgTypes,
	}.Build()
	File_internal_proto_prometheus_remote_proto = out.File
	file_internal_proto_prometheus_remote_proto_goTypes = nil
	file_internal_proto_prometheus_remote_proto_depIdxs = nil
}
