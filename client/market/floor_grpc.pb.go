// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: market/floor.proto

package market

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

// FloorServiceClient is the client API for FloorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type FloorServiceClient interface {
	CreateFloor(ctx context.Context, in *UpsertFloorRequest, opts ...grpc.CallOption) (*UpsertFloorResponse, error)
	UpdateFloor(ctx context.Context, in *UpsertFloorRequest, opts ...grpc.CallOption) (*UpsertFloorResponse, error)
	GetFloor(ctx context.Context, in *GetFloorRequest, opts ...grpc.CallOption) (*GetFloorResponse, error)
	GetFloorCodeAndMarketCode(ctx context.Context, in *common.FindByIdRequest, opts ...grpc.CallOption) (*GetFloorCodeAndMarketCodeResponse, error)
	GetPublishedFloor(ctx context.Context, in *GetPublishedFloorRequest, opts ...grpc.CallOption) (*GetFloorResponse, error)
	ListFloors(ctx context.Context, in *ListFloorsRequest, opts ...grpc.CallOption) (*ListFloorsResponse, error)
	ListPublishedFloors(ctx context.Context, in *common.FindByIdRequest, opts ...grpc.CallOption) (*ListFloorsResponse, error)
	DeleteFloor(ctx context.Context, in *DeleteFloorRequest, opts ...grpc.CallOption) (*common.NoContentResponse, error)
}

type floorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewFloorServiceClient(cc grpc.ClientConnInterface) FloorServiceClient {
	return &floorServiceClient{cc}
}

