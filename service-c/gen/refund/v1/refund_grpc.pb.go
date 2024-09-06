// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.3
// source: service-c/proto/refund/refund.proto

package v1

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
	RefundService_GetRefund_FullMethodName = "/ggoh.refund.RefundService/GetRefund"
)

// RefundServiceClient is the client API for RefundService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RefundServiceClient interface {
	// Retrieves Refund details
	GetRefund(ctx context.Context, in *RefundRequest, opts ...grpc.CallOption) (*RefundResponse, error)
}

type refundServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewRefundServiceClient(cc grpc.ClientConnInterface) RefundServiceClient {
	return &refundServiceClient{cc}
}

func (c *refundServiceClient) GetRefund(ctx context.Context, in *RefundRequest, opts ...grpc.CallOption) (*RefundResponse, error) {
	out := new(RefundResponse)
	err := c.cc.Invoke(ctx, RefundService_GetRefund_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RefundServiceServer is the server API for RefundService service.
// All implementations must embed UnimplementedRefundServiceServer
// for forward compatibility
type RefundServiceServer interface {
	// Retrieves Refund details
	GetRefund(context.Context, *RefundRequest) (*RefundResponse, error)
	mustEmbedUnimplementedRefundServiceServer()
}

// UnimplementedRefundServiceServer must be embedded to have forward compatible implementations.
type UnimplementedRefundServiceServer struct {
}

func (UnimplementedRefundServiceServer) GetRefund(context.Context, *RefundRequest) (*RefundResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRefund not implemented")
}
func (UnimplementedRefundServiceServer) mustEmbedUnimplementedRefundServiceServer() {}

// UnsafeRefundServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RefundServiceServer will
// result in compilation errors.
type UnsafeRefundServiceServer interface {
	mustEmbedUnimplementedRefundServiceServer()
}

func RegisterRefundServiceServer(s grpc.ServiceRegistrar, srv RefundServiceServer) {
	s.RegisterService(&RefundService_ServiceDesc, srv)
}

func _RefundService_GetRefund_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefundRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RefundServiceServer).GetRefund(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RefundService_GetRefund_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RefundServiceServer).GetRefund(ctx, req.(*RefundRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RefundService_ServiceDesc is the grpc.ServiceDesc for RefundService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RefundService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ggoh.refund.RefundService",
	HandlerType: (*RefundServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetRefund",
			Handler:    _RefundService_GetRefund_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service-c/proto/refund/refund.proto",
}
