package controller

import (
	"gateway-service/client"
	grpc "gateway-service/client/common"
	"gateway-service/client/market"
	"gateway-service/client/rental"
	"gateway-service/common"
	"gateway-service/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type ApplicationController struct {
	applicationClient *client.ApplicationClient
	stallClient       *client.StallClient
}

func NewApplicationController(
	applicationClient *client.ApplicationClient,
	stallClient *client.StallClient,
) ApplicationController {
	return ApplicationController{
		applicationClient: applicationClient,
		stallClient:       stallClient,
	}
}

// ListApplication
// @Summary List Application
// @Description Get applications
// @Tags Application
// @Accept json
// @Param size query integer false "Page size"
// @Param page query integer false "Page number"
// @Param sort query string false "Sort field. Accept value[ first_name, last_name, code]"
// @Param direction query string false "Sort direction"
// @Param code query string false "Filter By Application Code"
// @Param first_name query string false "Filter by owner first name"
// @Param last_name query string false "Filter by owner last name"
// @Success 200 {object} model.PageResponse
// @Failure 401,400,500 {object} model.ErrorResponse
// @Router /api/v2/applications [GET]
func (controller *ApplicationController) ListApplication(ctx *gin.Context) {
	req := buildListApplicationsRequest(ctx)

	res, err := controller.applicationClient.ListApplication(req, common.GetMetadataFromContext(ctx))

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

func buildListApplicationsRequest(ctx *gin.Context) *rental.ListApplicationRequest {
	pageRequest := common.AsPageRequest(ctx)
	code := ctx.Query("code")
	firstName := ctx.Query("first_name")
	lastName := ctx.Query("last_name")

	request := &rental.ListApplicationRequest{
		PageRequest: &pageRequest,
		Code:        code,
		FirstName:   firstName,
		LastName:    lastName,
	}
	return request
}

// CancelApplication
// @Summary Cancel application
// @Tags Application
// @Accept json
// @Param id path string true "Application ID"
// @Success 200 {object} common.NoContentResponse
// @Failure 401,400,500 {object} model.ErrorResponse
// @Router /api/v2/applications/{id} [DELETE]
func (controller *ApplicationController) CancelApplication(ctx *gin.Context) {

	applicationId := ctx.Param("id")

	res, err := controller.applicationClient.CancelApplication(&grpc.FindByIdRequest{
		Id: applicationId,
	}, common.GetMetadataFromContext(ctx))

	if err != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if res.GetSuccess() {
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, model.AsErrorResponse(res.GetError()))
	}
}
