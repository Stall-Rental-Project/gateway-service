package client

import (
	"context"
	"gateway-service/client/account"
	"gateway-service/common/constants"
	"google.golang.org/grpc"
	"log"
	"time"
)

/*
 * gRPC Account Client
 */

type AuthClient struct {
	authClient account.AuthServiceClient
}

func NewAuthClient(cc *grpc.ClientConn) *AuthClient {
	accountClient := account.NewAuthServiceClient(cc)
	return &AuthClient{authClient: accountClient}
}

func (service *AuthClient) Login(req *account.LoginRequest) (res *account.LoginResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	defer cancel()
	res, err = service.authClient.Login(ctx, req)
	log.Println("Failed when retrieving user permission", err.Error())
	return
}
