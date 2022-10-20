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

type MarketClient struct {
	marketClient market.MarketServiceClient
}

func NewMarketClient(cc *grpc.ClientConn) *MarketClient {
	marketClient := market.NewMarketServiceClient(cc)
	return &MarketClient{marketClient: marketClient}
}

func (client *MarketClient) ListMarkets(req *market.ListMarketsRequest, md metadata.MD) (res *common.PageResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()

	res, err = client.marketClient.ListMarkets(ctx, req)
	return
}

func (client *MarketClient) CreateMarket(req *market.UpsertMarketRequest, md metadata.MD) (res *common.OnlyIdResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()

	res, err = client.marketClient.CreateMarket(ctx, req)
	return
}

func (client *MarketClient) UpdateMarket(req *market.UpsertMarketRequest, md metadata.MD) (res *market.UpdateMarketResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()

	res, err = client.marketClient.UpdateMarket(ctx, req)
	return
}

func (client *MarketClient) GetMarket(req *market.GetMarketRequest, md metadata.MD) (res *market.GetMarketResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()

	res, err = client.marketClient.GetMarket(ctx, req)
	return
}

func (client *MarketClient) DeleteMarket(req *common.FindByIdRequest, md metadata.MD) (res *common.NoContentResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()

	res, err = client.marketClient.DeleteMarket(ctx, req)
	return
}
