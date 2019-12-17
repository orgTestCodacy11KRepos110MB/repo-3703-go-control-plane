// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/load_stats/v3alpha/lrs.proto

package envoy_service_load_stats_v3alpha

import (
	context "context"
	fmt "fmt"
	_ "github.com/cncf/udpa/go/udpa/annotations"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/core"
	endpoint "github.com/envoyproxy/go-control-plane/envoy/api/v3alpha/endpoint"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type LoadStatsRequest struct {
	Node                 *core.Node               `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	ClusterStats         []*endpoint.ClusterStats `protobuf:"bytes,2,rep,name=cluster_stats,json=clusterStats,proto3" json:"cluster_stats,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                 `json:"-"`
	XXX_unrecognized     []byte                   `json:"-"`
	XXX_sizecache        int32                    `json:"-"`
}

func (m *LoadStatsRequest) Reset()         { *m = LoadStatsRequest{} }
func (m *LoadStatsRequest) String() string { return proto.CompactTextString(m) }
func (*LoadStatsRequest) ProtoMessage()    {}
func (*LoadStatsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_5313d77ae958858e, []int{0}
}

func (m *LoadStatsRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadStatsRequest.Unmarshal(m, b)
}
func (m *LoadStatsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadStatsRequest.Marshal(b, m, deterministic)
}
func (m *LoadStatsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadStatsRequest.Merge(m, src)
}
func (m *LoadStatsRequest) XXX_Size() int {
	return xxx_messageInfo_LoadStatsRequest.Size(m)
}
func (m *LoadStatsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadStatsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_LoadStatsRequest proto.InternalMessageInfo

func (m *LoadStatsRequest) GetNode() *core.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func (m *LoadStatsRequest) GetClusterStats() []*endpoint.ClusterStats {
	if m != nil {
		return m.ClusterStats
	}
	return nil
}

type LoadStatsResponse struct {
	Clusters                  []string           `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty"`
	LoadReportingInterval     *duration.Duration `protobuf:"bytes,2,opt,name=load_reporting_interval,json=loadReportingInterval,proto3" json:"load_reporting_interval,omitempty"`
	ReportEndpointGranularity bool               `protobuf:"varint,3,opt,name=report_endpoint_granularity,json=reportEndpointGranularity,proto3" json:"report_endpoint_granularity,omitempty"`
	XXX_NoUnkeyedLiteral      struct{}           `json:"-"`
	XXX_unrecognized          []byte             `json:"-"`
	XXX_sizecache             int32              `json:"-"`
}

func (m *LoadStatsResponse) Reset()         { *m = LoadStatsResponse{} }
func (m *LoadStatsResponse) String() string { return proto.CompactTextString(m) }
func (*LoadStatsResponse) ProtoMessage()    {}
func (*LoadStatsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_5313d77ae958858e, []int{1}
}

func (m *LoadStatsResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LoadStatsResponse.Unmarshal(m, b)
}
func (m *LoadStatsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LoadStatsResponse.Marshal(b, m, deterministic)
}
func (m *LoadStatsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LoadStatsResponse.Merge(m, src)
}
func (m *LoadStatsResponse) XXX_Size() int {
	return xxx_messageInfo_LoadStatsResponse.Size(m)
}
func (m *LoadStatsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_LoadStatsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_LoadStatsResponse proto.InternalMessageInfo

func (m *LoadStatsResponse) GetClusters() []string {
	if m != nil {
		return m.Clusters
	}
	return nil
}

func (m *LoadStatsResponse) GetLoadReportingInterval() *duration.Duration {
	if m != nil {
		return m.LoadReportingInterval
	}
	return nil
}

func (m *LoadStatsResponse) GetReportEndpointGranularity() bool {
	if m != nil {
		return m.ReportEndpointGranularity
	}
	return false
}

func init() {
	proto.RegisterType((*LoadStatsRequest)(nil), "envoy.service.load_stats.v3alpha.LoadStatsRequest")
	proto.RegisterType((*LoadStatsResponse)(nil), "envoy.service.load_stats.v3alpha.LoadStatsResponse")
}

func init() {
	proto.RegisterFile("envoy/service/load_stats/v3alpha/lrs.proto", fileDescriptor_5313d77ae958858e)
}

var fileDescriptor_5313d77ae958858e = []byte{
	// 474 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xc1, 0x8e, 0xd3, 0x3e,
	0x10, 0xc6, 0xff, 0x6e, 0xff, 0xa0, 0xe0, 0x05, 0x01, 0x11, 0x68, 0xb3, 0x05, 0xad, 0xb2, 0x15,
	0x87, 0x08, 0x2d, 0xf6, 0x2a, 0xe5, 0xb4, 0x07, 0x90, 0x0a, 0x08, 0x21, 0x15, 0xb4, 0xa4, 0x0f,
	0x50, 0xb9, 0xcd, 0x10, 0x2c, 0x05, 0x4f, 0xb0, 0x9d, 0x88, 0xde, 0x10, 0xa7, 0x3d, 0x73, 0x83,
	0xf7, 0x41, 0xe2, 0x99, 0x38, 0xa1, 0xc4, 0x4e, 0xb7, 0x62, 0x41, 0x0b, 0xb7, 0xd8, 0xf3, 0xfb,
	0xc6, 0xf3, 0x7d, 0x13, 0x7a, 0x1f, 0x54, 0x83, 0x6b, 0x6e, 0x40, 0x37, 0x72, 0x05, 0xbc, 0x44,
	0x91, 0x2f, 0x8c, 0x15, 0xd6, 0xf0, 0x66, 0x22, 0xca, 0xea, 0xad, 0xe0, 0xa5, 0x36, 0xac, 0xd2,
	0x68, 0x31, 0x8c, 0x3b, 0x96, 0x79, 0x96, 0x9d, 0xb1, 0xcc, 0xb3, 0xa3, 0x03, 0xd7, 0x4d, 0x54,
	0x72, 0x23, 0x5f, 0xa1, 0x06, 0xbe, 0x14, 0x06, 0x5c, 0x93, 0xd1, 0xe1, 0x79, 0x04, 0x54, 0x5e,
	0xa1, 0x54, 0xd6, 0xbd, 0xae, 0xa1, 0x42, 0x6d, 0x3d, 0xbd, 0x5f, 0x20, 0x16, 0x25, 0xf0, 0xee,
	0xb4, 0xac, 0xdf, 0xf0, 0xbc, 0xd6, 0xc2, 0x4a, 0x54, 0xbe, 0x7e, 0x50, 0xe7, 0x95, 0xe0, 0x42,
	0x29, 0xb4, 0xdd, 0xb5, 0xe1, 0x0d, 0x68, 0x23, 0x51, 0x49, 0x55, 0x78, 0x64, 0xb7, 0x11, 0xa5,
	0xcc, 0x85, 0x05, 0xde, 0x7f, 0xb8, 0xc2, 0xf8, 0x3b, 0xa1, 0x37, 0x66, 0x28, 0xf2, 0x79, 0x6b,
	0x21, 0x83, 0xf7, 0x35, 0x18, 0x1b, 0x1e, 0xd1, 0xff, 0x15, 0xe6, 0x10, 0x91, 0x98, 0x24, 0x3b,
	0xe9, 0x5d, 0xe6, 0x2c, 0x8b, 0x4a, 0xf6, 0x1e, 0x59, 0x6b, 0x88, 0xbd, 0xc2, 0x1c, 0xb2, 0x8e,
	0x0c, 0x5f, 0xd2, 0x6b, 0xab, 0xb2, 0x36, 0x16, 0xb4, 0x0b, 0x23, 0x1a, 0xc4, 0xc3, 0x64, 0x27,
	0x4d, 0x7e, 0x23, 0xed, 0x8d, 0xb2, 0x27, 0x4e, 0xe0, 0x5e, 0xbe, 0xba, 0xda, 0x3a, 0x1d, 0x4f,
	0xbe, 0x7e, 0x3b, 0xdd, 0x67, 0xf4, 0xf0, 0xcf, 0x59, 0xa7, 0xec, 0xd7, 0xa9, 0xc7, 0x9f, 0x06,
	0xf4, 0xe6, 0xd6, 0xa5, 0xa9, 0x50, 0x19, 0x08, 0xef, 0xd1, 0xc0, 0xb7, 0x36, 0x11, 0x89, 0x87,
	0xc9, 0x95, 0x69, 0xf0, 0x63, 0x7a, 0xe9, 0x33, 0x19, 0x04, 0x24, 0xdb, 0x54, 0xc2, 0xd7, 0x74,
	0x77, 0x2b, 0x77, 0xa9, 0x8a, 0x85, 0x54, 0x16, 0x74, 0x23, 0xca, 0x68, 0xd0, 0x85, 0xb0, 0xc7,
	0xdc, 0x12, 0x58, 0xbf, 0x04, 0xf6, 0xd4, 0x2f, 0x21, 0xbb, 0xdd, 0x2a, 0xb3, 0x5e, 0xf8, 0xc2,
	0xeb, 0xc2, 0x47, 0xf4, 0x8e, 0xeb, 0xb6, 0xe8, 0x1d, 0x2f, 0x0a, 0x2d, 0x54, 0x5d, 0x0a, 0x2d,
	0xed, 0x3a, 0x1a, 0xc6, 0x24, 0x09, 0xb2, 0x3d, 0x87, 0x3c, 0xf3, 0xc4, 0xf3, 0x33, 0xe0, 0xf8,
	0x61, 0x9b, 0x01, 0xa7, 0x0f, 0xfe, 0x32, 0x03, 0x67, 0x37, 0xfd, 0x42, 0xe8, 0xad, 0xd9, 0xf6,
	0x3c, 0x73, 0x27, 0x0c, 0x3f, 0x12, 0x7a, 0x7d, 0x6e, 0x35, 0x88, 0x77, 0x1b, 0x51, 0x98, 0xb2,
	0x8b, 0x7e, 0xe6, 0x73, 0x29, 0x8f, 0x26, 0xff, 0xa4, 0x71, 0x53, 0x8d, 0xff, 0x4b, 0xc8, 0x11,
	0x99, 0x3e, 0xa6, 0x4c, 0xa2, 0x93, 0x57, 0x1a, 0x3f, 0xac, 0x2f, 0xec, 0x34, 0x0d, 0x66, 0xda,
	0x9c, 0xb4, 0x81, 0x9f, 0x90, 0x53, 0x42, 0x96, 0x97, 0xbb, 0xf0, 0x27, 0x3f, 0x03, 0x00, 0x00,
	0xff, 0xff, 0xcb, 0x5d, 0x22, 0x6c, 0xb0, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// LoadReportingServiceClient is the client API for LoadReportingService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type LoadReportingServiceClient interface {
	StreamLoadStats(ctx context.Context, opts ...grpc.CallOption) (LoadReportingService_StreamLoadStatsClient, error)
}

type loadReportingServiceClient struct {
	cc *grpc.ClientConn
}

func NewLoadReportingServiceClient(cc *grpc.ClientConn) LoadReportingServiceClient {
	return &loadReportingServiceClient{cc}
}

func (c *loadReportingServiceClient) StreamLoadStats(ctx context.Context, opts ...grpc.CallOption) (LoadReportingService_StreamLoadStatsClient, error) {
	stream, err := c.cc.NewStream(ctx, &_LoadReportingService_serviceDesc.Streams[0], "/envoy.service.load_stats.v3alpha.LoadReportingService/StreamLoadStats", opts...)
	if err != nil {
		return nil, err
	}
	x := &loadReportingServiceStreamLoadStatsClient{stream}
	return x, nil
}

type LoadReportingService_StreamLoadStatsClient interface {
	Send(*LoadStatsRequest) error
	Recv() (*LoadStatsResponse, error)
	grpc.ClientStream
}

type loadReportingServiceStreamLoadStatsClient struct {
	grpc.ClientStream
}

func (x *loadReportingServiceStreamLoadStatsClient) Send(m *LoadStatsRequest) error {
	return x.ClientStream.SendMsg(m)
}

func (x *loadReportingServiceStreamLoadStatsClient) Recv() (*LoadStatsResponse, error) {
	m := new(LoadStatsResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LoadReportingServiceServer is the server API for LoadReportingService service.
type LoadReportingServiceServer interface {
	StreamLoadStats(LoadReportingService_StreamLoadStatsServer) error
}

// UnimplementedLoadReportingServiceServer can be embedded to have forward compatible implementations.
type UnimplementedLoadReportingServiceServer struct {
}

func (*UnimplementedLoadReportingServiceServer) StreamLoadStats(srv LoadReportingService_StreamLoadStatsServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamLoadStats not implemented")
}

func RegisterLoadReportingServiceServer(s *grpc.Server, srv LoadReportingServiceServer) {
	s.RegisterService(&_LoadReportingService_serviceDesc, srv)
}

func _LoadReportingService_StreamLoadStats_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(LoadReportingServiceServer).StreamLoadStats(&loadReportingServiceStreamLoadStatsServer{stream})
}

type LoadReportingService_StreamLoadStatsServer interface {
	Send(*LoadStatsResponse) error
	Recv() (*LoadStatsRequest, error)
	grpc.ServerStream
}

type loadReportingServiceStreamLoadStatsServer struct {
	grpc.ServerStream
}

func (x *loadReportingServiceStreamLoadStatsServer) Send(m *LoadStatsResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *loadReportingServiceStreamLoadStatsServer) Recv() (*LoadStatsRequest, error) {
	m := new(LoadStatsRequest)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _LoadReportingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.load_stats.v3alpha.LoadReportingService",
	HandlerType: (*LoadReportingServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamLoadStats",
			Handler:       _LoadReportingService_StreamLoadStats_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/load_stats/v3alpha/lrs.proto",
}