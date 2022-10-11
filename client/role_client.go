package client

import (
	"context"
	"gateway-service/client/account"
	"gateway-service/client/common"
	"gateway-service/common/constants"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

/*
 * gRPC Role Client
 */

type RoleClient struct {
	roleClient account.RoleServiceClient
}

func NewRoleClient(cc *grpc.ClientConn) *RoleClient {
	roleClient := account.NewRoleServiceClient(cc)
	return &RoleClient{roleClient: roleClient}
}

func (service *RoleClient) CreateRole(req *account.UpsertRoleRequest, md metadata.MD) (res *common.NoContentResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.roleClient.CreateRole(ctx, req)
	return
}

func (service *RoleClient) UpdateRole(req *account.UpsertRoleRequest, md metadata.MD) (res *common.NoContentResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.roleClient.UpdateRole(ctx, req)
	return
}

func (service *RoleClient) DeleteRole(req *common.FindByIdRequest, md metadata.MD) (res *common.NoContentResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.roleClient.DeleteRole(ctx, req)
	return
}

func (service *RoleClient) ListRoles(req *account.ListRoleRequest, md metadata.MD) (res *common.PageResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.roleClient.ListRoles(ctx, req)
	return
}

func (service *RoleClient) GetRole(req *common.FindByIdRequest, md metadata.MD) (res *account.GetRoleResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.roleClient.GetRole(ctx, req)
	return
}

func (service *RoleClient) FindRolesByIds(req *common.FindByIdsRequest) (res *common.PageResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	defer cancel()
	res, err = service.roleClient.FindAllByIds(ctx, req)
	return
}
