package client

import (
	"context"
	"gateway-service/client/common"
	"gateway-service/client/market"
	"gateway-service/common/constants"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

type StallClient struct {
	client market.StallServiceClient
}

func NewStallClient(cc *grpc.ClientConn) *StallClient {
	return &StallClient{client: market.NewStallServiceClient(cc)}
}

func (service *StallClient) CreateStall(req *market.CreateStallRequest, md metadata.MD) (res *market.GetStallResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.client.CreateStall(ctx, req)
	return
}

func (service *StallClient) UpdateStallMetadata(req *market.UpdateStallMetadataRequest, md metadata.MD) (res *market.GetStallResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.client.UpdateStallMetadata(ctx, req)
	return
}

func (service *StallClient) UpdateStallPosition(req *market.UpdateStallPositionRequest, md metadata.MD) (res *market.GetStallResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.client.UpdateStallPosition(ctx, req)
	return
}

func (service *StallClient) GetStall(req *market.GetStallRequest, md metadata.MD) (res *market.GetStallResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.client.GetStall(ctx, req)
	return
}

func (service *StallClient) GetPublishedStall(req *common.FindByIdRequest, md metadata.MD) (res *market.GetStallResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.client.GetPublishedStall(ctx, req)
	return
}
func (service *StallClient) GetStallInfo(req *market.GetStallInfoRequest, md metadata.MD) (res *market.GetStallInfoResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.client.GetStallInfo(ctx, req)
	return
}

func (service *StallClient) ListStallsInfo(req *market.ListStallsInfoRequest, md metadata.MD) (res *market.ListStallsInfoResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.client.ListStallsInfo(ctx, req)
	return
}
