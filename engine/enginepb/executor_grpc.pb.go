// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package enginepb

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

// ExecutorClient is the client API for Executor service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ExecutorClient interface {
	PreDispatchTask(ctx context.Context, in *PreDispatchTaskRequest, opts ...grpc.CallOption) (*PreDispatchTaskResponse, error)
	ConfirmDispatchTask(ctx context.Context, in *ConfirmDispatchTaskRequest, opts ...grpc.CallOption) (*ConfirmDispatchTaskResponse, error)
}

type executorClient struct {
	cc grpc.ClientConnInterface
}

func NewExecutorClient(cc grpc.ClientConnInterface) ExecutorClient {
	return &executorClient{cc}
}

func (c *executorClient) PreDispatchTask(ctx context.Context, in *PreDispatchTaskRequest, opts ...grpc.CallOption) (*PreDispatchTaskResponse, error) {
	out := new(PreDispatchTaskResponse)
	err := c.cc.Invoke(ctx, "/enginepb.Executor/PreDispatchTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *executorClient) ConfirmDispatchTask(ctx context.Context, in *ConfirmDispatchTaskRequest, opts ...grpc.CallOption) (*ConfirmDispatchTaskResponse, error) {
	out := new(ConfirmDispatchTaskResponse)
	err := c.cc.Invoke(ctx, "/enginepb.Executor/ConfirmDispatchTask", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ExecutorServer is the server API for Executor service.
// All implementations should embed UnimplementedExecutorServer
// for forward compatibility
type ExecutorServer interface {
	PreDispatchTask(context.Context, *PreDispatchTaskRequest) (*PreDispatchTaskResponse, error)
	ConfirmDispatchTask(context.Context, *ConfirmDispatchTaskRequest) (*ConfirmDispatchTaskResponse, error)
}

// UnimplementedExecutorServer should be embedded to have forward compatible implementations.
type UnimplementedExecutorServer struct {
}

func (UnimplementedExecutorServer) PreDispatchTask(context.Context, *PreDispatchTaskRequest) (*PreDispatchTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PreDispatchTask not implemented")
}
func (UnimplementedExecutorServer) ConfirmDispatchTask(context.Context, *ConfirmDispatchTaskRequest) (*ConfirmDispatchTaskResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmDispatchTask not implemented")
}

// UnsafeExecutorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ExecutorServer will
// result in compilation errors.
type UnsafeExecutorServer interface {
	mustEmbedUnimplementedExecutorServer()
}

func RegisterExecutorServer(s grpc.ServiceRegistrar, srv ExecutorServer) {
	s.RegisterService(&Executor_ServiceDesc, srv)
}

func _Executor_PreDispatchTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PreDispatchTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).PreDispatchTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/enginepb.Executor/PreDispatchTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).PreDispatchTask(ctx, req.(*PreDispatchTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Executor_ConfirmDispatchTask_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmDispatchTaskRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ExecutorServer).ConfirmDispatchTask(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/enginepb.Executor/ConfirmDispatchTask",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ExecutorServer).ConfirmDispatchTask(ctx, req.(*ConfirmDispatchTaskRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Executor_ServiceDesc is the grpc.ServiceDesc for Executor service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Executor_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "enginepb.Executor",
	HandlerType: (*ExecutorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PreDispatchTask",
			Handler:    _Executor_PreDispatchTask_Handler,
		},
		{
			MethodName: "ConfirmDispatchTask",
			Handler:    _Executor_ConfirmDispatchTask_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "engine/proto/executor.proto",
}

// BrokerServiceClient is the client API for BrokerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BrokerServiceClient interface {
	RemoveResource(ctx context.Context, in *RemoveLocalResourceRequest, opts ...grpc.CallOption) (*RemoveLocalResourceResponse, error)
}

type brokerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBrokerServiceClient(cc grpc.ClientConnInterface) BrokerServiceClient {
	return &brokerServiceClient{cc}
}

func (c *brokerServiceClient) RemoveResource(ctx context.Context, in *RemoveLocalResourceRequest, opts ...grpc.CallOption) (*RemoveLocalResourceResponse, error) {
	out := new(RemoveLocalResourceResponse)
	err := c.cc.Invoke(ctx, "/enginepb.BrokerService/RemoveResource", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BrokerServiceServer is the server API for BrokerService service.
// All implementations should embed UnimplementedBrokerServiceServer
// for forward compatibility
type BrokerServiceServer interface {
	RemoveResource(context.Context, *RemoveLocalResourceRequest) (*RemoveLocalResourceResponse, error)
}

// UnimplementedBrokerServiceServer should be embedded to have forward compatible implementations.
type UnimplementedBrokerServiceServer struct {
}

func (UnimplementedBrokerServiceServer) RemoveResource(context.Context, *RemoveLocalResourceRequest) (*RemoveLocalResourceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveResource not implemented")
}

// UnsafeBrokerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BrokerServiceServer will
// result in compilation errors.
type UnsafeBrokerServiceServer interface {
	mustEmbedUnimplementedBrokerServiceServer()
}

func RegisterBrokerServiceServer(s grpc.ServiceRegistrar, srv BrokerServiceServer) {
	s.RegisterService(&BrokerService_ServiceDesc, srv)
}

func _BrokerService_RemoveResource_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveLocalResourceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BrokerServiceServer).RemoveResource(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/enginepb.BrokerService/RemoveResource",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BrokerServiceServer).RemoveResource(ctx, req.(*RemoveLocalResourceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BrokerService_ServiceDesc is the grpc.ServiceDesc for BrokerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BrokerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "enginepb.BrokerService",
	HandlerType: (*BrokerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RemoveResource",
			Handler:    _BrokerService_RemoveResource_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "engine/proto/executor.proto",
}
