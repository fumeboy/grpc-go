// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package grpc_testing

import (
	context "context"
	grpc "google.golang.org/grpc/v2"
	codes "google.golang.org/grpc/v2/codes"
	status "google.golang.org/grpc/v2/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// MetricsServiceClient is the client API for MetricsService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MetricsServiceClient interface {
	// Returns the values of all the gauges that are currently being maintained by
	// the service
	GetAllGauges(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (MetricsService_GetAllGaugesClient, error)
	// Returns the value of one gauge
	GetGauge(ctx context.Context, in *GaugeRequest, opts ...grpc.CallOption) (*GaugeResponse, error)
}

type metricsServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMetricsServiceClient(cc grpc.ClientConnInterface) MetricsServiceClient {
	return &metricsServiceClient{cc}
}

func (c *metricsServiceClient) GetAllGauges(ctx context.Context, in *EmptyMessage, opts ...grpc.CallOption) (MetricsService_GetAllGaugesClient, error) {
	stream, err := c.cc.NewStream(ctx, &MetricsService_ServiceDesc.Streams[0], "/grpc.testing.MetricsService/GetAllGauges", opts...)
	if err != nil {
		return nil, err
	}
	x := &metricsServiceGetAllGaugesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type MetricsService_GetAllGaugesClient interface {
	Recv() (*GaugeResponse, error)
	grpc.ClientStream
}

type metricsServiceGetAllGaugesClient struct {
	grpc.ClientStream
}

func (x *metricsServiceGetAllGaugesClient) Recv() (*GaugeResponse, error) {
	m := new(GaugeResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *metricsServiceClient) GetGauge(ctx context.Context, in *GaugeRequest, opts ...grpc.CallOption) (*GaugeResponse, error) {
	out := new(GaugeResponse)
	err := c.cc.Invoke(ctx, "/grpc.testing.MetricsService/GetGauge", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MetricsServiceServer is the server API for MetricsService service.
// All implementations must embed UnimplementedMetricsServiceServer
// for forward compatibility
type MetricsServiceServer interface {
	// Returns the values of all the gauges that are currently being maintained by
	// the service
	GetAllGauges(*EmptyMessage, MetricsService_GetAllGaugesServer) error
	// Returns the value of one gauge
	GetGauge(context.Context, *GaugeRequest) (*GaugeResponse, error)
	mustEmbedUnimplementedMetricsServiceServer()
}

// UnimplementedMetricsServiceServer must be embedded to have forward compatible implementations.
type UnimplementedMetricsServiceServer struct {
}

func (UnimplementedMetricsServiceServer) GetAllGauges(*EmptyMessage, MetricsService_GetAllGaugesServer) error {
	return status.Errorf(codes.Unimplemented, "method GetAllGauges not implemented")
}
func (UnimplementedMetricsServiceServer) GetGauge(context.Context, *GaugeRequest) (*GaugeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGauge not implemented")
}
func (UnimplementedMetricsServiceServer) mustEmbedUnimplementedMetricsServiceServer() {}

// UnsafeMetricsServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MetricsServiceServer will
// result in compilation errors.
type UnsafeMetricsServiceServer interface {
	mustEmbedUnimplementedMetricsServiceServer()
}

func RegisterMetricsServiceServer(s grpc.ServiceRegistrar, srv MetricsServiceServer) {
	s.RegisterService(&MetricsService_ServiceDesc, srv)
}

func _MetricsService_GetAllGauges_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(EmptyMessage)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(MetricsServiceServer).GetAllGauges(m, &metricsServiceGetAllGaugesServer{stream})
}

type MetricsService_GetAllGaugesServer interface {
	Send(*GaugeResponse) error
	grpc.ServerStream
}

type metricsServiceGetAllGaugesServer struct {
	grpc.ServerStream
}

func (x *metricsServiceGetAllGaugesServer) Send(m *GaugeResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _MetricsService_GetGauge_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GaugeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MetricsServiceServer).GetGauge(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/grpc.testing.MetricsService/GetGauge",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MetricsServiceServer).GetGauge(ctx, req.(*GaugeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// MetricsService_ServiceDesc is the grpc.ServiceDesc for MetricsService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MetricsService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "grpc.testing.MetricsService",
	HandlerType: (*MetricsServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGauge",
			Handler:    _MetricsService_GetGauge_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetAllGauges",
			Handler:       _MetricsService_GetAllGauges_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "stress/grpc_testing/metrics.proto",
}
