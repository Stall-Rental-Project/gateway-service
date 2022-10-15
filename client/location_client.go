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

type LocationClient struct {
	locationClient market.LocationServiceClient
}

func NewLocationClient(cc *grpc.ClientConn) *LocationClient {
	locationClient := market.NewLocationServiceClient(cc)
	return &LocationClient{locationClient: locationClient}
}

func (client *LocationClient) ListProvinces(req *market.ListProvinceRequest, md metadata.MD) (res *market.ListProvinceResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()

	res, err = client.locationClient.ListProvinces(ctx, req)
	return
}

func (client *LocationClient) ListCities(req *market.ListCityRequest, md metadata.MD) (res *market.ListCityResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()

	res, err = client.locationClient.ListCities(ctx, req)
	return
}

func (client *LocationClient) ListWards(req *market.ListWardRequest, md metadata.MD) (res *market.ListWardResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()

	res, err = client.locationClient.ListWards(ctx, req)
	return
}

func (client *LocationClient) ListLocations(req *common.FindByIdsRequest, md metadata.MD) (res *common.ListResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()

	res, err = client.locationClient.ListLocations(ctx, req)
	return
}

func (client *LocationClient) GetLocation(req *market.GetLocationRequest, md metadata.MD) (res *market.GetLocationResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()

	res, err = client.locationClient.GetLocation(ctx, req)
	return
}
