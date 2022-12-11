package client

import (
	"context"
	"gateway-service/client/market"
	"gateway-service/common/constants"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

type MarketAnalyticClient struct {
	client market.MarketAnalyticsServiceClient
}

func NewMarketAnalyticClient(cc *grpc.ClientConn) *MarketAnalyticClient {
	return &MarketAnalyticClient{
		client: market.NewMarketAnalyticsServiceClient(cc),
	}
}

func (service *MarketAnalyticClient) GetMarketStallAnalytics(req *market.GetMarketStallAnalyticsRequest, md metadata.MD) (res *market.GetMarketStallAnalyticsResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.client.GetMarketStallAnalytics(ctx, req)
	return
}
