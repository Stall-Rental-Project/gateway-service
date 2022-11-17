// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: rental/lease.proto

package rental

import (
	context "context"
	common "gateway-service/client/common"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LeaseServiceClient is the client API for LeaseService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LeaseServiceClient interface {
	ListLeases(ctx context.Context, in *ListLeasesRequest, opts ...grpc.CallOption) (*common.PageResponse, error)
	GetLease(ctx context.Context, in *common.FindByIdRequest, opts ...grpc.CallOption) (*GetApplicationResponse, error)
}

type leaseServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLeaseServiceClient(cc grpc.ClientConnInterface) LeaseServiceClient {
	return &leaseServiceClient{cc}
}

func (c *leaseServiceClient) ListLeases(ctx context.Context, in *ListLeasesRequest, opts ...grpc.CallOption) (*common.PageResponse, error) {
	out := new(common.PageResponse)
	err := c.cc.Invoke(ctx, "/rental.LeaseService/ListLeases", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *leaseServiceClient) GetLease(ctx context.Context, in *common.FindByIdRequest, opts ...grpc.CallOption) (*GetApplicationResponse, error) {
	out := new(GetApplicationResponse)
	err := c.cc.Invoke(ctx, "/rental.LeaseService/GetLease", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LeaseServiceServer is the server API for LeaseService service.
// All implementations must embed UnimplementedLeaseServiceServer
// for forward compatibility
type LeaseServiceServer interface {
	ListLeases(context.Context, *ListLeasesRequest) (*common.PageResponse, error)
	GetLease(context.Context, *common.FindByIdRequest) (*GetApplicationResponse, error)
	mustEmbedUnimplementedLeaseServiceServer()
}

// UnimplementedLeaseServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLeaseServiceServer struct {
}

func (UnimplementedLeaseServiceServer) ListLeases(context.Context, *ListLeasesRequest) (*common.PageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLeases not implemented")
}
func (UnimplementedLeaseServiceServer) GetLease(context.Context, *common.FindByIdRequest) (*GetApplicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLease not implemented")
}
func (UnimplementedLeaseServiceServer) mustEmbedUnimplementedLeaseServiceServer() {}

// UnsafeLeaseServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LeaseServiceServer will
// result in compilation errors.
type UnsafeLeaseServiceServer interface {
	mustEmbedUnimplementedLeaseServiceServer()
}

func RegisterLeaseServiceServer(s grpc.ServiceRegistrar, srv LeaseServiceServer) {
	s.RegisterService(&LeaseService_ServiceDesc, srv)
}

func _LeaseService_ListLeases_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListLeasesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaseServiceServer).ListLeases(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rental.LeaseService/ListLeases",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaseServiceServer).ListLeases(ctx, req.(*ListLeasesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LeaseService_GetLease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.FindByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LeaseServiceServer).GetLease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rental.LeaseService/GetLease",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LeaseServiceServer).GetLease(ctx, req.(*common.FindByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LeaseService_ServiceDesc is the grpc.ServiceDesc for LeaseService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LeaseService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rental.LeaseService",
	HandlerType: (*LeaseServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListLeases",
			Handler:    _LeaseService_ListLeases_Handler,
		},
		{
			MethodName: "GetLease",
			Handler:    _LeaseService_GetLease_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rental/lease.proto",
}