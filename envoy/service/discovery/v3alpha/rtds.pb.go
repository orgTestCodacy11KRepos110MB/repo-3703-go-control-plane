// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/discovery/v3alpha/rtds.proto

package envoy_service_discovery_v3alpha

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	v3alpha "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_struct "github.com/golang/protobuf/ptypes/struct"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type RtdsDummy struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RtdsDummy) Reset()         { *m = RtdsDummy{} }
func (m *RtdsDummy) String() string { return proto.CompactTextString(m) }
func (*RtdsDummy) ProtoMessage()    {}
func (*RtdsDummy) Descriptor() ([]byte, []int) {
	return fileDescriptor_29c3a2e120a075f2, []int{0}
}

func (m *RtdsDummy) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RtdsDummy.Unmarshal(m, b)
}
func (m *RtdsDummy) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RtdsDummy.Marshal(b, m, deterministic)
}
func (m *RtdsDummy) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RtdsDummy.Merge(m, src)
}
func (m *RtdsDummy) XXX_Size() int {
	return xxx_messageInfo_RtdsDummy.Size(m)
}
func (m *RtdsDummy) XXX_DiscardUnknown() {
	xxx_messageInfo_RtdsDummy.DiscardUnknown(m)
}

var xxx_messageInfo_RtdsDummy proto.InternalMessageInfo

type Runtime struct {
	Name                 string          `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Layer                *_struct.Struct `protobuf:"bytes,2,opt,name=layer,proto3" json:"layer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Runtime) Reset()         { *m = Runtime{} }
func (m *Runtime) String() string { return proto.CompactTextString(m) }
func (*Runtime) ProtoMessage()    {}
func (*Runtime) Descriptor() ([]byte, []int) {
	return fileDescriptor_29c3a2e120a075f2, []int{1}
}

func (m *Runtime) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Runtime.Unmarshal(m, b)
}
func (m *Runtime) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Runtime.Marshal(b, m, deterministic)
}
func (m *Runtime) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Runtime.Merge(m, src)
}
func (m *Runtime) XXX_Size() int {
	return xxx_messageInfo_Runtime.Size(m)
}
func (m *Runtime) XXX_DiscardUnknown() {
	xxx_messageInfo_Runtime.DiscardUnknown(m)
}

var xxx_messageInfo_Runtime proto.InternalMessageInfo

func (m *Runtime) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Runtime) GetLayer() *_struct.Struct {
	if m != nil {
		return m.Layer
	}
	return nil
}

func init() {
	proto.RegisterType((*RtdsDummy)(nil), "envoy.service.discovery.v3alpha.RtdsDummy")
	proto.RegisterType((*Runtime)(nil), "envoy.service.discovery.v3alpha.Runtime")
}

func init() {
	proto.RegisterFile("envoy/service/discovery/v3alpha/rtds.proto", fileDescriptor_29c3a2e120a075f2)
}

