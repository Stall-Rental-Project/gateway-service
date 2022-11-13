package controller

import (
	"gateway-service/client"
	"gateway-service/client/market"
	"gateway-service/client/rental"
	"gateway-service/common"
	"gateway-service/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type NSAController struct {
	nsaClient    *client.NSAClient
	stallClient  *client.StallClient
	marketClient *client.MarketClient
	userClient   *client.UserClient
	rateClient   *client.RateClient
}

func NewNSAController(
	nsaClient *client.NSAClient,
	stallClient *client.StallClient,
	marketClient *client.MarketClient,
	userClient *client.UserClient,
	rateClient *client.RateClient,
) NSAController {
	return NSAController{
		nsaClient:    nsaClient,
		stallClient:  stallClient,
		marketClient: marketClient,
		userClient:   userClient,
		rateClient:   rateClient,
	}
}

// SubmitRequest
// @Summary Submit Application
// @Description Submit new Application
// @Tags Application
// @Accept json
// @Param data body rental.SubmitApplicationRequest true "Application Information"
// @Success 200 {object} model.SuccessResponse
// @Failure 401,400,500 {object} model.ErrorResponse
// @Router /api/v2/applications [POST]
func (controller *NSAController) SubmitRequest(ctx *gin.Context) {
	req := new(rental.SubmitApplicationRequest)

	if err := ctx.ShouldBindJSON(req); err != nil {
		common.ReturnErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	stallResp, stallErr := controller.stallClient.GetStallInfo(&market.GetStallInfoRequest{Searcher: &market.StallSearcher{
		MarketCode: req.MarketCode,
		FloorCode:  req.FloorCode,
		StallCode:  req.StallCode,
	}}, common.GetMetadataFromContext(ctx))

	if stallErr != nil {
		common.ReturnErrorResponse(ctx, http.StatusBadRequest, stallErr.Error())
		return
	} else if !stallResp.Success {
		ctx.JSON(http.StatusInternalServerError, model.AsErrorResponse(stallResp.GetError()))
		return
	} else if stallResp.GetData().GetStall().GetLeaseStatus() != market.StallLeaseStatus_STALL_AVAILABLE {
		common.ReturnErrorResponse(ctx, http.StatusBadRequest, "Stall has been already occupied")
		return
	} else {
		stall := stallResp.GetData().GetStall()
		req.MarketType = stall.GetMarketType()
		req.MarketClass = stall.GetMarketClass()
		req.StallClass = stall.GetStallClass()
		req.StallType = stall.GetStallType()
		req.StallArea = stall.GetArea()
	}

	res, err := controller.nsaClient.SubmitApplication(req, common.GetMetadataFromContext(ctx))

	if err != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if res.GetSuccess() {
		data := res.GetData().GetApplication()
		ctx.JSON(http.StatusOK, model.SuccessResponse{
			Message: "success",
			Data:    data,
		})
	} else {
		ctx.JSON(http.StatusBadRequest, model.AsErrorResponse(res.GetError()))
	}
}

// SubmitDocs
// @Summary Submit Application Documents
// @Description Submit new application documents or update the application documents
// @Tags Application
// @Accept json
// @Param id path string true "Application ID"
// @Param data body rental.SubmitApplicationDocsRequest true "Application Docs Information"
// @Success 200 {object} common.NoContentResponse
// @Failure 401,400,500 {object} model.ErrorResponse
// @Router /api/v2/applications/{id}/docs [PUT]
func (controller *NSAController) SubmitDocs(ctx *gin.Context) {
	applicationId := ctx.Param("id")
	applicationRequest := new(rental.SubmitApplicationDocsRequest)
	applicationRequest.ApplicationId = applicationId

	if err := ctx.ShouldBindJSON(applicationRequest); err != nil {
		common.ReturnErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	res, err := controller.nsaClient.SubmitApplicationDocs(applicationRequest, common.GetMetadataFromContext(ctx))

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

// UpdateApplication
// @Summary Update Application
// @Description Update Application
// @Tags Application
// @Accept json
// @Param id path string true "Application ID"
// @Param data body rental.Application true "Application Information"
// @Success 200 {object} rental.SubmitApplicationRequest
// @Failure 401,400,500 {object} model.ErrorResponse
// @Router /api/v2/applications/{id} [PUT]
func (controller *NSAController) UpdateApplication(ctx *gin.Context) {
	applicationId := ctx.Param("id")

	req := new(rental.SubmitApplicationRequest)
	req.ApplicationId = applicationId
	if err := ctx.ShouldBindJSON(req); err != nil {
		common.ReturnErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	stallResp, stallErr := controller.stallClient.GetStallInfo(&market.GetStallInfoRequest{Searcher: &market.StallSearcher{
		MarketCode: req.MarketCode,
		FloorCode:  req.FloorCode,
		StallCode:  req.StallCode,
	}}, common.GetMetadataFromContext(ctx))

	if stallErr != nil {
		common.ReturnErrorResponse(ctx, http.StatusBadRequest, stallErr.Error())
		return
	} else if !stallResp.Success {
		ctx.JSON(http.StatusBadRequest, model.AsErrorResponse(stallResp.GetError()))
		return
	} else {
		stall := stallResp.GetData().GetStall()
		req.MarketType = stall.GetMarketType()
		req.MarketClass = stall.GetMarketClass()
		req.StallClass = stall.GetStallClass()
		req.StallArea = stall.GetArea()
	}

	res, err := controller.nsaClient.UpdateApplication(req, common.GetMetadataFromContext(ctx))

	if err != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if res.GetSuccess() {
		ctx.JSON(http.StatusOK, res.GetData().GetApplication())
	} else {
		ctx.JSON(http.StatusBadRequest, model.AsErrorResponse(res.GetError()))
	}
}

// GetApplication
// @Summary Get Application
// @Description Get Application
// @Tags Application
// @Accept json
// @Param id path string true "Application ID"
// @Success 200 {object} rental.Application
// @Failure 401,400,500 {object} model.ErrorResponse
// @Router /api/v2/applications/{id} [GET]
func (controller *NSAController) GetApplication(ctx *gin.Context) {
	applicationId := ctx.Param("id")

	res, err := controller.nsaClient.GetApplication(applicationId, common.GetMetadataFromContext(ctx))

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

// SubmitPayment
// @Summary Submit Application Payment
// @Param id path string true "Application ID"
// @Description Submit new application payment or update the application payment
// @Tags Application
// @Accept json
// @Param data body rental.SubmitApplicationPaymentRequest true "Application Payment Information"
// @Success 200 {object} model.SuccessResponse
// @Failure 401,400,500 {object} model.ErrorResponse
// @Router /api/v2/applications/{id}/payment [PUT]
func (controller *NSAController) SubmitPayment(ctx *gin.Context) {
	applicationId := ctx.Param("id")

	applicationRequest := new(rental.SubmitApplicationPaymentRequest)
	applicationRequest.ApplicationId = applicationId
	if err := ctx.ShouldBindJSON(applicationRequest); err != nil {
		common.ReturnErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res, err := controller.nsaClient.SubmitApplicationPayment(applicationRequest, common.GetMetadataFromContext(ctx))

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

// ConfirmApplication
// @Summary Confirm Application
// @Param id path string true "Application ID"
// @Description Confirm the Application to finish the application process
// @Tags Application
// @Accept json
// @Param data body rental.ConfirmApplicationRequest true "Confirm Application Information"
// @Success 200 {object} model.SuccessResponse
// @Failure 401,400,500 {object} model.ErrorResponse
// @Router /api/v2/applications/{id}/confirm [PUT]
func (controller *NSAController) ConfirmApplication(ctx *gin.Context) {
	applicationId := ctx.Param("id")

	applicationRequest := new(rental.ConfirmApplicationRequest)
	applicationRequest.ApplicationId = applicationId
	if err := ctx.ShouldBindJSON(applicationRequest); err != nil {
		common.ReturnErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res, err := controller.nsaClient.ConfirmApplication(applicationRequest, common.GetMetadataFromContext(ctx))

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
