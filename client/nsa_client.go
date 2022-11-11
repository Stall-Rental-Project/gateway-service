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

type NSAClient struct {
	nsaClient rental.NSAServiceClient
}

func NewNSAClient(cc *grpc.ClientConn) *NSAClient {
	nsaClient := rental.NewNSAServiceClient(cc)
	return &NSAClient{
		nsaClient: nsaClient,
	}
}

func (service *NSAClient) SubmitApplication(req *rental.SubmitApplicationRequest, md metadata.MD) (res *rental.GetApplicationResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.nsaClient.SubmitApplication(ctx, req)
	return
}

func (service *NSAClient) SubmitApplicationDocs(req *rental.SubmitApplicationDocsRequest, md metadata.MD) (res *common.NoContentResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.nsaClient.SubmitApplicationDocs(ctx, req)
	return
}

func (service *NSAClient) UpdateApplication(req *rental.SubmitApplicationRequest, md metadata.MD) (res *rental.GetApplicationResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.nsaClient.UpdateApplication(ctx, req)
	return
}

func (service *NSAClient) GetApplication(applicationId string, md metadata.MD) (res *rental.GetApplicationResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.nsaClient.GetApplication(ctx, &common.FindByIdRequest{Id: applicationId})
	return
}
func (service *NSAClient) CheckExistApplication(req *rental.CheckExistApplicationRequest, md metadata.MD) (res *common.BooleanResponse, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
	ctx = metadata.NewOutgoingContext(ctx, md)
	defer cancel()
	res, err = service.nsaClient.CheckExistApplication(ctx, req)
	return
}

//func (service *NSAClient) SubmitApplicationPayment(req *rental.SubmitApplicationPaymentRequest, md metadata.MD) (res *common.NoContentResponse, err error) {
//	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
//	ctx = metadata.NewOutgoingContext(ctx, md)
//	defer cancel()
//	res, err = service.nsaClient.SubmitApplicationPayment(ctx, req)
//	return
//}
//

//
//func (service *NSAClient) CheckExistApplications(req *rental.CheckExistApplicationsRequest, md metadata.MD) (res *rental.CheckExistApplicationsResponse, err error) {
//	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
//	ctx = metadata.NewOutgoingContext(ctx, md)
//	defer cancel()
//	res, err = service.nsaClient.CheckExistApplications(ctx, req)
//	return
//}
//
//func (service *NSAClient) SubmitOrderPayment(req *rental.SubmitOrderPaymentRequest, md metadata.MD) (res *common.NoContentResponse, err error) {
//	ctx, cancel := context.WithTimeout(context.Background(), constants.GrpcTimeoutInSecs*time.Second)
//	ctx = metadata.NewOutgoingContext(ctx, md)
//	defer cancel()
//	res, err = service.nsaClient.SubmitOrderPayment(ctx, req)
//	return
//}
//
