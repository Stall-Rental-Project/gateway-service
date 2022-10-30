// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: market/location.proto

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

// LocationServiceClient is the client API for LocationService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LocationServiceClient interface {
	ListProvinces(ctx context.Context, in *ListProvinceRequest, opts ...grpc.CallOption) (*ListProvinceResponse, error)
	ListCities(ctx context.Context, in *ListCityRequest, opts ...grpc.CallOption) (*ListCityResponse, error)
	ListWards(ctx context.Context, in *ListWardRequest, opts ...grpc.CallOption) (*ListWardResponse, error)
	ListLocations(ctx context.Context, in *common.FindByIdsRequest, opts ...grpc.CallOption) (*common.ListResponse, error)
	GetLocation(ctx context.Context, in *GetLocationRequest, opts ...grpc.CallOption) (*GetLocationResponse, error)
}

type locationServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLocationServiceClient(cc grpc.ClientConnInterface) LocationServiceClient {
	return &locationServiceClient{cc}
}

func (c *locationServiceClient) ListProvinces(ctx context.Context, in *ListProvinceRequest, opts ...grpc.CallOption) (*ListProvinceResponse, error) {
	out := new(ListProvinceResponse)
	err := c.cc.Invoke(ctx, "/market.LocationService/ListProvinces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationServiceClient) ListCities(ctx context.Context, in *ListCityRequest, opts ...grpc.CallOption) (*ListCityResponse, error) {
	out := new(ListCityResponse)
	err := c.cc.Invoke(ctx, "/market.LocationService/ListCities", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationServiceClient) ListWards(ctx context.Context, in *ListWardRequest, opts ...grpc.CallOption) (*ListWardResponse, error) {
	out := new(ListWardResponse)
	err := c.cc.Invoke(ctx, "/market.LocationService/ListWards", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationServiceClient) ListLocations(ctx context.Context, in *common.FindByIdsRequest, opts ...grpc.CallOption) (*common.ListResponse, error) {
	out := new(common.ListResponse)
	err := c.cc.Invoke(ctx, "/market.LocationService/ListLocations", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *locationServiceClient) GetLocation(ctx context.Context, in *GetLocationRequest, opts ...grpc.CallOption) (*GetLocationResponse, error) {
	out := new(GetLocationResponse)
	err := c.cc.Invoke(ctx, "/market.LocationService/GetLocation", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// LocationServiceServer is the server API for LocationService service.
// All implementations must embed UnimplementedLocationServiceServer
// for forward compatibility
type LocationServiceServer interface {
	ListProvinces(context.Context, *ListProvinceRequest) (*ListProvinceResponse, error)
	ListCities(context.Context, *ListCityRequest) (*ListCityResponse, error)
	ListWards(context.Context, *ListWardRequest) (*ListWardResponse, error)
	ListLocations(context.Context, *common.FindByIdsRequest) (*common.ListResponse, error)
	GetLocation(context.Context, *GetLocationRequest) (*GetLocationResponse, error)
	mustEmbedUnimplementedLocationServiceServer()
}

// UnimplementedLocationServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLocationServiceServer struct {
}

func (UnimplementedLocationServiceServer) ListProvinces(context.Context, *ListProvinceRequest) (*ListProvinceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProvinces not implemented")
}
func (UnimplementedLocationServiceServer) ListCities(context.Context, *ListCityRequest) (*ListCityResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListCities not implemented")
}
func (UnimplementedLocationServiceServer) ListWards(context.Context, *ListWardRequest) (*ListWardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWards not implemented")
}
func (UnimplementedLocationServiceServer) ListLocations(context.Context, *common.FindByIdsRequest) (*common.ListResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListLocations not implemented")
}
func (UnimplementedLocationServiceServer) GetLocation(context.Context, *GetLocationRequest) (*GetLocationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetLocation not implemented")
}
func (UnimplementedLocationServiceServer) mustEmbedUnimplementedLocationServiceServer() {}

// UnsafeLocationServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LocationServiceServer will
// result in compilation errors.
type UnsafeLocationServiceServer interface {
	mustEmbedUnimplementedLocationServiceServer()
}

func RegisterLocationServiceServer(s grpc.ServiceRegistrar, srv LocationServiceServer) {
	s.RegisterService(&LocationService_ServiceDesc, srv)
}

func _LocationService_ListProvinces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProvinceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationServiceServer).ListProvinces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/market.LocationService/ListProvinces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationServiceServer).ListProvinces(ctx, req.(*ListProvinceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocationService_ListCities_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListCityRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationServiceServer).ListCities(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/market.LocationService/ListCities",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationServiceServer).ListCities(ctx, req.(*ListCityRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocationService_ListWards_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListWardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationServiceServer).ListWards(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/market.LocationService/ListWards",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationServiceServer).ListWards(ctx, req.(*ListWardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocationService_ListLocations_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.FindByIdsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationServiceServer).ListLocations(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/market.LocationService/ListLocations",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationServiceServer).ListLocations(ctx, req.(*common.FindByIdsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _LocationService_GetLocation_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetLocationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LocationServiceServer).GetLocation(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/market.LocationService/GetLocation",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LocationServiceServer).GetLocation(ctx, req.(*GetLocationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// LocationService_ServiceDesc is the grpc.ServiceDesc for LocationService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LocationService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "market.LocationService",
	HandlerType: (*LocationServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListProvinces",
			Handler:    _LocationService_ListProvinces_Handler,
		},
		{
			MethodName: "ListCities",
			Handler:    _LocationService_ListCities_Handler,
		},
		{
			MethodName: "ListWards",
			Handler:    _LocationService_ListWards_Handler,
		},
		{
			MethodName: "ListLocations",
			Handler:    _LocationService_ListLocations_Handler,
		},
		{
			MethodName: "GetLocation",
			Handler:    _LocationService_GetLocation_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "market/location.proto",
}
