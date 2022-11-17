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

/*
 * gRPC Role Client
 */

type RateClient struct {
	rateClient rental.RateServiceClient
}

func NewRateClient(cc *grpc.ClientConn) *RateClient {
	rateClient := rental.NewRateServiceClient(cc)
	return &RateClient{rateClient: rateClient}
}

func (service *RateClient) CreateRate(req *rental.UpsertRateRequest, md metadata.MD) (res *common.OnlyCodeResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.rateClient.CreateRate(ctx, req)
	return
}

func (service *RateClient) UpdateRate(req *rental.UpsertRateRequest, md metadata.MD) (res *common.OnlyCodeResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.rateClient.UpdateRate(ctx, req)
	return
}

func (service *RateClient) ListRates(req *rental.ListRatesRequest, md metadata.MD) (res *common.PageResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.rateClient.ListRates(ctx, req)
	return
}

func (service *RateClient) GetRate(req *rental.GetRateRequest, md metadata.MD) (res *rental.GetRateResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.rateClient.GetRate(ctx, req)
	return
}

func (service *RateClient) DeleteRate(req *common.FindByIdRequest, md metadata.MD) (res *common.NoContentResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.rateClient.DeleteRate(ctx, req)
	return
}

func (service RateClient) CalculateApplicationRate(req *rental.CalculateRateRequest, md metadata.MD) (res *rental.CalculateRateResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.rateClient.CalculateRates(ctx, req)
	return
}
