package client

import (
	"context"
	"gateway-service/client/account"
	"gateway-service/client/common"
	"gateway-service/common/constants"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/types/known/emptypb"
	"time"
)

type UserClient struct {
	client account.UserServiceClient
}

func NewUserClient(cc *grpc.ClientConn) *UserClient {
	return &UserClient{
		client: account.NewUserServiceClient(cc),
	}
}

func (service UserClient) CreateUser(req *account.UpsertUserRequest, md metadata.MD) (res *common.NoContentResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.client.CreateUser(ctx, req)
	return
}

func (service UserClient) ListUsers(req *account.ListUsersRequest, md metadata.MD) (res *common.PageResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.client.ListUsers(ctx, req)
	return
}

func (service UserClient) ListUsersByEmail(req *common.FindByCodeRequest, md metadata.MD) (res *common.ListResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.client.ListUsersByEmail(ctx, req)
	return
}

func (service UserClient) GetUser(req *common.FindByCodeRequest, md metadata.MD) (res *account.GetUserResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.client.GetUser(ctx, req)
	return
}

func (service UserClient) GetCurrentUser(md metadata.MD) (res *account.GetCurrentUserResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.client.GetCurrentUser(ctx, &emptypb.Empty{})
	return
}

func (service UserClient) UpdateUser(req *account.UpsertUserRequest, md metadata.MD) (res *common.NoContentResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.client.UpdateUser(ctx, req)
	return
}

func (service UserClient) DeleteUser(req *common.FindByCodeRequest, md metadata.MD) (res *common.NoContentResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.client.DeleteUser(ctx, req)
	return
}
