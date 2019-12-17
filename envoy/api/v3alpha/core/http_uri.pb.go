// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/api/v3alpha/core/http_uri.proto

package envoy_api_v3alpha_core

import (
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type HttpUri struct {
	Uri string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	// Types that are valid to be assigned to HttpUpstreamType:
	//	*HttpUri_Cluster
	HttpUpstreamType     isHttpUri_HttpUpstreamType `protobuf_oneof:"http_upstream_type"`
	Timeout              *duration.Duration         `protobuf:"bytes,3,opt,name=timeout,proto3" json:"timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *HttpUri) Reset()         { *m = HttpUri{} }
func (m *HttpUri) String() string { return proto.CompactTextString(m) }
func (*HttpUri) ProtoMessage()    {}
func (*HttpUri) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a09f7c0d394634e, []int{0}
}

func (m *HttpUri) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HttpUri.Unmarshal(m, b)
}
func (m *HttpUri) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HttpUri.Marshal(b, m, deterministic)
}
func (m *HttpUri) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HttpUri.Merge(m, src)
}
func (m *HttpUri) XXX_Size() int {
	return xxx_messageInfo_HttpUri.Size(m)
}
func (m *HttpUri) XXX_DiscardUnknown() {
	xxx_messageInfo_HttpUri.DiscardUnknown(m)
}

var xxx_messageInfo_HttpUri proto.InternalMessageInfo

func (m *HttpUri) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

type isHttpUri_HttpUpstreamType interface {
	isHttpUri_HttpUpstreamType()
}

type HttpUri_Cluster struct {
	Cluster string `protobuf:"bytes,2,opt,name=cluster,proto3,oneof"`
}

func (*HttpUri_Cluster) isHttpUri_HttpUpstreamType() {}

func (m *HttpUri) GetHttpUpstreamType() isHttpUri_HttpUpstreamType {
	if m != nil {
		return m.HttpUpstreamType
	}
	return nil
}

func (m *HttpUri) GetCluster() string {
	if x, ok := m.GetHttpUpstreamType().(*HttpUri_Cluster); ok {
		return x.Cluster
	}
	return ""
}

func (m *HttpUri) GetTimeout() *duration.Duration {
	if m != nil {
		return m.Timeout
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*HttpUri) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*HttpUri_Cluster)(nil),
	}
}

func init() {
	proto.RegisterType((*HttpUri)(nil), "envoy.api.v3alpha.core.HttpUri")
}

func init() {
	proto.RegisterFile("envoy/api/v3alpha/core/http_uri.proto", fileDescriptor_6a09f7c0d394634e)
}

var fileDescriptor_6a09f7c0d394634e = []byte{
	// 305 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0xcd, 0x4a, 0x03, 0x31,
	0x14, 0x85, 0x9b, 0xb6, 0x38, 0x1a, 0x5d, 0x65, 0xa1, 0x9d, 0x0a, 0x65, 0xfc, 0x83, 0xae, 0x12,
	0x68, 0xc1, 0x85, 0x1b, 0x21, 0xb8, 0xe8, 0xb2, 0x14, 0x5c, 0x97, 0xb4, 0x8d, 0x6d, 0x60, 0x9a,
	0x1b, 0x32, 0x37, 0x83, 0x7d, 0x03, 0x9f, 0xc1, 0x47, 0xf0, 0x3d, 0xc4, 0x57, 0x92, 0xae, 0xa4,
	0x99, 0x19, 0x50, 0x70, 0x17, 0xee, 0xf9, 0x4e, 0xf8, 0x38, 0xf4, 0x4e, 0xdb, 0x12, 0x76, 0x42,
	0x39, 0x23, 0xca, 0xb1, 0xca, 0xdd, 0x46, 0x89, 0x25, 0x78, 0x2d, 0x36, 0x88, 0x6e, 0x1e, 0xbc,
	0xe1, 0xce, 0x03, 0x02, 0x3b, 0x8f, 0x18, 0x57, 0xce, 0xf0, 0x1a, 0xe3, 0x07, 0xac, 0x3f, 0x58,
	0x03, 0xac, 0x73, 0x2d, 0x22, 0xb5, 0x08, 0x2f, 0x62, 0x15, 0xbc, 0x42, 0x03, 0xb6, 0xea, 0xf5,
	0xaf, 0xc2, 0xca, 0x29, 0xa1, 0xac, 0x05, 0x8c, 0xe7, 0x42, 0x94, 0xda, 0x17, 0x06, 0xac, 0xb1,
	0xeb, 0x1a, 0xb9, 0x28, 0x55, 0x6e, 0x56, 0x0a, 0xb5, 0x68, 0x1e, 0x55, 0x70, 0xfd, 0x45, 0x68,
	0x32, 0x41, 0x74, 0xcf, 0xde, 0xb0, 0x94, 0x76, 0x82, 0x37, 0x3d, 0x92, 0x91, 0xe1, 0x89, 0x4c,
	0xf6, 0xb2, 0xeb, 0xdb, 0x19, 0x99, 0x1d, 0x6e, 0xec, 0x86, 0x26, 0xcb, 0x3c, 0x14, 0xa8, 0x7d,
	0xaf, 0xfd, 0x27, 0x9e, 0xb4, 0x66, 0x4d, 0xc2, 0x1e, 0x69, 0x82, 0x66, 0xab, 0x21, 0x60, 0xaf,
	0x93, 0x91, 0xe1, 0xe9, 0x28, 0xe5, 0x95, 0x39, 0x6f, 0xcc, 0xf9, 0x53, 0x6d, 0x2e, 0xe9, 0x5e,
	0x26, 0x1f, 0xa4, 0x7b, 0x4c, 0x46, 0xad, 0x59, 0xd3, 0x7a, 0xc8, 0xde, 0x3f, 0xdf, 0x06, 0x97,
	0x34, 0xfd, 0xb5, 0xc3, 0x28, 0x4e, 0xc0, 0x6b, 0x45, 0x99, 0x52, 0x56, 0x8d, 0xe6, 0x0a, 0xf4,
	0x5a, 0x6d, 0xe7, 0xb8, 0x73, 0x9a, 0x75, 0xbe, 0x25, 0x91, 0xf7, 0xf4, 0xd6, 0x00, 0x8f, 0x55,
	0xe7, 0xe1, 0x75, 0xc7, 0xff, 0x5f, 0x53, 0x9e, 0xd5, 0x7f, 0x4d, 0x0f, 0x4e, 0x53, 0xb2, 0x38,
	0x8a, 0x72, 0xe3, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x7c, 0x2b, 0xe1, 0xec, 0xa5, 0x01, 0x00,
	0x00,
}