var fileDescriptor_29c3a2e120a075f2 = []byte{
	// 412 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0xbf, 0x8e, 0xd3, 0x40,
	0x10, 0xc6, 0x59, 0x73, 0x70, 0xba, 0xe5, 0x68, 0xdc, 0xe4, 0x64, 0x4e, 0x90, 0x33, 0x11, 0xf2,
	0x25, 0xca, 0x1a, 0x25, 0x0d, 0x72, 0x83, 0x64, 0x45, 0xd4, 0x91, 0xf3, 0x02, 0x6c, 0xec, 0x21,
	0x59, 0xc9, 0xde, 0x35, 0xbb, 0x6b, 0x0b, 0x77, 0x88, 0x02, 0xa5, 0xa7, 0xe3, 0x7d, 0x78, 0x02,
	0x5e, 0x81, 0x9e, 0x9e, 0x0a, 0x65, 0xd7, 0x36, 0x7f, 0x02, 0x81, 0xe6, 0xba, 0xb5, 0xe7, 0xf7,
	0xcd, 0x37, 0x9f, 0x66, 0xf0, 0x18, 0x78, 0x2d, 0x9a, 0x50, 0x81, 0xac, 0x59, 0x0a, 0x61, 0xc6,
	0x54, 0x2a, 0x6a, 0x90, 0x4d, 0x58, 0xcf, 0x69, 0x5e, 0x6e, 0x69, 0x28, 0x75, 0xa6, 0x48, 0x29,
	0x85, 0x16, 0xee, 0x23, 0xc3, 0x92, 0x96, 0x25, 0x3d, 0x4b, 0x5a, 0xd6, 0xbb, 0xb2, 0xcd, 0x68,
	0xc9, 0x7a, 0xf9, 0x0f, 0xc8, 0xf4, 0xf0, 0x2e, 0x37, 0x42, 0x6c, 0x72, 0x30, 0x0c, 0xe5, 0x5c,
	0x68, 0xaa, 0x99, 0xe0, 0xea, 0xb7, 0xaa, 0xf9, 0x5a, 0x57, 0xaf, 0x42, 0xa5, 0x65, 0x95, 0xea,
	0xb6, 0x7a, 0x55, 0x65, 0x25, 0xfd, 0x59, 0x15, 0xd6, 0x20, 0x15, 0x13, 0x9c, 0xf1, 0x4d, 0x8b,
	0x0c, 0x6a, 0x9a, 0xb3, 0x8c, 0x6a, 0x08, 0xbb, 0x87, 0x2d, 0xf8, 0xcf, 0xf0, 0x59, 0xa2, 0x33,
	0xb5, 0xa8, 0x8a, 0xa2, 0x89, 0x26, 0x1f, 0x3f, 0xed, 0x1e, 0x3e, 0xc1, 0xa3, 0xbf, 0xe6, 0x99,
	0x91, 0x1e, 0xf6, 0xdf, 0x22, 0x7c, 0x9a, 0x54, 0x5c, 0xb3, 0x02, 0xdc, 0x07, 0xf8, 0x84, 0xd3,
	0x02, 0x2e, 0xd0, 0x10, 0x05, 0x67, 0xf1, 0xe9, 0xb7, 0xf8, 0x44, 0x3a, 0x43, 0x94, 0x98, 0x9f,
	0xee, 0x14, 0xdf, 0xc9, 0x69, 0x03, 0xf2, 0xc2, 0x19, 0xa2, 0xe0, 0xde, 0x6c, 0x40, 0x6c, 0x18,
	0xd2, 0x85, 0x21, 0x2b, 0x13, 0x26, 0xb1, 0x54, 0x74, 0xbd, 0x1f, 0x62, 0x84, 0xfd, 0x63, 0x43,
	0x58, 0xdb, 0xd9, 0x57, 0x07, 0x0f, 0xda, 0xf7, 0xa2, 0xab, 0xaf, 0xac, 0xc0, 0x7d, 0x89, 0xef,
	0xaf, 0xb4, 0x04, 0x5a, 0x74, 0x33, 0x3e, 0x26, 0xb6, 0x23, 0x2d, 0x59, 0xb7, 0x18, 0xd2, 0xab,
	0x12, 0x78, 0x5d, 0x81, 0xd2, 0xde, 0xe8, 0x38, 0xa4, 0x4a, 0xc1, 0x15, 0xf8, 0xb7, 0x02, 0xf4,
	0x14, 0xb9, 0x0c, 0x9f, 0x2f, 0x20, 0xd7, 0xb4, 0x33, 0x08, 0xfe, 0xa4, 0xdd, 0x03, 0x07, 0x2e,
	0xd7, 0xff, 0x41, 0xfe, 0x62, 0xf5, 0x1e, 0xe1, 0xf3, 0x17, 0xa0, 0xd3, 0xed, 0x0d, 0x84, 0x99,
	0xbc, 0xfb, 0xfc, 0xe5, 0x83, 0x73, 0xe9, 0x7b, 0x87, 0xb7, 0x19, 0x49, 0x6b, 0x67, 0x88, 0xdb,
	0x11, 0x1a, 0xc7, 0xcf, 0xf1, 0x94, 0x09, 0xdb, 0xb6, 0x94, 0xe2, 0x4d, 0x43, 0xfe, 0x71, 0xfa,
	0xb1, 0xb9, 0xae, 0xe5, 0x7e, 0xd3, 0x4b, 0xb4, 0x43, 0x68, 0x7d, 0xd7, 0x6c, 0x7d, 0xfe, 0x3d,
	0x00, 0x00, 0xff, 0xff, 0x62, 0xbe, 0x53, 0x4b, 0x60, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RuntimeDiscoveryServiceClient is the client API for RuntimeDiscoveryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RuntimeDiscoveryServiceClient interface {
	StreamRuntime(ctx context.Context, opts ...grpc.CallOption) (RuntimeDiscoveryService_StreamRuntimeClient, error)
	DeltaRuntime(ctx context.Context, opts ...grpc.CallOption) (RuntimeDiscoveryService_DeltaRuntimeClient, error)
	FetchRuntime(ctx context.Context, in *v3alpha.DiscoveryRequest, opts ...grpc.CallOption) (*v3alpha.DiscoveryResponse, error)
}

type runtimeDiscoveryServiceClient struct {
	cc *grpc.ClientConn
}

func NewRuntimeDiscoveryServiceClient(cc *grpc.ClientConn) RuntimeDiscoveryServiceClient {
	return &runtimeDiscoveryServiceClient{cc}
}

func (c *runtimeDiscoveryServiceClient) StreamRuntime(ctx context.Context, opts ...grpc.CallOption) (RuntimeDiscoveryService_StreamRuntimeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RuntimeDiscoveryService_serviceDesc.Streams[0], "/envoy.service.discovery.v3alpha.RuntimeDiscoveryService/StreamRuntime", opts...)
	if err != nil {
		return nil, err
	}
	x := &runtimeDiscoveryServiceStreamRuntimeClient{stream}
	return x, nil
}

