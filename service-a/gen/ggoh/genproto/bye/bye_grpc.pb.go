// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: service-a/proto/bye/bye.proto

package bye

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

const (
	Byer_SayBye_FullMethodName = "/ggoh.bye.Byer/SayBye"
)

// ByerClient is the client API for Byer service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ByerClient interface {
	// Sends a greeting
	SayBye(ctx context.Context, in *ByeRequest, opts ...grpc.CallOption) (*ByeReply, error)
}

type byerClient struct {
	cc grpc.ClientConnInterface
}

func NewByerClient(cc grpc.ClientConnInterface) ByerClient {
	return &byerClient{cc}
}

func (c *byerClient) SayBye(ctx context.Context, in *ByeRequest, opts ...grpc.CallOption) (*ByeReply, error) {
	out := new(ByeReply)
	err := c.cc.Invoke(ctx, Byer_SayBye_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ByerServer is the server API for Byer service.
// All implementations must embed UnimplementedByerServer
// for forward compatibility
type ByerServer interface {
	// Sends a greeting
	SayBye(context.Context, *ByeRequest) (*ByeReply, error)
	mustEmbedUnimplementedByerServer()
}

// UnimplementedByerServer must be embedded to have forward compatible implementations.
type UnimplementedByerServer struct {
}

func (UnimplementedByerServer) SayBye(context.Context, *ByeRequest) (*ByeReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayBye not implemented")
}
func (UnimplementedByerServer) mustEmbedUnimplementedByerServer() {}

// UnsafeByerServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ByerServer will
// result in compilation errors.
type UnsafeByerServer interface {
	mustEmbedUnimplementedByerServer()
}

func RegisterByerServer(s grpc.ServiceRegistrar, srv ByerServer) {
	s.RegisterService(&Byer_ServiceDesc, srv)
}

func _Byer_SayBye_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ByeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ByerServer).SayBye(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Byer_SayBye_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ByerServer).SayBye(ctx, req.(*ByeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Byer_ServiceDesc is the grpc.ServiceDesc for Byer service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Byer_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ggoh.bye.Byer",
	HandlerType: (*ByerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayBye",
			Handler:    _Byer_SayBye_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service-a/proto/bye/bye.proto",
}
