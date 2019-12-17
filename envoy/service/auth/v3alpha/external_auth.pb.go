// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/auth/v3alpha/external_auth.proto

package envoy_service_auth_v3alpha

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/core"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/type/v3alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	status "google.golang.org/genproto/googleapis/rpc/status"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status1 "google.golang.org/grpc/status"
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

type CheckRequest struct {
	Attributes           *AttributeContext `protobuf:"bytes,1,opt,name=attributes,proto3" json:"attributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *CheckRequest) Reset()         { *m = CheckRequest{} }
func (m *CheckRequest) String() string { return proto.CompactTextString(m) }
func (*CheckRequest) ProtoMessage()    {}
func (*CheckRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_52e899d802958088, []int{0}
}

func (m *CheckRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckRequest.Unmarshal(m, b)
}
func (m *CheckRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckRequest.Marshal(b, m, deterministic)
}
func (m *CheckRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckRequest.Merge(m, src)
}
func (m *CheckRequest) XXX_Size() int {
	return xxx_messageInfo_CheckRequest.Size(m)
}
func (m *CheckRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CheckRequest proto.InternalMessageInfo

func (m *CheckRequest) GetAttributes() *AttributeContext {
	if m != nil {
		return m.Attributes
	}
	return nil
}

type DeniedHttpResponse struct {
	Status               *v3alpha.HttpStatus       `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	Headers              []*core.HeaderValueOption `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty"`
	Body                 string                    `protobuf:"bytes,3,opt,name=body,proto3" json:"body,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *DeniedHttpResponse) Reset()         { *m = DeniedHttpResponse{} }
func (m *DeniedHttpResponse) String() string { return proto.CompactTextString(m) }
func (*DeniedHttpResponse) ProtoMessage()    {}
func (*DeniedHttpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_52e899d802958088, []int{1}
}

func (m *DeniedHttpResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DeniedHttpResponse.Unmarshal(m, b)
}
func (m *DeniedHttpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DeniedHttpResponse.Marshal(b, m, deterministic)
}
func (m *DeniedHttpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DeniedHttpResponse.Merge(m, src)
}
func (m *DeniedHttpResponse) XXX_Size() int {
	return xxx_messageInfo_DeniedHttpResponse.Size(m)
}
func (m *DeniedHttpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DeniedHttpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DeniedHttpResponse proto.InternalMessageInfo

func (m *DeniedHttpResponse) GetStatus() *v3alpha.HttpStatus {
	if m != nil {
		return m.Status
	}
	return nil
}

func (m *DeniedHttpResponse) GetHeaders() []*core.HeaderValueOption {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *DeniedHttpResponse) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

type OkHttpResponse struct {
	Headers              []*core.HeaderValueOption `protobuf:"bytes,2,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *OkHttpResponse) Reset()         { *m = OkHttpResponse{} }
func (m *OkHttpResponse) String() string { return proto.CompactTextString(m) }
func (*OkHttpResponse) ProtoMessage()    {}
func (*OkHttpResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_52e899d802958088, []int{2}
}

func (m *OkHttpResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OkHttpResponse.Unmarshal(m, b)
}
func (m *OkHttpResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OkHttpResponse.Marshal(b, m, deterministic)
}
func (m *OkHttpResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OkHttpResponse.Merge(m, src)
}
func (m *OkHttpResponse) XXX_Size() int {
	return xxx_messageInfo_OkHttpResponse.Size(m)
}
func (m *OkHttpResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OkHttpResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OkHttpResponse proto.InternalMessageInfo

func (m *OkHttpResponse) GetHeaders() []*core.HeaderValueOption {
	if m != nil {
		return m.Headers
	}
	return nil
}

type CheckResponse struct {
	Status *status.Status `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// Types that are valid to be assigned to HttpResponse:
	//	*CheckResponse_DeniedResponse
	//	*CheckResponse_OkResponse
	HttpResponse         isCheckResponse_HttpResponse `protobuf_oneof:"http_response"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *CheckResponse) Reset()         { *m = CheckResponse{} }
func (m *CheckResponse) String() string { return proto.CompactTextString(m) }
func (*CheckResponse) ProtoMessage()    {}
func (*CheckResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_52e899d802958088, []int{3}
}

func (m *CheckResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CheckResponse.Unmarshal(m, b)
}
func (m *CheckResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CheckResponse.Marshal(b, m, deterministic)
}
func (m *CheckResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CheckResponse.Merge(m, src)
}
func (m *CheckResponse) XXX_Size() int {
	return xxx_messageInfo_CheckResponse.Size(m)
}
func (m *CheckResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CheckResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CheckResponse proto.InternalMessageInfo

func (m *CheckResponse) GetStatus() *status.Status {
	if m != nil {
		return m.Status
	}
	return nil
}

type isCheckResponse_HttpResponse interface {
	isCheckResponse_HttpResponse()
}

type CheckResponse_DeniedResponse struct {
	DeniedResponse *DeniedHttpResponse `protobuf:"bytes,2,opt,name=denied_response,json=deniedResponse,proto3,oneof"`
}

type CheckResponse_OkResponse struct {
	OkResponse *OkHttpResponse `protobuf:"bytes,3,opt,name=ok_response,json=okResponse,proto3,oneof"`
}

func (*CheckResponse_DeniedResponse) isCheckResponse_HttpResponse() {}

func (*CheckResponse_OkResponse) isCheckResponse_HttpResponse() {}

func (m *CheckResponse) GetHttpResponse() isCheckResponse_HttpResponse {
	if m != nil {
		return m.HttpResponse
	}
	return nil
}

func (m *CheckResponse) GetDeniedResponse() *DeniedHttpResponse {
	if x, ok := m.GetHttpResponse().(*CheckResponse_DeniedResponse); ok {
		return x.DeniedResponse
	}
	return nil
}

func (m *CheckResponse) GetOkResponse() *OkHttpResponse {
	if x, ok := m.GetHttpResponse().(*CheckResponse_OkResponse); ok {
		return x.OkResponse
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*CheckResponse) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*CheckResponse_DeniedResponse)(nil),
		(*CheckResponse_OkResponse)(nil),
	}
}

func init() {
	proto.RegisterType((*CheckRequest)(nil), "envoy.service.auth.v3alpha.CheckRequest")
	proto.RegisterType((*DeniedHttpResponse)(nil), "envoy.service.auth.v3alpha.DeniedHttpResponse")
	proto.RegisterType((*OkHttpResponse)(nil), "envoy.service.auth.v3alpha.OkHttpResponse")
	proto.RegisterType((*CheckResponse)(nil), "envoy.service.auth.v3alpha.CheckResponse")
}

func init() {
	proto.RegisterFile("envoy/service/auth/v3alpha/external_auth.proto", fileDescriptor_52e899d802958088)
}

var fileDescriptor_52e899d802958088 = []byte{
	// 552 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x54, 0x41, 0x6f, 0xd3, 0x30,
	0x18, 0x5d, 0x3a, 0x36, 0xc0, 0xa3, 0x1b, 0xe4, 0x42, 0xd5, 0xc3, 0xd4, 0x95, 0x82, 0xd2, 0x82,
	0x6c, 0xa9, 0xbb, 0xf5, 0xc4, 0x52, 0x90, 0x7a, 0x00, 0x6d, 0x0a, 0x12, 0x12, 0x17, 0x2a, 0x37,
	0xf9, 0xd4, 0x58, 0xad, 0x62, 0x63, 0x3b, 0x55, 0xcb, 0x91, 0x0b, 0x13, 0x3f, 0x81, 0x23, 0xff,
	0x85, 0xff, 0xc1, 0xef, 0xe0, 0x84, 0x62, 0xbb, 0xa1, 0x65, 0x25, 0xe2, 0xc0, 0x2d, 0xaa, 0xdf,
	0x7b, 0x7e, 0xef, 0xf9, 0xfb, 0x8a, 0x30, 0x64, 0x0b, 0xbe, 0x22, 0x0a, 0xe4, 0x82, 0xc5, 0x40,
	0x68, 0xae, 0x53, 0xb2, 0x38, 0xa7, 0x73, 0x91, 0x52, 0x02, 0x4b, 0x0d, 0x32, 0xa3, 0xf3, 0x71,
	0xf1, 0x2b, 0x16, 0x92, 0x6b, 0xee, 0x37, 0x0d, 0x1e, 0x3b, 0x3c, 0x36, 0x27, 0x0e, 0xdf, 0x3c,
	0xb3, 0x5a, 0x54, 0xb0, 0x52, 0x22, 0xe6, 0x12, 0xc8, 0x84, 0x2a, 0xb0, 0xf4, 0x66, 0xbf, 0xe2,
	0x3a, 0xaa, 0xb5, 0x64, 0x93, 0x5c, 0xc3, 0x38, 0xe6, 0x99, 0x86, 0xa5, 0x76, 0x9c, 0x8e, 0xe5,
	0xe8, 0x95, 0x80, 0x12, 0x9b, 0x6a, 0x2d, 0xc6, 0x4a, 0x53, 0x9d, 0x2b, 0x87, 0x7a, 0x38, 0xe5,
	0x7c, 0x3a, 0x07, 0x22, 0x45, 0x4c, 0xb6, 0x0e, 0xce, 0xf2, 0x44, 0x50, 0x42, 0xb3, 0x8c, 0x6b,
	0xaa, 0x19, 0xcf, 0x14, 0x59, 0x80, 0x54, 0x8c, 0x67, 0x2c, 0x9b, 0xae, 0xb9, 0x0b, 0x3a, 0x67,
	0x09, 0xd5, 0x40, 0xd6, 0x1f, 0xf6, 0xa0, 0xfd, 0xd9, 0x43, 0xf7, 0x86, 0x29, 0xc4, 0xb3, 0x08,
	0x3e, 0xe4, 0xa0, 0xb4, 0xff, 0x0a, 0xa1, 0xd2, 0xa6, 0x6a, 0x78, 0x2d, 0x2f, 0x38, 0xea, 0x3f,
	0xc3, 0x7f, 0xef, 0x04, 0x5f, 0xac, 0xd1, 0x43, 0x9b, 0x29, 0xda, 0xe0, 0x0f, 0xba, 0x5f, 0xbf,
	0x5f, 0x9f, 0x76, 0x50, 0x7b, 0x17, 0xbf, 0x8f, 0x37, 0x2f, 0x6e, 0xff, 0xf0, 0x90, 0xff, 0x02,
	0x32, 0x06, 0xc9, 0x48, 0x6b, 0x11, 0x81, 0x12, 0x3c, 0x53, 0xe0, 0x3f, 0x47, 0x87, 0x36, 0xac,
	0xf3, 0x72, 0xea, 0xbc, 0x14, 0x65, 0x95, 0x1e, 0x0a, 0xc6, 0x1b, 0x83, 0x0a, 0xef, 0xfc, 0x0c,
	0x0f, 0xbe, 0x78, 0xb5, 0xfb, 0x5e, 0xe4, 0x78, 0xfe, 0x10, 0xdd, 0x4e, 0x81, 0x26, 0x20, 0x55,
	0xa3, 0xd6, 0xda, 0x0f, 0x8e, 0xfa, 0x5d, 0x27, 0x41, 0x05, 0x2b, 0x15, 0x8a, 0x67, 0xc4, 0x23,
	0x03, 0x7b, 0x4b, 0xe7, 0x39, 0x5c, 0x8a, 0xa2, 0xc9, 0x68, 0xcd, 0xf4, 0x7d, 0x74, 0x6b, 0xc2,
	0x93, 0x55, 0x63, 0xbf, 0xe5, 0x05, 0x77, 0x23, 0xf3, 0x3d, 0x20, 0x45, 0xb8, 0x1e, 0x0a, 0x76,
	0x87, 0xbb, 0x99, 0xa5, 0xfd, 0xc9, 0x43, 0xc7, 0x97, 0xb3, 0xad, 0x78, 0xff, 0xc3, 0xdc, 0xe0,
	0x69, 0x61, 0xe4, 0x09, 0xea, 0xec, 0x36, 0xb2, 0x7d, 0x63, 0xfb, 0x5b, 0x0d, 0xd5, 0x5d, 0xf1,
	0xce, 0x43, 0xef, 0x8f, 0x8a, 0x7d, 0x6c, 0x27, 0x0d, 0x4b, 0x11, 0x63, 0x5b, 0x6b, 0x59, 0xe6,
	0x3b, 0x74, 0x92, 0x98, 0x60, 0x63, 0xe9, 0xe8, 0x8d, 0x9a, 0x21, 0xe1, 0xaa, 0x19, 0xb9, 0xd9,
	0xc5, 0x68, 0x2f, 0x3a, 0xb6, 0x42, 0xa5, 0x8d, 0xd7, 0xe8, 0x88, 0xcf, 0x7e, 0xcb, 0xee, 0x1b,
	0xd9, 0x5e, 0x95, 0xec, 0x76, 0xb2, 0xd1, 0x5e, 0x84, 0x78, 0x99, 0x6a, 0xd0, 0x2b, 0x4a, 0x79,
	0x8c, 0x1e, 0x55, 0x8e, 0x9e, 0xc5, 0x86, 0x27, 0xa8, 0x6e, 0xf6, 0x6d, 0x7d, 0x79, 0x9f, 0xa3,
	0xfa, 0x45, 0xae, 0x53, 0x2e, 0xd9, 0x47, 0xb3, 0x52, 0xfe, 0x7b, 0x74, 0x60, 0x28, 0x7e, 0x50,
	0x65, 0x68, 0x73, 0xa0, 0x9b, 0xdd, 0x7f, 0x40, 0xba, 0x37, 0xd9, 0x0b, 0x43, 0x14, 0x30, 0x6e,
	0x09, 0x42, 0xf2, 0xe5, 0xaa, 0x82, 0x1b, 0x3e, 0x78, 0xe9, 0xfe, 0xb6, 0x0a, 0x8b, 0x57, 0xc5,
	0x1a, 0x5f, 0x79, 0xd7, 0x9e, 0x37, 0x39, 0x34, 0x2b, 0x7d, 0xfe, 0x2b, 0x00, 0x00, 0xff, 0xff,
	0x03, 0x64, 0xd2, 0x2c, 0xf2, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// AuthorizationClient is the client API for Authorization service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type AuthorizationClient interface {
	Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error)
}

