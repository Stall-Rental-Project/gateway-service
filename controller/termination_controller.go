package controller

import (
	"gateway-service/client"
	common2 "gateway-service/client/common"
	"gateway-service/client/rental"
	"gateway-service/common"
	"gateway-service/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type TerminationController struct {
	terminationClient *client.TerminationClient
}

func NewTerminationController(terminationClient *client.TerminationClient) *TerminationController {
	return &TerminationController{
		terminationClient: terminationClient,
	}
}

// GetLeaseTermination
// @Summary Get Application Lease Termination
// @Tags Lease Termination
// @Accept json
// @Param id path string true "Application ID"
// @Success 200 {object} rental.GetLeaseTerminationResponse_Data
// @Failure 401,400,500 {object} model.ErrorResponse
// @Router /api/v2/applications/{id}/termination [GET]
func (controller *TerminationController) GetLeaseTermination(ctx *gin.Context) {
	applicationId := ctx.Param("id")

	res, err := controller.terminationClient.GetTermination(&common2.FindByIdRequest{Id: applicationId}, common.GetMetadataFromContext(ctx))

	if err != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if res.GetSuccess() {
		ctx.JSON(http.StatusOK, res.GetData())
	} else {
		ctx.JSON(http.StatusBadRequest, model.AsErrorResponse(res.GetError()))
	}
}

// SubmitLeaseTermination
// @Summary Submit Application Lease Termination
// @Tags Lease Termination
// @Accept json
// @Param id path string true "Application ID"
// @Param data body rental.CreateLeaseTerminationRequest true "Termination information"
// @Success 200 {object} common.NoContentResponse
// @Failure 401,400,500 {object} model.ErrorResponse
// @Router /api/v2/applications/{id}/termination [POST]
func (controller *TerminationController) SubmitLeaseTermination(ctx *gin.Context) {
	applicationId := ctx.Param("id")

	req := new(rental.CreateLeaseTerminationRequest)

	if err := ctx.ShouldBindJSON(req); err != nil {
		common.ReturnErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	req.ApplicationId = applicationId

	res, err := controller.terminationClient.SubmitTermination(req, common.GetMetadataFromContext(ctx))

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

// CancelLeaseTermination
// @Summary Cancel Lease Termination
// @Tags Lease Termination
// @Accept json
// @Param id path string true "Application ID"
// @Param tid path string true "Termination ID"
// @Param data body rental.ProceedLeaseTerminationRequest true "Termination information"
// @Success 200 {object} common.NoContentResponse
// @Failure 401,400,500 {object} model.ErrorResponse
// @Router /api/v2/applications/{id}/termination/{tid} [PUT]
func (controller *TerminationController) CancelLeaseTermination(ctx *gin.Context) {
	terminationId := ctx.Param("tid")

	req := new(rental.ProceedLeaseTerminationRequest)

	if err := ctx.ShouldBindJSON(req); err != nil {
		common.ReturnErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	req.TerminationId = terminationId

	res, err := controller.terminationClient.ProceedTermination(req, common.GetMetadataFromContext(ctx))

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
