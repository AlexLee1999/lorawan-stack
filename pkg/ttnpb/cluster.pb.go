// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorawan-stack/api/cluster.proto

package ttnpb

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	golang_proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = golang_proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// PeerInfo
type PeerInfo struct {
	// Port on which the gRPC server is exposed.
	GrpcPort uint32 `protobuf:"varint,1,opt,name=grpc_port,json=grpcPort,proto3" json:"grpc_port,omitempty"`
	// Indicates whether the gRPC server uses TLS.
	Tls bool `protobuf:"varint,2,opt,name=tls,proto3" json:"tls,omitempty"`
	// Roles of the peer.
	Roles []ClusterRole `protobuf:"varint,3,rep,packed,name=roles,proto3,enum=ttn.lorawan.v3.ClusterRole" json:"roles,omitempty"`
	// Tags of the peer
	Tags                 map[string]string `protobuf:"bytes,4,rep,name=tags,proto3" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *PeerInfo) Reset()         { *m = PeerInfo{} }
func (m *PeerInfo) String() string { return proto.CompactTextString(m) }
func (*PeerInfo) ProtoMessage()    {}
func (*PeerInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_5716c3fcd711eefd, []int{0}
}
func (m *PeerInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PeerInfo.Unmarshal(m, b)
}
func (m *PeerInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PeerInfo.Marshal(b, m, deterministic)
}
func (m *PeerInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PeerInfo.Merge(m, src)
}
func (m *PeerInfo) XXX_Size() int {
	return xxx_messageInfo_PeerInfo.Size(m)
}
func (m *PeerInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PeerInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PeerInfo proto.InternalMessageInfo

func (m *PeerInfo) GetGrpcPort() uint32 {
	if m != nil {
		return m.GrpcPort
	}
	return 0
}

func (m *PeerInfo) GetTls() bool {
	if m != nil {
		return m.Tls
	}
	return false
}

func (m *PeerInfo) GetRoles() []ClusterRole {
	if m != nil {
		return m.Roles
	}
	return nil
}

func (m *PeerInfo) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func init() {
	proto.RegisterType((*PeerInfo)(nil), "ttn.lorawan.v3.PeerInfo")
	golang_proto.RegisterType((*PeerInfo)(nil), "ttn.lorawan.v3.PeerInfo")
	proto.RegisterMapType((map[string]string)(nil), "ttn.lorawan.v3.PeerInfo.TagsEntry")
	golang_proto.RegisterMapType((map[string]string)(nil), "ttn.lorawan.v3.PeerInfo.TagsEntry")
}

func init() { proto.RegisterFile("lorawan-stack/api/cluster.proto", fileDescriptor_5716c3fcd711eefd) }
func init() {
	golang_proto.RegisterFile("lorawan-stack/api/cluster.proto", fileDescriptor_5716c3fcd711eefd)
}

var fileDescriptor_5716c3fcd711eefd = []byte{
	// 300 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0x4f, 0x4b, 0xc3, 0x30,
	0x18, 0xc6, 0xe9, 0xba, 0xc9, 0x1a, 0x71, 0x48, 0xf1, 0x50, 0x36, 0xd4, 0xb2, 0x53, 0x2f, 0x4b,
	0x70, 0xc3, 0x3f, 0x78, 0x54, 0x3c, 0x78, 0x1b, 0xc1, 0x93, 0x17, 0x49, 0x4b, 0x96, 0x8d, 0x76,
	0x79, 0x43, 0xf2, 0xb6, 0x63, 0xdf, 0xce, 0x2f, 0xe2, 0xc9, 0x2f, 0x22, 0x4d, 0x55, 0x98, 0xde,
	0x9e, 0xe7, 0xcd, 0xf3, 0x86, 0xdf, 0xfb, 0x90, 0xcb, 0x0a, 0xac, 0xd8, 0x09, 0x3d, 0x73, 0x28,
	0x8a, 0x92, 0x09, 0xb3, 0x61, 0x45, 0x55, 0x3b, 0x94, 0x96, 0x1a, 0x0b, 0x08, 0xf1, 0x08, 0x51,
	0xd3, 0xef, 0x10, 0x6d, 0x16, 0xe3, 0x99, 0xda, 0xe0, 0xba, 0xce, 0x69, 0x01, 0x5b, 0xa6, 0x40,
	0x01, 0xf3, 0xb1, 0xbc, 0x5e, 0x79, 0xe7, 0x8d, 0x57, 0xdd, 0xfa, 0xf8, 0xfc, 0xff, 0xff, 0x52,
	0xd7, 0x5b, 0xd7, 0x3d, 0x4f, 0x3f, 0x02, 0x32, 0x5c, 0x4a, 0x69, 0x9f, 0xf5, 0x0a, 0xe2, 0x09,
	0x89, 0x94, 0x35, 0xc5, 0x9b, 0x01, 0x8b, 0x49, 0x90, 0x06, 0xd9, 0x09, 0x1f, 0xb6, 0x83, 0x25,
	0x58, 0x8c, 0x4f, 0x49, 0x88, 0x95, 0x4b, 0x7a, 0x69, 0x90, 0x0d, 0x79, 0x2b, 0xe3, 0x2b, 0x32,
	0xb0, 0x50, 0x49, 0x97, 0x84, 0x69, 0x98, 0x8d, 0xe6, 0x13, 0x7a, 0x48, 0x4a, 0x1f, 0xbb, 0x3b,
	0x38, 0x54, 0x92, 0x77, 0xc9, 0xf8, 0x86, 0xf4, 0x51, 0x28, 0x97, 0xf4, 0xd3, 0x30, 0x3b, 0x9e,
	0x4f, 0xff, 0x6e, 0xfc, 0x90, 0xd0, 0x17, 0xa1, 0xdc, 0x93, 0x46, 0xbb, 0xe7, 0x3e, 0x3f, 0xbe,
	0x25, 0xd1, 0xef, 0xa8, 0x25, 0x29, 0xe5, 0xde, 0x03, 0x46, 0xbc, 0x95, 0xf1, 0x19, 0x19, 0x34,
	0xa2, 0xaa, 0xa5, 0xa7, 0x8b, 0x78, 0x67, 0xee, 0x7b, 0x77, 0xc1, 0xc3, 0xf5, 0xfb, 0xe7, 0x45,
	0xf0, 0xca, 0x14, 0x50, 0x5c, 0x4b, 0x5c, 0x6f, 0xb4, 0x72, 0x54, 0x4b, 0xdc, 0x81, 0x2d, 0xd9,
	0x61, 0x33, 0xcd, 0x82, 0x99, 0x52, 0x31, 0x44, 0x6d, 0xf2, 0xfc, 0xc8, 0xb7, 0xb3, 0xf8, 0x0a,
	0x00, 0x00, 0xff, 0xff, 0x59, 0xea, 0x2f, 0x88, 0x9e, 0x01, 0x00, 0x00,
}
