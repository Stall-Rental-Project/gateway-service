package client

import (
	"context"
	"gateway-service/client/common"
	"gateway-service/client/rental"
	"gateway-service/common/constants"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

type ApplicationClient struct {
	applicationClient rental.ApplicationServiceClient
}

func NewApplicationClient(cc *grpc.ClientConn) *ApplicationClient {
	applicationClient := rental.NewApplicationServiceClient(cc)
	return &ApplicationClient{
		applicationClient: applicationClient,
	}
}

func (service *ApplicationClient) ListApplication(req *rental.ListApplicationRequest, md metadata.MD) (res *common.PageResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.applicationClient.ListApplications(ctx, req)
	return
}

func (service *ApplicationClient) CancelApplication(req *rental.CancelApplicationRequest, md metadata.MD) (res *common.NoContentResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.applicationClient.CancelApplication(ctx, req)
	return
}
