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

type FloorClient struct {
	client market.FloorServiceClient
}

func NewFloorClient(cc *grpc.ClientConn) *FloorClient {
	return &FloorClient{client: market.NewFloorServiceClient(cc)}
}

func (service FloorClient) ListFloors(req *market.ListFloorsRequest, md metadata.MD) (res *market.ListFloorsResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.client.ListFloors(ctx, req)
	return
}

func (service FloorClient) GetFloor(req *market.GetFloorRequest, md metadata.MD) (res *market.GetFloorResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.client.GetFloor(ctx, req)
	return
}

func (service FloorClient) CreateFloor(req *market.CreateFloorRequest, md metadata.MD) (res *market.CreateFloorResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.client.CreateFloor(ctx, req)
	return
}

func (service FloorClient) UpdateFloor(req *market.UpdateFloorRequest, md metadata.MD) (res *market.UpdateFloorResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.client.UpdateFloor(ctx, req)
	return
}

func (service FloorClient) DeleteFloor(req *market.DeleteFloorRequest, md metadata.MD) (res *common.NoContentResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.client.DeleteFloor(ctx, req)
	return
}
