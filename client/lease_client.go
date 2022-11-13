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
 * gRPC Application Client
 */

type LeaseClient struct {
	LeaseClient rental.LeaseServiceClient
}

func NewLeaseClient(cc *grpc.ClientConn) *LeaseClient {
	leaseClient := rental.NewLeaseServiceClient(cc)
	return &LeaseClient{
		LeaseClient: leaseClient,
	}
}

func (service *LeaseClient) ListLeases(req *rental.ListLeasesRequest, md metadata.MD) (res *common.PageResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.LeaseClient.ListLeases(ctx, req)
	return
}

func (service *LeaseClient) GetLease(applicationId string, md metadata.MD) (res *rental.GetApplicationResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.LeaseClient.GetLease(ctx, &common.FindByIdRequest{
		Id: applicationId,
	})
	return
}
