// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: metric.proto

package types

import (
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

type MetricType int32

const (
	// NOTHING should not be used.
	MetricType_NOTHING MetricType = 0
	// Metric for p2p network transfer
	MetricType_P2P_NETWORK MetricType = 1
)

// Enum value maps for MetricType.
var (
	MetricType_name = map[int32]string{
		0: "NOTHING",
		1: "P2P_NETWORK",
	}
	MetricType_value = map[string]int32{
		"NOTHING":     0,
		"P2P_NETWORK": 1,
	}
)

func (x MetricType) Enum() *MetricType {
	p := new(MetricType)
	*p = x
	return p
}

func (x MetricType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (MetricType) Descriptor() protoreflect.EnumDescriptor {
	return file_metric_proto_enumTypes[0].Descriptor()
}

func (MetricType) Type() protoreflect.EnumType {
	return &file_metric_proto_enumTypes[0]
}

func (x MetricType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use MetricType.Descriptor instead.
func (MetricType) EnumDescriptor() ([]byte, []int) {
	return file_metric_proto_rawDescGZIP(), []int{0}
}

type MetricsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Types []MetricType `protobuf:"varint,1,rep,packed,name=types,proto3,enum=types.MetricType" json:"types,omitempty"`
}

func (x *MetricsRequest) Reset() {
	*x = MetricsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metric_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricsRequest) ProtoMessage() {}

func (x *MetricsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_metric_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricsRequest.ProtoReflect.Descriptor instead.
func (*MetricsRequest) Descriptor() ([]byte, []int) {
	return file_metric_proto_rawDescGZIP(), []int{0}
}

func (x *MetricsRequest) GetTypes() []MetricType {
	if x != nil {
		return x.Types
	}
	return nil
}

type Metrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Peers []*PeerMetric `protobuf:"bytes,1,rep,name=peers,proto3" json:"peers,omitempty"`
}

func (x *Metrics) Reset() {
	*x = Metrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metric_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Metrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Metrics) ProtoMessage() {}

func (x *Metrics) ProtoReflect() protoreflect.Message {
	mi := &file_metric_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Metrics.ProtoReflect.Descriptor instead.
func (*Metrics) Descriptor() ([]byte, []int) {
	return file_metric_proto_rawDescGZIP(), []int{1}
}

func (x *Metrics) GetPeers() []*PeerMetric {
	if x != nil {
		return x.Peers
	}
	return nil
}

type PeerMetric struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PeerID []byte `protobuf:"bytes,1,opt,name=peerID,proto3" json:"peerID,omitempty"`
	SumIn  int64  `protobuf:"varint,2,opt,name=sumIn,proto3" json:"sumIn,omitempty"`
	AvrIn  int64  `protobuf:"varint,3,opt,name=avrIn,proto3" json:"avrIn,omitempty"`
	SumOut int64  `protobuf:"varint,4,opt,name=sumOut,proto3" json:"sumOut,omitempty"`
	AvrOut int64  `protobuf:"varint,5,opt,name=avrOut,proto3" json:"avrOut,omitempty"`
}

func (x *PeerMetric) Reset() {
	*x = PeerMetric{}
	if protoimpl.UnsafeEnabled {
		mi := &file_metric_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerMetric) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerMetric) ProtoMessage() {}

func (x *PeerMetric) ProtoReflect() protoreflect.Message {
	mi := &file_metric_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerMetric.ProtoReflect.Descriptor instead.
func (*PeerMetric) Descriptor() ([]byte, []int) {
	return file_metric_proto_rawDescGZIP(), []int{2}
}

func (x *PeerMetric) GetPeerID() []byte {
	if x != nil {
		return x.PeerID
	}
	return nil
}

func (x *PeerMetric) GetSumIn() int64 {
	if x != nil {
		return x.SumIn
	}
	return 0
}

func (x *PeerMetric) GetAvrIn() int64 {
	if x != nil {
		return x.AvrIn
	}
	return 0
}

func (x *PeerMetric) GetSumOut() int64 {
	if x != nil {
		return x.SumOut
	}
	return 0
}

func (x *PeerMetric) GetAvrOut() int64 {
	if x != nil {
		return x.AvrOut
	}
	return 0
}

var File_metric_proto protoreflect.FileDescriptor

var file_metric_proto_rawDesc = []byte{
	0x0a, 0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05,
	0x74, 0x79, 0x70, 0x65, 0x73, 0x22, 0x39, 0x0a, 0x0e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x27, 0x0a, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0e, 0x32, 0x11, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x54, 0x79, 0x70, 0x65, 0x52, 0x05, 0x74, 0x79, 0x70, 0x65, 0x73,
	0x22, 0x32, 0x0a, 0x07, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x27, 0x0a, 0x05, 0x70,
	0x65, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x74, 0x79, 0x70,
	0x65, 0x73, 0x2e, 0x50, 0x65, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x52, 0x05, 0x70,
	0x65, 0x65, 0x72, 0x73, 0x22, 0x80, 0x01, 0x0a, 0x0a, 0x50, 0x65, 0x65, 0x72, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x73,
	0x75, 0x6d, 0x49, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x75, 0x6d, 0x49,
	0x6e, 0x12, 0x14, 0x0a, 0x05, 0x61, 0x76, 0x72, 0x49, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x61, 0x76, 0x72, 0x49, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x75, 0x6d, 0x4f, 0x75,
	0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x73, 0x75, 0x6d, 0x4f, 0x75, 0x74, 0x12,
	0x16, 0x0a, 0x06, 0x61, 0x76, 0x72, 0x4f, 0x75, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x06, 0x61, 0x76, 0x72, 0x4f, 0x75, 0x74, 0x2a, 0x2a, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x4e, 0x4f, 0x54, 0x48, 0x49, 0x4e, 0x47,
	0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x50, 0x32, 0x50, 0x5f, 0x4e, 0x45, 0x54, 0x57, 0x4f, 0x52,
	0x4b, 0x10, 0x01, 0x42, 0x08, 0x5a, 0x06, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_metric_proto_rawDescOnce sync.Once
	file_metric_proto_rawDescData = file_metric_proto_rawDesc
)

func file_metric_proto_rawDescGZIP() []byte {
	file_metric_proto_rawDescOnce.Do(func() {
		file_metric_proto_rawDescData = protoimpl.X.CompressGZIP(file_metric_proto_rawDescData)
	})
	return file_metric_proto_rawDescData
}

var file_metric_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_metric_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_metric_proto_goTypes = []interface{}{
	(MetricType)(0),        // 0: types.MetricType
	(*MetricsRequest)(nil), // 1: types.MetricsRequest
	(*Metrics)(nil),        // 2: types.Metrics
	(*PeerMetric)(nil),     // 3: types.PeerMetric
}
var file_metric_proto_depIdxs = []int32{
	0, // 0: types.MetricsRequest.types:type_name -> types.MetricType
	3, // 1: types.Metrics.peers:type_name -> types.PeerMetric
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_metric_proto_init() }
func file_metric_proto_init() {
	if File_metric_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_metric_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricsRequest); i {
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
		file_metric_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Metrics); i {
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
		file_metric_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerMetric); i {
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
			RawDescriptor: file_metric_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_metric_proto_goTypes,
		DependencyIndexes: file_metric_proto_depIdxs,
		EnumInfos:         file_metric_proto_enumTypes,
		MessageInfos:      file_metric_proto_msgTypes,
	}.Build()
	File_metric_proto = out.File
	file_metric_proto_rawDesc = nil
	file_metric_proto_goTypes = nil
	file_metric_proto_depIdxs = nil
}