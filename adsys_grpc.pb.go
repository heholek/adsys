// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package adsys

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ServiceClient is the client API for Service service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ServiceClient interface {
	Cat(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Service_CatClient, error)
	Version(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Service_VersionClient, error)
	Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (Service_StopClient, error)
	UpdatePolicy(ctx context.Context, in *UpdatePolicyRequest, opts ...grpc.CallOption) (Service_UpdatePolicyClient, error)
	DumpPolicies(ctx context.Context, in *DumpPoliciesRequest, opts ...grpc.CallOption) (Service_DumpPoliciesClient, error)
}

type serviceClient struct {
	cc grpc.ClientConnInterface
}

func NewServiceClient(cc grpc.ClientConnInterface) ServiceClient {
	return &serviceClient{cc}
}

func (c *serviceClient) Cat(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Service_CatClient, error) {
	stream, err := c.cc.NewStream(ctx, &Service_ServiceDesc.Streams[0], "/service/Cat", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceCatClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Service_CatClient interface {
	Recv() (*StringResponse, error)
	grpc.ClientStream
}

type serviceCatClient struct {
	grpc.ClientStream
}

func (x *serviceCatClient) Recv() (*StringResponse, error) {
	m := new(StringResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceClient) Version(ctx context.Context, in *Empty, opts ...grpc.CallOption) (Service_VersionClient, error) {
	stream, err := c.cc.NewStream(ctx, &Service_ServiceDesc.Streams[1], "/service/Version", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceVersionClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Service_VersionClient interface {
	Recv() (*StringResponse, error)
	grpc.ClientStream
}

type serviceVersionClient struct {
	grpc.ClientStream
}

func (x *serviceVersionClient) Recv() (*StringResponse, error) {
	m := new(StringResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceClient) Stop(ctx context.Context, in *StopRequest, opts ...grpc.CallOption) (Service_StopClient, error) {
	stream, err := c.cc.NewStream(ctx, &Service_ServiceDesc.Streams[2], "/service/Stop", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceStopClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Service_StopClient interface {
	Recv() (*Empty, error)
	grpc.ClientStream
}

type serviceStopClient struct {
	grpc.ClientStream
}

func (x *serviceStopClient) Recv() (*Empty, error) {
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceClient) UpdatePolicy(ctx context.Context, in *UpdatePolicyRequest, opts ...grpc.CallOption) (Service_UpdatePolicyClient, error) {
	stream, err := c.cc.NewStream(ctx, &Service_ServiceDesc.Streams[3], "/service/UpdatePolicy", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceUpdatePolicyClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Service_UpdatePolicyClient interface {
	Recv() (*Empty, error)
	grpc.ClientStream
}

type serviceUpdatePolicyClient struct {
	grpc.ClientStream
}

func (x *serviceUpdatePolicyClient) Recv() (*Empty, error) {
	m := new(Empty)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *serviceClient) DumpPolicies(ctx context.Context, in *DumpPoliciesRequest, opts ...grpc.CallOption) (Service_DumpPoliciesClient, error) {
	stream, err := c.cc.NewStream(ctx, &Service_ServiceDesc.Streams[4], "/service/DumpPolicies", opts...)
	if err != nil {
		return nil, err
	}
	x := &serviceDumpPoliciesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type Service_DumpPoliciesClient interface {
	Recv() (*StringResponse, error)
	grpc.ClientStream
}

type serviceDumpPoliciesClient struct {
	grpc.ClientStream
}

func (x *serviceDumpPoliciesClient) Recv() (*StringResponse, error) {
	m := new(StringResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ServiceServer is the server API for Service service.
// All implementations must embed UnimplementedServiceServer
// for forward compatibility
type ServiceServer interface {
	Cat(*Empty, Service_CatServer) error
	Version(*Empty, Service_VersionServer) error
	Stop(*StopRequest, Service_StopServer) error
	UpdatePolicy(*UpdatePolicyRequest, Service_UpdatePolicyServer) error
	DumpPolicies(*DumpPoliciesRequest, Service_DumpPoliciesServer) error
	mustEmbedUnimplementedServiceServer()
}

// UnimplementedServiceServer must be embedded to have forward compatible implementations.
type UnimplementedServiceServer struct {
}

func (UnimplementedServiceServer) Cat(*Empty, Service_CatServer) error {
	return status.Errorf(codes.Unimplemented, "method Cat not implemented")
}
func (UnimplementedServiceServer) Version(*Empty, Service_VersionServer) error {
	return status.Errorf(codes.Unimplemented, "method Version not implemented")
}
func (UnimplementedServiceServer) Stop(*StopRequest, Service_StopServer) error {
	return status.Errorf(codes.Unimplemented, "method Stop not implemented")
}
func (UnimplementedServiceServer) UpdatePolicy(*UpdatePolicyRequest, Service_UpdatePolicyServer) error {
	return status.Errorf(codes.Unimplemented, "method UpdatePolicy not implemented")
}
func (UnimplementedServiceServer) DumpPolicies(*DumpPoliciesRequest, Service_DumpPoliciesServer) error {
	return status.Errorf(codes.Unimplemented, "method DumpPolicies not implemented")
}
func (UnimplementedServiceServer) mustEmbedUnimplementedServiceServer() {}

// UnsafeServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ServiceServer will
// result in compilation errors.
type UnsafeServiceServer interface {
	mustEmbedUnimplementedServiceServer()
}

func RegisterServiceServer(s grpc.ServiceRegistrar, srv ServiceServer) {
	s.RegisterService(&Service_ServiceDesc, srv)
}

func _Service_Cat_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceServer).Cat(m, &serviceCatServer{stream})
}

type Service_CatServer interface {
	Send(*StringResponse) error
	grpc.ServerStream
}

type serviceCatServer struct {
	grpc.ServerStream
}

func (x *serviceCatServer) Send(m *StringResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Service_Version_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceServer).Version(m, &serviceVersionServer{stream})
}

type Service_VersionServer interface {
	Send(*StringResponse) error
	grpc.ServerStream
}

type serviceVersionServer struct {
	grpc.ServerStream
}

func (x *serviceVersionServer) Send(m *StringResponse) error {
	return x.ServerStream.SendMsg(m)
}

func _Service_Stop_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(StopRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceServer).Stop(m, &serviceStopServer{stream})
}

type Service_StopServer interface {
	Send(*Empty) error
	grpc.ServerStream
}

type serviceStopServer struct {
	grpc.ServerStream
}

func (x *serviceStopServer) Send(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func _Service_UpdatePolicy_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(UpdatePolicyRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceServer).UpdatePolicy(m, &serviceUpdatePolicyServer{stream})
}

type Service_UpdatePolicyServer interface {
	Send(*Empty) error
	grpc.ServerStream
}

type serviceUpdatePolicyServer struct {
	grpc.ServerStream
}

func (x *serviceUpdatePolicyServer) Send(m *Empty) error {
	return x.ServerStream.SendMsg(m)
}

func _Service_DumpPolicies_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(DumpPoliciesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ServiceServer).DumpPolicies(m, &serviceDumpPoliciesServer{stream})
}

type Service_DumpPoliciesServer interface {
	Send(*StringResponse) error
	grpc.ServerStream
}

type serviceDumpPoliciesServer struct {
	grpc.ServerStream
}

func (x *serviceDumpPoliciesServer) Send(m *StringResponse) error {
	return x.ServerStream.SendMsg(m)
}

// Service_ServiceDesc is the grpc.ServiceDesc for Service service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Service_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service",
	HandlerType: (*ServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Cat",
			Handler:       _Service_Cat_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Version",
			Handler:       _Service_Version_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "Stop",
			Handler:       _Service_Stop_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "UpdatePolicy",
			Handler:       _Service_UpdatePolicy_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "DumpPolicies",
			Handler:       _Service_DumpPolicies_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "adsys.proto",
}