func (c *floorServiceClient) CreateFloor(ctx context.Context, in *UpsertFloorRequest, opts ...grpc.CallOption) (*UpsertFloorResponse, error) {
	out := new(UpsertFloorResponse)
	err := c.cc.Invoke(ctx, "/market.FloorService/CreateFloor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *floorServiceClient) UpdateFloor(ctx context.Context, in *UpsertFloorRequest, opts ...grpc.CallOption) (*UpsertFloorResponse, error) {
	out := new(UpsertFloorResponse)
	err := c.cc.Invoke(ctx, "/market.FloorService/UpdateFloor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *floorServiceClient) GetFloor(ctx context.Context, in *GetFloorRequest, opts ...grpc.CallOption) (*GetFloorResponse, error) {
	out := new(GetFloorResponse)
	err := c.cc.Invoke(ctx, "/market.FloorService/GetFloor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *floorServiceClient) GetFloorCodeAndMarketCode(ctx context.Context, in *common.FindByIdRequest, opts ...grpc.CallOption) (*GetFloorCodeAndMarketCodeResponse, error) {
	out := new(GetFloorCodeAndMarketCodeResponse)
	err := c.cc.Invoke(ctx, "/market.FloorService/GetFloorCodeAndMarketCode", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *floorServiceClient) GetPublishedFloor(ctx context.Context, in *GetPublishedFloorRequest, opts ...grpc.CallOption) (*GetFloorResponse, error) {
	out := new(GetFloorResponse)
	err := c.cc.Invoke(ctx, "/market.FloorService/GetPublishedFloor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *floorServiceClient) ListFloors(ctx context.Context, in *ListFloorsRequest, opts ...grpc.CallOption) (*ListFloorsResponse, error) {
	out := new(ListFloorsResponse)
	err := c.cc.Invoke(ctx, "/market.FloorService/ListFloors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *floorServiceClient) ListPublishedFloors(ctx context.Context, in *common.FindByIdRequest, opts ...grpc.CallOption) (*ListFloorsResponse, error) {
	out := new(ListFloorsResponse)
	err := c.cc.Invoke(ctx, "/market.FloorService/ListPublishedFloors", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *floorServiceClient) DeleteFloor(ctx context.Context, in *DeleteFloorRequest, opts ...grpc.CallOption) (*common.NoContentResponse, error) {
	out := new(common.NoContentResponse)
	err := c.cc.Invoke(ctx, "/market.FloorService/DeleteFloor", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// FloorServiceServer is the server API for FloorService service.
// All implementations must embed UnimplementedFloorServiceServer
// for forward compatibility
type FloorServiceServer interface {
	CreateFloor(context.Context, *UpsertFloorRequest) (*UpsertFloorResponse, error)
	UpdateFloor(context.Context, *UpsertFloorRequest) (*UpsertFloorResponse, error)
	GetFloor(context.Context, *GetFloorRequest) (*GetFloorResponse, error)
	GetFloorCodeAndMarketCode(context.Context, *common.FindByIdRequest) (*GetFloorCodeAndMarketCodeResponse, error)
	GetPublishedFloor(context.Context, *GetPublishedFloorRequest) (*GetFloorResponse, error)
	ListFloors(context.Context, *ListFloorsRequest) (*ListFloorsResponse, error)
	ListPublishedFloors(context.Context, *common.FindByIdRequest) (*ListFloorsResponse, error)
	DeleteFloor(context.Context, *DeleteFloorRequest) (*common.NoContentResponse, error)
	mustEmbedUnimplementedFloorServiceServer()
}

// UnimplementedFloorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedFloorServiceServer struct {
}

func (UnimplementedFloorServiceServer) CreateFloor(context.Context, *UpsertFloorRequest) (*UpsertFloorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateFloor not implemented")
}
func (UnimplementedFloorServiceServer) UpdateFloor(context.Context, *UpsertFloorRequest) (*UpsertFloorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateFloor not implemented")
}
func (UnimplementedFloorServiceServer) GetFloor(context.Context, *GetFloorRequest) (*GetFloorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFloor not implemented")
}
func (UnimplementedFloorServiceServer) GetFloorCodeAndMarketCode(context.Context, *common.FindByIdRequest) (*GetFloorCodeAndMarketCodeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetFloorCodeAndMarketCode not implemented")
}
func (UnimplementedFloorServiceServer) GetPublishedFloor(context.Context, *GetPublishedFloorRequest) (*GetFloorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPublishedFloor not implemented")
}
func (UnimplementedFloorServiceServer) ListFloors(context.Context, *ListFloorsRequest) (*ListFloorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListFloors not implemented")
}
func (UnimplementedFloorServiceServer) ListPublishedFloors(context.Context, *common.FindByIdRequest) (*ListFloorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListPublishedFloors not implemented")
}
func (UnimplementedFloorServiceServer) DeleteFloor(context.Context, *DeleteFloorRequest) (*common.NoContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteFloor not implemented")
}
func (UnimplementedFloorServiceServer) mustEmbedUnimplementedFloorServiceServer() {}

// UnsafeFloorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to FloorServiceServer will
// result in compilation errors.
type UnsafeFloorServiceServer interface {
	mustEmbedUnimplementedFloorServiceServer()
}

func RegisterFloorServiceServer(s grpc.ServiceRegistrar, srv FloorServiceServer) {
	s.RegisterService(&FloorService_ServiceDesc, srv)
}

func _FloorService_CreateFloor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertFloorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FloorServiceServer).CreateFloor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/market.FloorService/CreateFloor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FloorServiceServer).CreateFloor(ctx, req.(*UpsertFloorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FloorService_UpdateFloor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpsertFloorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FloorServiceServer).UpdateFloor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/market.FloorService/UpdateFloor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FloorServiceServer).UpdateFloor(ctx, req.(*UpsertFloorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FloorService_GetFloor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetFloorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FloorServiceServer).GetFloor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/market.FloorService/GetFloor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FloorServiceServer).GetFloor(ctx, req.(*GetFloorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FloorService_GetFloorCodeAndMarketCode_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.FindByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FloorServiceServer).GetFloorCodeAndMarketCode(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/market.FloorService/GetFloorCodeAndMarketCode",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FloorServiceServer).GetFloorCodeAndMarketCode(ctx, req.(*common.FindByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FloorService_GetPublishedFloor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPublishedFloorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FloorServiceServer).GetPublishedFloor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/market.FloorService/GetPublishedFloor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FloorServiceServer).GetPublishedFloor(ctx, req.(*GetPublishedFloorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FloorService_ListFloors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListFloorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FloorServiceServer).ListFloors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/market.FloorService/ListFloors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FloorServiceServer).ListFloors(ctx, req.(*ListFloorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FloorService_ListPublishedFloors_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.FindByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FloorServiceServer).ListPublishedFloors(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/market.FloorService/ListPublishedFloors",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FloorServiceServer).ListPublishedFloors(ctx, req.(*common.FindByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _FloorService_DeleteFloor_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteFloorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(FloorServiceServer).DeleteFloor(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/market.FloorService/DeleteFloor",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(FloorServiceServer).DeleteFloor(ctx, req.(*DeleteFloorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// FloorService_ServiceDesc is the grpc.ServiceDesc for FloorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var FloorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "market.FloorService",
	HandlerType: (*FloorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateFloor",
			Handler:    _FloorService_CreateFloor_Handler,
		},
		{
			MethodName: "UpdateFloor",
			Handler:    _FloorService_UpdateFloor_Handler,
		},
		{
			MethodName: "GetFloor",
			Handler:    _FloorService_GetFloor_Handler,
		},
		{
			MethodName: "GetFloorCodeAndMarketCode",
			Handler:    _FloorService_GetFloorCodeAndMarketCode_Handler,
		},
		{
			MethodName: "GetPublishedFloor",
			Handler:    _FloorService_GetPublishedFloor_Handler,
		},
		{
			MethodName: "ListFloors",
			Handler:    _FloorService_ListFloors_Handler,
		},
		{
			MethodName: "ListPublishedFloors",
			Handler:    _FloorService_ListPublishedFloors_Handler,
		},
		{
			MethodName: "DeleteFloor",
			Handler:    _FloorService_DeleteFloor_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "market/floor.proto",
}