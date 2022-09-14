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
