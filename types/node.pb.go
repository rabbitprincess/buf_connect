// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: node.proto

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

type PeerRole int32

const (
	PeerRole_LegacyVersion PeerRole = 0
	PeerRole_Producer      PeerRole = 1
	PeerRole_Watcher       PeerRole = 2
	PeerRole_Agent         PeerRole = 3
)

// Enum value maps for PeerRole.
var (
	PeerRole_name = map[int32]string{
		0: "LegacyVersion",
		1: "Producer",
		2: "Watcher",
		3: "Agent",
	}
	PeerRole_value = map[string]int32{
		"LegacyVersion": 0,
		"Producer":      1,
		"Watcher":       2,
		"Agent":         3,
	}
)

func (x PeerRole) Enum() *PeerRole {
	p := new(PeerRole)
	*p = x
	return p
}

func (x PeerRole) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (PeerRole) Descriptor() protoreflect.EnumDescriptor {
	return file_node_proto_enumTypes[0].Descriptor()
}

func (PeerRole) Type() protoreflect.EnumType {
	return &file_node_proto_enumTypes[0]
}

func (x PeerRole) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use PeerRole.Descriptor instead.
func (PeerRole) EnumDescriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{0}
}

// PeerAddress contains static information of peer and addresses to connect peer
type PeerAddress struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// @Deprecated advertised address and port will be in addresses field in aergo v2.
	// address is string representation of ip address or domain name.
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	// @Deprecated
	Port        uint32   `protobuf:"varint,2,opt,name=port,proto3" json:"port,omitempty"`
	PeerID      []byte   `protobuf:"bytes,3,opt,name=peerID,proto3" json:"peerID,omitempty"`
	Role        PeerRole `protobuf:"varint,4,opt,name=role,proto3,enum=types.PeerRole" json:"role,omitempty"`
	Version     string   `protobuf:"bytes,5,opt,name=version,proto3" json:"version,omitempty"`
	Addresses   []string `protobuf:"bytes,6,rep,name=addresses,proto3" json:"addresses,omitempty"`
	ProducerIDs [][]byte `protobuf:"bytes,7,rep,name=producerIDs,proto3" json:"producerIDs,omitempty"`
}

func (x *PeerAddress) Reset() {
	*x = PeerAddress{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PeerAddress) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PeerAddress) ProtoMessage() {}

func (x *PeerAddress) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PeerAddress.ProtoReflect.Descriptor instead.
func (*PeerAddress) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{0}
}

func (x *PeerAddress) GetAddress() string {
	if x != nil {
		return x.Address
	}
	return ""
}

func (x *PeerAddress) GetPort() uint32 {
	if x != nil {
		return x.Port
	}
	return 0
}

func (x *PeerAddress) GetPeerID() []byte {
	if x != nil {
		return x.PeerID
	}
	return nil
}

func (x *PeerAddress) GetRole() PeerRole {
	if x != nil {
		return x.Role
	}
	return PeerRole_LegacyVersion
}

func (x *PeerAddress) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *PeerAddress) GetAddresses() []string {
	if x != nil {
		return x.Addresses
	}
	return nil
}

func (x *PeerAddress) GetProducerIDs() [][]byte {
	if x != nil {
		return x.ProducerIDs
	}
	return nil
}

