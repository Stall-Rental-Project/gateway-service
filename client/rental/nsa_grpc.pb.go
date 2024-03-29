// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.8
// source: rental/nsa.proto

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

// NSAServiceClient is the client API for NSAService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NSAServiceClient interface {
	// unary
	SubmitApplication(ctx context.Context, in *SubmitApplicationRequest, opts ...grpc.CallOption) (*GetApplicationResponse, error)
	SubmitApplicationDocs(ctx context.Context, in *SubmitApplicationDocsRequest, opts ...grpc.CallOption) (*common.NoContentResponse, error)
	SubmitApplicationPayment(ctx context.Context, in *SubmitApplicationPaymentRequest, opts ...grpc.CallOption) (*common.NoContentResponse, error)
	CheckExistApplication(ctx context.Context, in *CheckExistApplicationRequest, opts ...grpc.CallOption) (*common.BooleanResponse, error)
	CheckExistApplications(ctx context.Context, in *CheckExistApplicationsRequest, opts ...grpc.CallOption) (*CheckExistApplicationsResponse, error)
	GetApplication(ctx context.Context, in *common.FindByIdRequest, opts ...grpc.CallOption) (*GetApplicationResponse, error)
	UpdateApplication(ctx context.Context, in *SubmitApplicationRequest, opts ...grpc.CallOption) (*GetApplicationResponse, error)
	ConfirmApplication(ctx context.Context, in *ConfirmApplicationRequest, opts ...grpc.CallOption) (*common.NoContentResponse, error)
}

type nSAServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewNSAServiceClient(cc grpc.ClientConnInterface) NSAServiceClient {
	return &nSAServiceClient{cc}
}

