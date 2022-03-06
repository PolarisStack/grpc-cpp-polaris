// Code generated by protoc-gen-go. DO NOT EDIT.
// source: polaris_request.proto

package v1

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type DiscoverRequest_DiscoverRequestType int32

const (
	DiscoverRequest_UNKNOWN     DiscoverRequest_DiscoverRequestType = 0
	DiscoverRequest_INSTANCE    DiscoverRequest_DiscoverRequestType = 1
	DiscoverRequest_CLUSTER     DiscoverRequest_DiscoverRequestType = 2
	DiscoverRequest_ROUTING     DiscoverRequest_DiscoverRequestType = 3
	DiscoverRequest_RATE_LIMIT  DiscoverRequest_DiscoverRequestType = 4
	DiscoverRequest_SERVICES    DiscoverRequest_DiscoverRequestType = 6
	DiscoverRequest_MESH        DiscoverRequest_DiscoverRequestType = 7
	DiscoverRequest_MESH_CONFIG DiscoverRequest_DiscoverRequestType = 8
)

var DiscoverRequest_DiscoverRequestType_name = map[int32]string{
	0: "UNKNOWN",
	1: "INSTANCE",
	2: "CLUSTER",
	3: "ROUTING",
	4: "RATE_LIMIT",
	6: "SERVICES",
	7: "MESH",
	8: "MESH_CONFIG",
}
var DiscoverRequest_DiscoverRequestType_value = map[string]int32{
	"UNKNOWN":     0,
	"INSTANCE":    1,
	"CLUSTER":     2,
	"ROUTING":     3,
	"RATE_LIMIT":  4,
	"SERVICES":    6,
	"MESH":        7,
	"MESH_CONFIG": 8,
}

func (x DiscoverRequest_DiscoverRequestType) String() string {
	return proto.EnumName(DiscoverRequest_DiscoverRequestType_name, int32(x))
}
func (DiscoverRequest_DiscoverRequestType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_polaris_request_92cf89e67c2c4bdf, []int{0, 0}
}

