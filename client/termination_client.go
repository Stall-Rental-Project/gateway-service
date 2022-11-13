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

type TerminationClient struct {
	client rental.LeaseTerminationServiceClient
}

func NewTerminationClient(cc *grpc.ClientConn) *TerminationClient {
	return &TerminationClient{
		client: rental.NewLeaseTerminationServiceClient(cc),
	}
}

func (service *TerminationClient) GetTermination(req *common.FindByIdRequest, md metadata.MD) (res *rental.GetLeaseTerminationResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.client.GetLeaseTermination(ctx, req)
	return
}

func (service *TerminationClient) SubmitTermination(req *rental.CreateLeaseTerminationRequest, md metadata.MD) (res *common.NoContentResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.client.CreateLeaseTermination(ctx, req)
	return
}

func (service *TerminationClient) ProceedTermination(req *rental.ProceedLeaseTerminationRequest, md metadata.MD) (res *common.NoContentResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.client.ProceedLeaseTermination(ctx, req)
	return
}