type AgentCertificate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CertVersion uint32 `protobuf:"varint,1,opt,name=certVersion,proto3" json:"certVersion,omitempty"`
	BPID        []byte `protobuf:"bytes,2,opt,name=BPID,proto3" json:"BPID,omitempty"`
	BPPubKey    []byte `protobuf:"bytes,3,opt,name=BPPubKey,proto3" json:"BPPubKey,omitempty"`
	// CreateTime is the number of nanoseconds elapsed since January 1, 1970 UTC
	CreateTime int64 `protobuf:"varint,4,opt,name=createTime,proto3" json:"createTime,omitempty"`
	// CreateTime is the number of nanoseconds elapsed since January 1, 1970 UTC
	ExpireTime   int64    `protobuf:"varint,5,opt,name=expireTime,proto3" json:"expireTime,omitempty"`
	AgentID      []byte   `protobuf:"bytes,6,opt,name=agentID,proto3" json:"agentID,omitempty"`
	AgentAddress [][]byte `protobuf:"bytes,7,rep,name=AgentAddress,proto3" json:"AgentAddress,omitempty"`
	Signature    []byte   `protobuf:"bytes,8,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *AgentCertificate) Reset() {
	*x = AgentCertificate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_node_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AgentCertificate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AgentCertificate) ProtoMessage() {}

func (x *AgentCertificate) ProtoReflect() protoreflect.Message {
	mi := &file_node_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AgentCertificate.ProtoReflect.Descriptor instead.
func (*AgentCertificate) Descriptor() ([]byte, []int) {
	return file_node_proto_rawDescGZIP(), []int{1}
}

func (x *AgentCertificate) GetCertVersion() uint32 {
	if x != nil {
		return x.CertVersion
	}
	return 0
}

func (x *AgentCertificate) GetBPID() []byte {
	if x != nil {
		return x.BPID
	}
	return nil
}

func (x *AgentCertificate) GetBPPubKey() []byte {
	if x != nil {
		return x.BPPubKey
	}
	return nil
}

func (x *AgentCertificate) GetCreateTime() int64 {
	if x != nil {
		return x.CreateTime
	}
	return 0
}

func (x *AgentCertificate) GetExpireTime() int64 {
	if x != nil {
		return x.ExpireTime
	}
	return 0
}

func (x *AgentCertificate) GetAgentID() []byte {
	if x != nil {
		return x.AgentID
	}
	return nil
}

func (x *AgentCertificate) GetAgentAddress() [][]byte {
	if x != nil {
		return x.AgentAddress
	}
	return nil
}

func (x *AgentCertificate) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

var File_node_proto protoreflect.FileDescriptor

var file_node_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x6e, 0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x74, 0x79,
	0x70, 0x65, 0x73, 0x22, 0xd2, 0x01, 0x0a, 0x0b, 0x50, 0x65, 0x65, 0x72, 0x41, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x12, 0x0a,
	0x04, 0x70, 0x6f, 0x72, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x04, 0x70, 0x6f, 0x72,
	0x74, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x70, 0x65, 0x65, 0x72, 0x49, 0x44, 0x12, 0x23, 0x0a, 0x04, 0x72, 0x6f, 0x6c,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x0f, 0x2e, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e,
	0x50, 0x65, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x52, 0x04, 0x72, 0x6f, 0x6c, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x0a, 0x09, 0x61, 0x64, 0x64, 0x72,
	0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x65, 0x72, 0x49, 0x44, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x0b, 0x70, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x65, 0x72, 0x49, 0x44, 0x73, 0x22, 0x80, 0x02, 0x0a, 0x10, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x63, 0x65, 0x72, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0d, 0x52, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12,
	0x12, 0x0a, 0x04, 0x42, 0x50, 0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x42,
	0x50, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x42, 0x50, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x42, 0x50, 0x50, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12,
	0x1e, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x1e, 0x0a, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0a, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x18, 0x0a, 0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x07, 0x61, 0x67, 0x65, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x41, 0x67, 0x65,
	0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0c, 0x52,
	0x0c, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1c, 0x0a,
	0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x2a, 0x43, 0x0a, 0x08, 0x50,
	0x65, 0x65, 0x72, 0x52, 0x6f, 0x6c, 0x65, 0x12, 0x11, 0x0a, 0x0d, 0x4c, 0x65, 0x67, 0x61, 0x63,
	0x79, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x65, 0x72, 0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x57, 0x61, 0x74, 0x63,
	0x68, 0x65, 0x72, 0x10, 0x02, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x67, 0x65, 0x6e, 0x74, 0x10, 0x03,
	0x42, 0x08, 0x5a, 0x06, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_node_proto_rawDescOnce sync.Once
	file_node_proto_rawDescData = file_node_proto_rawDesc
)

func file_node_proto_rawDescGZIP() []byte {
	file_node_proto_rawDescOnce.Do(func() {
		file_node_proto_rawDescData = protoimpl.X.CompressGZIP(file_node_proto_rawDescData)
	})
	return file_node_proto_rawDescData
}

var file_node_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_node_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_node_proto_goTypes = []interface{}{
	(PeerRole)(0),            // 0: types.PeerRole
	(*PeerAddress)(nil),      // 1: types.PeerAddress
	(*AgentCertificate)(nil), // 2: types.AgentCertificate
}
var file_node_proto_depIdxs = []int32{
	0, // 0: types.PeerAddress.role:type_name -> types.PeerRole
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_node_proto_init() }
func file_node_proto_init() {
	if File_node_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_node_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PeerAddress); i {
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
		file_node_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AgentCertificate); i {
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
			RawDescriptor: file_node_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_node_proto_goTypes,
		DependencyIndexes: file_node_proto_depIdxs,
		EnumInfos:         file_node_proto_enumTypes,
		MessageInfos:      file_node_proto_msgTypes,
	}.Build()
	File_node_proto = out.File
	file_node_proto_rawDesc = nil
	file_node_proto_goTypes = nil
	file_node_proto_depIdxs = nil
}