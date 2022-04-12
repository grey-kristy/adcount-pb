// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package pb

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

// ADCounterClient is the client API for ADCounter service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ADCounterClient interface {
	UpdateCounter(ctx context.Context, in *CounterKey, opts ...grpc.CallOption) (*CounterValue, error)
}

type aDCounterClient struct {
	cc grpc.ClientConnInterface
}

func NewADCounterClient(cc grpc.ClientConnInterface) ADCounterClient {
	return &aDCounterClient{cc}
}

func (c *aDCounterClient) UpdateCounter(ctx context.Context, in *CounterKey, opts ...grpc.CallOption) (*CounterValue, error) {
	out := new(CounterValue)
	err := c.cc.Invoke(ctx, "/counter.ADCounter/UpdateCounter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ADCounterServer is the server API for ADCounter service.
// All implementations must embed UnimplementedADCounterServer
// for forward compatibility
type ADCounterServer interface {
	UpdateCounter(context.Context, *CounterKey) (*CounterValue, error)
	mustEmbedUnimplementedADCounterServer()
}

// UnimplementedADCounterServer must be embedded to have forward compatible implementations.
type UnimplementedADCounterServer struct {
}

func (UnimplementedADCounterServer) UpdateCounter(context.Context, *CounterKey) (*CounterValue, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateCounter not implemented")
}
func (UnimplementedADCounterServer) mustEmbedUnimplementedADCounterServer() {}

// UnsafeADCounterServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ADCounterServer will
// result in compilation errors.
type UnsafeADCounterServer interface {
	mustEmbedUnimplementedADCounterServer()
}

func RegisterADCounterServer(s grpc.ServiceRegistrar, srv ADCounterServer) {
	s.RegisterService(&ADCounter_ServiceDesc, srv)
}

func _ADCounter_UpdateCounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CounterKey)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ADCounterServer).UpdateCounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/counter.ADCounter/UpdateCounter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ADCounterServer).UpdateCounter(ctx, req.(*CounterKey))
	}
	return interceptor(ctx, in, info, handler)
}

// ADCounter_ServiceDesc is the grpc.ServiceDesc for ADCounter service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ADCounter_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "counter.ADCounter",
	HandlerType: (*ADCounterServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateCounter",
			Handler:    _ADCounter_UpdateCounter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "pb/adcounter.proto",
}
