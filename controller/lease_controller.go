package controller

import (
	"gateway-service/client"
	"gateway-service/client/market"
	"gateway-service/client/rental"
	"gateway-service/common"
	"gateway-service/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type LeaseController struct {
	leaseClient  *client.LeaseClient
	stallClient  *client.StallClient
	marketClient *client.MarketClient
	userClient   *client.UserClient
	rateClient   *client.RateClient
}

func NewLeaseController(
	leaseClient *client.LeaseClient,
	stallClient *client.StallClient,
	marketClient *client.MarketClient,
	userClient *client.UserClient,
	rateClient *client.RateClient,
) LeaseController {
	return LeaseController{
		leaseClient:  leaseClient,
		stallClient:  stallClient,
		marketClient: marketClient,
		userClient:   userClient,
		rateClient:   rateClient,
	}
}

// ListLeases
// @Summary List In-Lease Application
// @Tags Lease
// @Accept json
// @Param size query integer false "Page size. Default to 20"
// @Param page query integer false "Page number. Default to 1"
// @Param sort query string false "Sort field. Accepts [firstName, lastName]. Default to firstName if not given"
// @Param direction query string false "Sort direction. Accepts [asc, desc]. Default to asc if not given"
// @Param lease_code query string false "Filter By Lease Code/ID"
// @Param first_name query string false "Filter By Owner Firstname"
// @Param last_name query string false "Filter By Owner Lastname"
// @Param market_codes query string false "Filter By market code. Accepts comma-separated market codes"
// @Param lease_statuses query string false "Filter By lease status. Accepts comma-separated numeric string of values [1,2,3,4,5]. Default to 1,2,3"
// @Param payment_statuses query string false "Filter by payment status. Accepts comma-separated numeric string of values [0,1,2,3]. Omit if not given"
// @Success 200 {object} model.PageResponse
// @Failure 401,400,500 {object} model.ErrorResponse
// @Router /api/v2/applications/in-lease [GET]
func (controller *LeaseController) ListLeases(ctx *gin.Context) {
	req := buildListInLeaseApplicationsRequest(ctx)

	res, err := controller.leaseClient.ListLeases(req, common.GetMetadataFromContext(ctx))

	if err != nil {
		log.Println("Failed when list applications", err.Error())
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if res.GetSuccess() {
		var applications []*rental.Application
		var searchers []*market.StallSearcher

		for _, value := range res.GetData().GetItems() {
			var application rental.Application
			err = value.UnmarshalTo(&application)

			if err != nil {
				log.Println("An application has failed to unmarshall", err.Error())
				continue
			}

			searcher := &market.StallSearcher{
				MarketCode: application.MarketCode,
				FloorCode:  application.FloorCode,
				StallCode:  application.StallCode,
			}
			searchers = append(searchers, searcher)

			applications = append(applications, &application)
		}

		/* Stall */
		stalls, err := controller.stallClient.ListStallsInfo(&market.ListStallsInfoRequest{Searchers: searchers}, common.GetMetadataFromContext(ctx))

		if err != nil {
			log.Println("Failed when list stalls info", err.Error())
			common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		mapStalls := make(map[string]*market.StallInfo)
		if stalls != nil && stalls.GetSuccess() && stalls.GetData() != nil && len(stalls.GetData().GetStalls()) > 0 {
			for _, value := range stalls.GetData().GetStalls() {
				mapStalls[value.MarketCode+"_"+value.FloorCode+"_"+value.Code] = value
			}
		}

		for _, application := range applications {
			key := application.MarketCode + "_" + application.FloorCode + "_" + application.StallCode
			if mapStalls[key] != nil {
				application.MarketName = mapStalls[key].MarketName
				application.FloorName = mapStalls[key].FloorName
				application.StallArea = mapStalls[key].Area
				application.StallName = mapStalls[key].StallName
			}
		}

		ctx.JSON(http.StatusOK, common.AsPageResponse(res, applications))
	} else {
		ctx.JSON(http.StatusBadRequest, model.AsErrorResponse(res.GetError()))
	}
}

func buildListInLeaseApplicationsRequest(ctx *gin.Context) *rental.ListLeasesRequest {
	pageRequest := common.AsPageRequest(ctx)
	leaseCode := ctx.Query("lease_code")
	firstName := ctx.Query("first_name")
	lastName := ctx.Query("last_name")

	request := &rental.ListLeasesRequest{
		PageRequest: &pageRequest,
		LeaseCode:   leaseCode,
		FirstName:   firstName,
		LastName:    lastName,
	}

	marketCodes := ctx.Query("market_codes")
	if marketCodes != "" && len(marketCodes) > 0 {
		var codes []string
		for _, s := range strings.Split(marketCodes, ",") {
			codes = append(codes, s)
		}
		request.MarketCodes = codes
	}

	leaseStatuses := ctx.Query("lease_statuses")
	if leaseStatuses != "" && len(leaseStatuses) > 0 {
		var codes []rental.LeaseStatus
		for _, s := range strings.Split(leaseStatuses, ",") {
			v, _ := strconv.ParseInt(s, 10, 32)
			codes = append(codes, rental.LeaseStatus(v))
		}
		request.LeaseStatus = codes
	}

	return request
}

// GetLease
// @Summary Get in-lease Application
// @Tags Lease
// @Accept json
// @Success 200 {object} rental.Application
// @Failure 401,400,500 {object} model.ErrorResponse
// @Router /api/v2/applications/in-lease/{id} [GET]
func (controller *LeaseController) GetLease(ctx *gin.Context) {
	applicationId := ctx.Param("id")

	res, err := controller.leaseClient.GetLease(applicationId, common.GetMetadataFromContext(ctx))

	if err != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if res.GetSuccess() {
		application := res.GetData().GetApplication()
		searcher := &market.StallSearcher{
			MarketCode: application.MarketCode,
			FloorCode:  application.FloorCode,
			StallCode:  application.StallCode,
		}
		stallInfo, err := controller.stallClient.GetStallInfo(&market.GetStallInfoRequest{Searcher: searcher}, common.GetMetadataFromContext(ctx))

		if err != nil {
			common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
			return
		}

		if stallInfo != nil && stallInfo.GetSuccess() {
			stall := stallInfo.GetData().GetStall()
			mapStalls := make(map[string]*market.StallInfo)
			mapStalls[stall.MarketCode+"_"+stall.FloorCode+"_"+stall.Code] = stall

			key := application.MarketCode + "_" + application.FloorCode + "_" + application.StallCode
			if mapStalls[key] != nil {
				application.MarketName = stall.MarketName
				application.MarketAddress = stall.MarketAddress
				application.FloorName = stall.FloorName
				application.StallArea = stall.Area
				application.StallName = stall.StallName
			}
		}

		ctx.JSON(http.StatusOK, application)
	} else {
		ctx.JSON(http.StatusBadRequest, model.AsErrorResponse(res.GetError()))
	}
}