func (c *nSAServiceClient) SubmitApplication(ctx context.Context, in *SubmitApplicationRequest, opts ...grpc.CallOption) (*GetApplicationResponse, error) {
	out := new(GetApplicationResponse)
	err := c.cc.Invoke(ctx, "/rental.NSAService/SubmitApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSAServiceClient) SubmitApplicationDocs(ctx context.Context, in *SubmitApplicationDocsRequest, opts ...grpc.CallOption) (*common.NoContentResponse, error) {
	out := new(common.NoContentResponse)
	err := c.cc.Invoke(ctx, "/rental.NSAService/SubmitApplicationDocs", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSAServiceClient) SubmitApplicationPayment(ctx context.Context, in *SubmitApplicationPaymentRequest, opts ...grpc.CallOption) (*common.NoContentResponse, error) {
	out := new(common.NoContentResponse)
	err := c.cc.Invoke(ctx, "/rental.NSAService/SubmitApplicationPayment", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSAServiceClient) CheckExistApplication(ctx context.Context, in *CheckExistApplicationRequest, opts ...grpc.CallOption) (*common.BooleanResponse, error) {
	out := new(common.BooleanResponse)
	err := c.cc.Invoke(ctx, "/rental.NSAService/CheckExistApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSAServiceClient) CheckExistApplications(ctx context.Context, in *CheckExistApplicationsRequest, opts ...grpc.CallOption) (*CheckExistApplicationsResponse, error) {
	out := new(CheckExistApplicationsResponse)
	err := c.cc.Invoke(ctx, "/rental.NSAService/CheckExistApplications", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSAServiceClient) GetApplication(ctx context.Context, in *common.FindByIdRequest, opts ...grpc.CallOption) (*GetApplicationResponse, error) {
	out := new(GetApplicationResponse)
	err := c.cc.Invoke(ctx, "/rental.NSAService/GetApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSAServiceClient) UpdateApplication(ctx context.Context, in *SubmitApplicationRequest, opts ...grpc.CallOption) (*GetApplicationResponse, error) {
	out := new(GetApplicationResponse)
	err := c.cc.Invoke(ctx, "/rental.NSAService/UpdateApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *nSAServiceClient) ConfirmApplication(ctx context.Context, in *ConfirmApplicationRequest, opts ...grpc.CallOption) (*common.NoContentResponse, error) {
	out := new(common.NoContentResponse)
	err := c.cc.Invoke(ctx, "/rental.NSAService/ConfirmApplication", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NSAServiceServer is the server API for NSAService service.
// All implementations must embed UnimplementedNSAServiceServer
// for forward compatibility
type NSAServiceServer interface {
	// unary
	SubmitApplication(context.Context, *SubmitApplicationRequest) (*GetApplicationResponse, error)
	SubmitApplicationDocs(context.Context, *SubmitApplicationDocsRequest) (*common.NoContentResponse, error)
	SubmitApplicationPayment(context.Context, *SubmitApplicationPaymentRequest) (*common.NoContentResponse, error)
	CheckExistApplication(context.Context, *CheckExistApplicationRequest) (*common.BooleanResponse, error)
	CheckExistApplications(context.Context, *CheckExistApplicationsRequest) (*CheckExistApplicationsResponse, error)
	GetApplication(context.Context, *common.FindByIdRequest) (*GetApplicationResponse, error)
	UpdateApplication(context.Context, *SubmitApplicationRequest) (*GetApplicationResponse, error)
	ConfirmApplication(context.Context, *ConfirmApplicationRequest) (*common.NoContentResponse, error)
	mustEmbedUnimplementedNSAServiceServer()
}

// UnimplementedNSAServiceServer must be embedded to have forward compatible implementations.
type UnimplementedNSAServiceServer struct {
}

func (UnimplementedNSAServiceServer) SubmitApplication(context.Context, *SubmitApplicationRequest) (*GetApplicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitApplication not implemented")
}
func (UnimplementedNSAServiceServer) SubmitApplicationDocs(context.Context, *SubmitApplicationDocsRequest) (*common.NoContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitApplicationDocs not implemented")
}
func (UnimplementedNSAServiceServer) SubmitApplicationPayment(context.Context, *SubmitApplicationPaymentRequest) (*common.NoContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SubmitApplicationPayment not implemented")
}
func (UnimplementedNSAServiceServer) CheckExistApplication(context.Context, *CheckExistApplicationRequest) (*common.BooleanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckExistApplication not implemented")
}
func (UnimplementedNSAServiceServer) CheckExistApplications(context.Context, *CheckExistApplicationsRequest) (*CheckExistApplicationsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckExistApplications not implemented")
}
func (UnimplementedNSAServiceServer) GetApplication(context.Context, *common.FindByIdRequest) (*GetApplicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetApplication not implemented")
}
func (UnimplementedNSAServiceServer) UpdateApplication(context.Context, *SubmitApplicationRequest) (*GetApplicationResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateApplication not implemented")
}
func (UnimplementedNSAServiceServer) ConfirmApplication(context.Context, *ConfirmApplicationRequest) (*common.NoContentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ConfirmApplication not implemented")
}
func (UnimplementedNSAServiceServer) mustEmbedUnimplementedNSAServiceServer() {}

// UnsafeNSAServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NSAServiceServer will
// result in compilation errors.
type UnsafeNSAServiceServer interface {
	mustEmbedUnimplementedNSAServiceServer()
}

func RegisterNSAServiceServer(s grpc.ServiceRegistrar, srv NSAServiceServer) {
	s.RegisterService(&NSAService_ServiceDesc, srv)
}

func _NSAService_SubmitApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSAServiceServer).SubmitApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rental.NSAService/SubmitApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSAServiceServer).SubmitApplication(ctx, req.(*SubmitApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSAService_SubmitApplicationDocs_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitApplicationDocsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSAServiceServer).SubmitApplicationDocs(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rental.NSAService/SubmitApplicationDocs",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSAServiceServer).SubmitApplicationDocs(ctx, req.(*SubmitApplicationDocsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSAService_SubmitApplicationPayment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitApplicationPaymentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSAServiceServer).SubmitApplicationPayment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rental.NSAService/SubmitApplicationPayment",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSAServiceServer).SubmitApplicationPayment(ctx, req.(*SubmitApplicationPaymentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSAService_CheckExistApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckExistApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSAServiceServer).CheckExistApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rental.NSAService/CheckExistApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSAServiceServer).CheckExistApplication(ctx, req.(*CheckExistApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSAService_CheckExistApplications_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckExistApplicationsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSAServiceServer).CheckExistApplications(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rental.NSAService/CheckExistApplications",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSAServiceServer).CheckExistApplications(ctx, req.(*CheckExistApplicationsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSAService_GetApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.FindByIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSAServiceServer).GetApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rental.NSAService/GetApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSAServiceServer).GetApplication(ctx, req.(*common.FindByIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSAService_UpdateApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SubmitApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSAServiceServer).UpdateApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rental.NSAService/UpdateApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSAServiceServer).UpdateApplication(ctx, req.(*SubmitApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _NSAService_ConfirmApplication_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ConfirmApplicationRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NSAServiceServer).ConfirmApplication(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rental.NSAService/ConfirmApplication",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NSAServiceServer).ConfirmApplication(ctx, req.(*ConfirmApplicationRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// NSAService_ServiceDesc is the grpc.ServiceDesc for NSAService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NSAService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "rental.NSAService",
	HandlerType: (*NSAServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SubmitApplication",
			Handler:    _NSAService_SubmitApplication_Handler,
		},
		{
			MethodName: "SubmitApplicationDocs",
			Handler:    _NSAService_SubmitApplicationDocs_Handler,
		},
		{
			MethodName: "SubmitApplicationPayment",
			Handler:    _NSAService_SubmitApplicationPayment_Handler,
		},
		{
			MethodName: "CheckExistApplication",
			Handler:    _NSAService_CheckExistApplication_Handler,
		},
		{
			MethodName: "CheckExistApplications",
			Handler:    _NSAService_CheckExistApplications_Handler,
		},
		{
			MethodName: "GetApplication",
			Handler:    _NSAService_GetApplication_Handler,
		},
		{
			MethodName: "UpdateApplication",
			Handler:    _NSAService_UpdateApplication_Handler,
		},
		{
			MethodName: "ConfirmApplication",
			Handler:    _NSAService_ConfirmApplication_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "rental/nsa.proto",
}