type DiscoverRequest struct {
	Type                 DiscoverRequest_DiscoverRequestType `protobuf:"varint,1,opt,name=type,proto3,enum=v1.DiscoverRequest_DiscoverRequestType" json:"type,omitempty"`
	Service              *Service                            `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	Mesh                 *Mesh                               `protobuf:"bytes,3,opt,name=mesh,proto3" json:"mesh,omitempty"`
	MeshConfig           *MeshConfig                         `protobuf:"bytes,4,opt,name=meshConfig,proto3" json:"meshConfig,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *DiscoverRequest) Reset()         { *m = DiscoverRequest{} }
func (m *DiscoverRequest) String() string { return proto.CompactTextString(m) }
func (*DiscoverRequest) ProtoMessage()    {}
func (*DiscoverRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_polaris_request_92cf89e67c2c4bdf, []int{0}
}
func (m *DiscoverRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscoverRequest.Unmarshal(m, b)
}
func (m *DiscoverRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscoverRequest.Marshal(b, m, deterministic)
}
func (dst *DiscoverRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscoverRequest.Merge(dst, src)
}
func (m *DiscoverRequest) XXX_Size() int {
	return xxx_messageInfo_DiscoverRequest.Size(m)
}
func (m *DiscoverRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscoverRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DiscoverRequest proto.InternalMessageInfo

func (m *DiscoverRequest) GetType() DiscoverRequest_DiscoverRequestType {
	if m != nil {
		return m.Type
	}
	return DiscoverRequest_UNKNOWN
}

func (m *DiscoverRequest) GetService() *Service {
	if m != nil {
		return m.Service
	}
	return nil
}

func (m *DiscoverRequest) GetMesh() *Mesh {
	if m != nil {
		return m.Mesh
	}
	return nil
}

func (m *DiscoverRequest) GetMeshConfig() *MeshConfig {
	if m != nil {
		return m.MeshConfig
	}
	return nil
}

func init() {
	proto.RegisterType((*DiscoverRequest)(nil), "v1.DiscoverRequest")
	proto.RegisterEnum("v1.DiscoverRequest_DiscoverRequestType", DiscoverRequest_DiscoverRequestType_name, DiscoverRequest_DiscoverRequestType_value)
}

func init() {
	proto.RegisterFile("polaris_request.proto", fileDescriptor_polaris_request_92cf89e67c2c4bdf)
}

var fileDescriptor_polaris_request_92cf89e67c2c4bdf = []byte{
	// 294 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x41, 0x4b, 0xfb, 0x30,
	0x18, 0xc6, 0xff, 0xed, 0xca, 0x56, 0xde, 0xfe, 0xd9, 0x42, 0x44, 0x28, 0xc3, 0xc3, 0x18, 0x88,
	0x3b, 0x15, 0x36, 0x8f, 0x9e, 0x46, 0x8d, 0x33, 0xb8, 0xa5, 0x90, 0xa4, 0x7a, 0x2c, 0x5a, 0xa2,
	0x2b, 0xa8, 0xad, 0x4d, 0x2d, 0xec, 0xec, 0xa7, 0xf3, 0x5b, 0x49, 0x52, 0x27, 0x32, 0x3c, 0x25,
	0xef, 0xf3, 0xfc, 0x9e, 0xe4, 0xe1, 0x85, 0xe3, 0xaa, 0x7c, 0xbe, 0xaf, 0x0b, 0x9d, 0xd5, 0xea,
	0xed, 0x5d, 0xe9, 0x26, 0xaa, 0xea, 0xb2, 0x29, 0xb1, 0xdb, 0xce, 0xc7, 0x3f, 0x96, 0x56, 0x75,
	0x5b, 0xe4, 0xaa, 0xb3, 0xc6, 0xe8, 0x45, 0xe9, 0x6d, 0x5e, 0xbe, 0x3e, 0x16, 0x4f, 0x9d, 0x32,
	0xfd, 0x74, 0x61, 0x74, 0x59, 0xe8, 0xbc, 0x6c, 0x55, 0xcd, 0xbb, 0x67, 0xf0, 0x05, 0x78, 0xcd,
	0xae, 0x52, 0xa1, 0x33, 0x71, 0x66, 0xc3, 0xc5, 0x59, 0xd4, 0xce, 0xa3, 0x03, 0xe4, 0x70, 0x96,
	0xbb, 0x4a, 0x71, 0x1b, 0xc2, 0xa7, 0x30, 0xf8, 0xfe, 0x33, 0x74, 0x27, 0xce, 0x2c, 0x58, 0x04,
	0x26, 0x2f, 0x3a, 0x89, 0xef, 0x3d, 0x7c, 0x02, 0x9e, 0xe9, 0x12, 0xf6, 0x2c, 0xe3, 0x1b, 0x66,
	0xa3, 0xf4, 0x96, 0x5b, 0x15, 0x47, 0x00, 0xe6, 0x8c, 0x6d, 0xd3, 0xd0, 0xb3, 0xcc, 0x70, 0xcf,
	0x74, 0x2a, 0xff, 0x45, 0x4c, 0x3f, 0x1c, 0x38, 0xfa, 0xa3, 0x12, 0x0e, 0x60, 0x90, 0xb2, 0x1b,
	0x96, 0xdc, 0x31, 0xf4, 0x0f, 0xff, 0x07, 0x9f, 0x32, 0x21, 0x97, 0x2c, 0x26, 0xc8, 0x31, 0x56,
	0xbc, 0x4e, 0x85, 0x24, 0x1c, 0xb9, 0x66, 0xe0, 0x49, 0x2a, 0x29, 0x5b, 0xa1, 0x1e, 0x1e, 0x02,
	0xf0, 0xa5, 0x24, 0xd9, 0x9a, 0x6e, 0xa8, 0x44, 0x9e, 0xc9, 0x09, 0xc2, 0x6f, 0x69, 0x4c, 0x04,
	0xea, 0x63, 0x1f, 0xbc, 0x0d, 0x11, 0xd7, 0x68, 0x80, 0x47, 0x10, 0x98, 0x5b, 0x16, 0x27, 0xec,
	0x8a, 0xae, 0x90, 0xff, 0xd0, 0xb7, 0x2b, 0x3d, 0xff, 0x0a, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x6a,
	0x4c, 0xc5, 0x98, 0x01, 0x00, 0x00,
}