type authorizationClient struct {
	cc *grpc.ClientConn
}

func NewAuthorizationClient(cc *grpc.ClientConn) AuthorizationClient {
	return &authorizationClient{cc}
}

func (c *authorizationClient) Check(ctx context.Context, in *CheckRequest, opts ...grpc.CallOption) (*CheckResponse, error) {
	out := new(CheckResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.auth.v3alpha.Authorization/Check", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AuthorizationServer is the server API for Authorization service.
type AuthorizationServer interface {
	Check(context.Context, *CheckRequest) (*CheckResponse, error)
}

// UnimplementedAuthorizationServer can be embedded to have forward compatible implementations.
type UnimplementedAuthorizationServer struct {
}

func (*UnimplementedAuthorizationServer) Check(ctx context.Context, req *CheckRequest) (*CheckResponse, error) {
	return nil, status1.Errorf(codes.Unimplemented, "method Check not implemented")
}

func RegisterAuthorizationServer(s *grpc.Server, srv AuthorizationServer) {
	s.RegisterService(&_Authorization_serviceDesc, srv)
}

func _Authorization_Check_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AuthorizationServer).Check(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.auth.v3alpha.Authorization/Check",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AuthorizationServer).Check(ctx, req.(*CheckRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Authorization_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.auth.v3alpha.Authorization",
	HandlerType: (*AuthorizationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Check",
			Handler:    _Authorization_Check_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "envoy/service/auth/v3alpha/external_auth.proto",
}