package client

import (
	"context"
	"gateway-service/client/account"
	"gateway-service/client/common"
	"gateway-service/common/constants"
	"google.golang.org/grpc"
	"time"
)

type PermissionClient struct {
	client account.PermissionServiceClient
}

func NewPermissionClient(cc *grpc.ClientConn) *PermissionClient {
	return &PermissionClient{
		client: account.NewPermissionServiceClient(cc),
	}
}

func (service PermissionClient) ListUserPermissions(req *common.FindByIdRequest) (res *account.ListPermissionByUserResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	defer cancel()
	res, err = service.client.ListPermissionsByUser(ctx, req)
	return
}

func (service PermissionClient) ListPermissions(req *account.ListPermissionsRequest) (res *common.ListResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	defer cancel()
	res, err = service.client.ListPermissions(ctx, req)
	return
}

func (service PermissionClient) ListPermissionCategories() (res *common.ListResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	defer cancel()
	res, err = service.client.ListPermissionCategories(ctx, &account.ListPermissionCategoryRequest{})
	return
}
