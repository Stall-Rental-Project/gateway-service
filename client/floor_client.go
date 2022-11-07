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

func (service FloorClient) CreateFloor(req *market.UpsertFloorRequest, md metadata.MD) (res *market.UpsertFloorResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.client.CreateFloor(ctx, req)
	return
}

func (service FloorClient) UpdateFloor(req *market.UpsertFloorRequest, md metadata.MD) (res *market.UpsertFloorResponse, err error) {
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

func (service FloorClient) ListPublishedFloors(req *common.FindByIdRequest, md metadata.MD) (res *market.ListFloorsResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.client.ListPublishedFloors(ctx, req)
	return
}
func (service FloorClient) GetPublishedFloor(req *market.GetPublishedFloorRequest, md metadata.MD) (res *market.GetFloorResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.client.GetPublishedFloor(ctx, req)
	return
}
