package controller

import (
	"gateway-service/client"
	common2 "gateway-service/client/common"
	"gateway-service/client/market"
	"gateway-service/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FloorController struct {
	floorClient *client.FloorClient
}

func NewFloorController(floorClient *client.FloorClient) FloorController {
	return FloorController{
		floorClient,
	}
}

// CreateFloor
// @Router /api/v2/floors [POST]
// @Summary Create floor
// @Param _ body market.UpsertFloorRequest true "request body"
// @Tags Floor
// @Accept json
// @Produce json
// @Success 200 {object} market.UpsertFloorResponse_Data
// @Failure 400,401,500 {object} model.ErrorResponse
func (controller *FloorController) CreateFloor(ctx *gin.Context) {
	req := new(market.UpsertFloorRequest)

	if err := ctx.ShouldBindJSON(req); err != nil {
		common.ReturnErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res, err := controller.floorClient.CreateFloor(req, common.GetMetadataFromContext(ctx))

	if err != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if res.Success {
		ctx.JSON(http.StatusOK, res.GetData())
	} else {
		common.AsErrorResponse(res.GetError(), ctx)
	}
}

// UpdateFloor
// @Router /api/v2/floors/{id} [PUT]
// @Summary Update floor
// @Param id path string true "ID"
// @Param _ body market.UpsertFloorRequest true "request body"
// @Tags Floor
// @Accept json
// @Produce json
// @Success 200 {object} market.UpsertFloorResponse_Data
// @Failure 400,401,500 {object} model.ErrorResponse
func (controller *FloorController) UpdateFloor(ctx *gin.Context) {
	req := new(market.UpsertFloorRequest)
	floorId := ctx.Param("id")

	if err := ctx.ShouldBindJSON(req); err != nil {
		common.ReturnErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	req.FloorplanId = floorId
	res, err := controller.floorClient.UpdateFloor(req, common.GetMetadataFromContext(ctx))

	if err != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if res.Success {
		ctx.JSON(http.StatusOK, res.GetData())
	} else {
		common.AsErrorResponse(res.GetError(), ctx)
	}
}

// GetFloor
// @Router /api/v2/floors/:id [GET]
// @Summary Get floor
// @Param id path string true "floor id"
// @Param draft query bool false "return draft version ? default to false"
// @Tags Floor
// @Accept json
// @Produce json
// @Success 200 {object} market.GetFloorResponse_Data
// @Failure 400,401,500 {object} model.ErrorResponse
func (controller *FloorController) GetFloor(ctx *gin.Context) {
	floorID := ctx.Param("id")

	draft := false

	if ctx.Query("draft") == "true" {
		draft = true
	}

	req := &market.GetFloorRequest{
		FloorplanId: floorID,
		Draft:       &draft,
	}

	res, err := controller.floorClient.GetFloor(req, common.GetMetadataFromContext(ctx))

	if err != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if res.Success {
		ctx.JSON(http.StatusOK, res.GetData().GetFloor())
	} else {
		common.AsErrorResponse(res.GetError(), ctx)
	}
}

type ListFloorsResponse struct {
	Floors []*market.Floor `json:"floors"`
}

// ListFloors
// @Router /api/v2/markets/:id/floors [GET]
// @Summary List floors
// @Param id path string true "market id"
// @Param draft query bool false "return draft version ? default to false"
// @Tags Floor
// @Accept json
// @Produce json
// @Success 200 {object} ListFloorsResponse
// @Failure 400,401,500 {object} model.ErrorResponse
func (controller *FloorController) ListFloors(ctx *gin.Context) {
	marketID := ctx.Param("id")
	draft := false

	if ctx.Query("draft") == "true" {
		draft = true
	}

	req := &market.ListFloorsRequest{
		MarketId: marketID,
		Draft:    &draft,
	}

	res, err := controller.floorClient.ListFloors(req, common.GetMetadataFromContext(ctx))

	if err != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if res.Success {
		ctx.JSON(http.StatusOK, ListFloorsResponse{
			Floors: res.GetData().GetFloors(),
		})
	} else {
		common.AsErrorResponse(res.GetError(), ctx)
	}
}

// DeleteFloor
// @Router /api/v2/floors [DELETE]
// @Summary Delete floor (single)
// @Param _ body market.DeleteFloorRequest true "request body"
// @Tags Floor
// @Accept json
// @Produce json
// @Success 200 {object} common.NoContentResponse
// @Failure 400,401,500 {object} model.ErrorResponse
func (controller *FloorController) DeleteFloor(ctx *gin.Context) {
	req := new(market.DeleteFloorRequest)

	if err := ctx.ShouldBindJSON(req); err != nil {
		common.ReturnErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	//floorResp, floorErr := controller.floorClient.GetFloorCodeAndMarketCode(&common2.FindByIdRequest{
	//	Id: req.FloorplanId,
	//}, common.GetMetadataFromContext(ctx))
	//
	//if floorErr != nil {
	//	common.ReturnErrorResponse(ctx, http.StatusBadRequest, floorErr.Error())
	//	return
	//} else if !floorResp.Success {
	//	common.AsErrorResponse(floorResp.GetError(), ctx)
	//	return
	//}
	//
	//applicationResp, applicationErr := controller.applicationClient.CheckExistApplication(&rental.CheckExistApplicationRequest{
	//	MarketCode: floorResp.GetData().MarketCode,
	//	FloorCode:  floorResp.GetData().FloorCode,
	//}, common.GetMetadataFromContext(ctx))
	//
	//if applicationErr != nil {
	//	common.ReturnErrorResponse(ctx, http.StatusBadRequest, applicationErr.Error())
	//	return
	//} else if applicationResp.GetErrorResponse() != nil {
	//	common.AsLegacyErrorResponse(applicationResp.GetErrorResponse(), ctx)
	//	return
	//} else if applicationResp.GetSuccessResponse().Result {
	//	common.AsLegacyErrorResponse(&common2.ErrorResponse{
	//		ErrorCode:        common2.ErrorCode_CANNOT_EXECUTE,
	//		ErrorDescription: "Some application already exists",
	//	}, ctx)
	//	return
	//}

	res, err := controller.floorClient.DeleteFloor(req, common.GetMetadataFromContext(ctx))

	if err != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if res.Success {
		ctx.JSON(http.StatusOK, res)
	} else {
		common.AsErrorResponse(res.GetError(), ctx)
	}
}

// GetPublishedFloor
// @Router /api/v2/floors/:id/published [GET]
// @Summary Get published floor
// @Param id path string true "floor id"
// @Tags Floor
// @Accept json
// @Produce json
// @Success 200 {object} market.Floor
// @Failure 400,401,500 {object} model.ErrorResponse
func (controller *FloorController) GetPublishedFloor(ctx *gin.Context) {
	floorID := ctx.Param("id")

	req := &market.GetPublishedFloorRequest{
		Id: floorID,
	}

	res, err := controller.floorClient.GetPublishedFloor(req, common.GetMetadataFromContext(ctx))

	if err != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if res.Success {
		ctx.JSON(http.StatusOK, res.GetData().GetFloor())
	} else {
		common.AsErrorResponse(res.GetError(), ctx)
	}
}

// ListPublishedFloors
// @Router /api/v2/markets/:id/floors/published [GET]
// @Summary List published floors
// @Param id path string true "market id"
// @Tags Floor
// @Accept json
// @Produce json
// @Success 200 {object} ListFloorsResponse
// @Failure 400,401,500 {object} model.ErrorResponse
func (controller *FloorController) ListPublishedFloors(ctx *gin.Context) {
	marketID := ctx.Param("id")

	req := &common2.FindByIdRequest{
		Id: marketID,
	}

	res, err := controller.floorClient.ListPublishedFloors(req, common.GetMetadataFromContext(ctx))

	if err != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if res.Success {
		ctx.JSON(http.StatusOK, ListFloorsResponse{
			Floors: res.GetData().GetFloors(),
		})
	} else {
		common.AsErrorResponse(res.GetError(), ctx)
	}
}