type RuntimeDiscoveryService_StreamRuntimeClient interface {
	Send(*v3alpha.DiscoveryRequest) error
	Recv() (*v3alpha.DiscoveryResponse, error)
	grpc.ClientStream
}

type runtimeDiscoveryServiceStreamRuntimeClient struct {
	grpc.ClientStream
}

func (x *runtimeDiscoveryServiceStreamRuntimeClient) Send(m *v3alpha.DiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *runtimeDiscoveryServiceStreamRuntimeClient) Recv() (*v3alpha.DiscoveryResponse, error) {
	m := new(v3alpha.DiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *runtimeDiscoveryServiceClient) DeltaRuntime(ctx context.Context, opts ...grpc.CallOption) (RuntimeDiscoveryService_DeltaRuntimeClient, error) {
	stream, err := c.cc.NewStream(ctx, &_RuntimeDiscoveryService_serviceDesc.Streams[1], "/envoy.service.discovery.v3alpha.RuntimeDiscoveryService/DeltaRuntime", opts...)
	if err != nil {
		return nil, err
	}
	x := &runtimeDiscoveryServiceDeltaRuntimeClient{stream}
	return x, nil
}

type RuntimeDiscoveryService_DeltaRuntimeClient interface {
	Send(*v3alpha.DeltaDiscoveryRequest) error
	Recv() (*v3alpha.DeltaDiscoveryResponse, error)
	grpc.ClientStream
}

type runtimeDiscoveryServiceDeltaRuntimeClient struct {
	grpc.ClientStream
}

func (x *runtimeDiscoveryServiceDeltaRuntimeClient) Send(m *v3alpha.DeltaDiscoveryRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *runtimeDiscoveryServiceDeltaRuntimeClient) Recv() (*v3alpha.DeltaDiscoveryResponse, error) {
	m := new(v3alpha.DeltaDiscoveryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *runtimeDiscoveryServiceClient) FetchRuntime(ctx context.Context, in *v3alpha.DiscoveryRequest, opts ...grpc.CallOption) (*v3alpha.DiscoveryResponse, error) {
	out := new(v3alpha.DiscoveryResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.discovery.v3alpha.RuntimeDiscoveryService/FetchRuntime", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RuntimeDiscoveryServiceServer is the server API for RuntimeDiscoveryService service.
type RuntimeDiscoveryServiceServer interface {
	StreamRuntime(RuntimeDiscoveryService_StreamRuntimeServer) error
	DeltaRuntime(RuntimeDiscoveryService_DeltaRuntimeServer) error
	FetchRuntime(context.Context, *v3alpha.DiscoveryRequest) (*v3alpha.DiscoveryResponse, error)
}

// UnimplementedRuntimeDiscoveryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedRuntimeDiscoveryServiceServer struct {
}

func (*UnimplementedRuntimeDiscoveryServiceServer) StreamRuntime(srv RuntimeDiscoveryService_StreamRuntimeServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamRuntime not implemented")
}
func (*UnimplementedRuntimeDiscoveryServiceServer) DeltaRuntime(srv RuntimeDiscoveryService_DeltaRuntimeServer) error {
	return status.Errorf(codes.Unimplemented, "method DeltaRuntime not implemented")
}
func (*UnimplementedRuntimeDiscoveryServiceServer) FetchRuntime(ctx context.Context, req *v3alpha.DiscoveryRequest) (*v3alpha.DiscoveryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchRuntime not implemented")
}

func RegisterRuntimeDiscoveryServiceServer(s *grpc.Server, srv RuntimeDiscoveryServiceServer) {
	s.RegisterService(&_RuntimeDiscoveryService_serviceDesc, srv)
}

func _RuntimeDiscoveryService_StreamRuntime_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RuntimeDiscoveryServiceServer).StreamRuntime(&runtimeDiscoveryServiceStreamRuntimeServer{stream})
}

type RuntimeDiscoveryService_StreamRuntimeServer interface {
	Send(*v3alpha.DiscoveryResponse) error
	Recv() (*v3alpha.DiscoveryRequest, error)
	grpc.ServerStream
}

type runtimeDiscoveryServiceStreamRuntimeServer struct {
	grpc.ServerStream
}

func (x *runtimeDiscoveryServiceStreamRuntimeServer) Send(m *v3alpha.DiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *runtimeDiscoveryServiceStreamRuntimeServer) Recv() (*v3alpha.DiscoveryRequest, error) {
	m := new(v3alpha.DiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RuntimeDiscoveryService_DeltaRuntime_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(RuntimeDiscoveryServiceServer).DeltaRuntime(&runtimeDiscoveryServiceDeltaRuntimeServer{stream})
}

type RuntimeDiscoveryService_DeltaRuntimeServer interface {
	Send(*v3alpha.DeltaDiscoveryResponse) error
	Recv() (*v3alpha.DeltaDiscoveryRequest, error)
	grpc.ServerStream
}

type runtimeDiscoveryServiceDeltaRuntimeServer struct {
	grpc.ServerStream
}

func (x *runtimeDiscoveryServiceDeltaRuntimeServer) Send(m *v3alpha.DeltaDiscoveryResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *runtimeDiscoveryServiceDeltaRuntimeServer) Recv() (*v3alpha.DeltaDiscoveryRequest, error) {
	m := new(v3alpha.DeltaDiscoveryRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _RuntimeDiscoveryService_FetchRuntime_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v3alpha.DiscoveryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RuntimeDiscoveryServiceServer).FetchRuntime(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.discovery.v3alpha.RuntimeDiscoveryService/FetchRuntime",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RuntimeDiscoveryServiceServer).FetchRuntime(ctx, req.(*v3alpha.DiscoveryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RuntimeDiscoveryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.discovery.v3alpha.RuntimeDiscoveryService",
	HandlerType: (*RuntimeDiscoveryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FetchRuntime",
			Handler:    _RuntimeDiscoveryService_FetchRuntime_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamRuntime",
			Handler:       _RuntimeDiscoveryService_StreamRuntime_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
		{
			StreamName:    "DeltaRuntime",
			Handler:       _RuntimeDiscoveryService_DeltaRuntime_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/discovery/v3alpha/rtds.proto",
}