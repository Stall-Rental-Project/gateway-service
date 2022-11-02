package controller

import (
	"gateway-service/client"
	common2 "gateway-service/client/common"
	"gateway-service/client/market"
	"gateway-service/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type StallController struct {
	stallClient *client.StallClient
	rateClient  *client.RateClient
}

func NewStallController(stallClient *client.StallClient,
	rateClient *client.RateClient) StallController {
	return StallController{
		stallClient: stallClient,
		rateClient:  rateClient,
	}
}

// CreateStall
// @Router /api/v2/stalls [POST]
// @Summary Create Stalls
// @Param _ body market.CreateStallRequest true "request body"
// @Tags Stall
// @Accept json
// @Produce json
// @Success 200 {object} market.Stall
// @Failure 400,401,500 {object} model.ErrorResponse
func (controller *StallController) CreateStall(ctx *gin.Context) {
	req := new(market.CreateStallRequest)

	if err := ctx.ShouldBindJSON(req); err != nil {
		common.ReturnErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res, err := controller.stallClient.CreateStall(req, common.GetMetadataFromContext(ctx))

	if err != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if res.Success {
		common.AsSuccessResponse(res.GetData().Stall, ctx)
	} else {
		common.AsErrorResponse(res.GetError(), ctx)
	}
}

// UpdateStallMetadata
// @Router /api/v2/stalls/metadata [PUT]
// @Summary Update Stall Metadata
// @Param _ body market.UpdateStallMetadataRequest true "request body"
// @Tags Stall
// @Accept json
// @Produce json
// @Success 200 {object} market.Stall
// @Failure 400,401,500 {object} model.ErrorResponse
func (controller *StallController) UpdateStallMetadata(ctx *gin.Context) {
	req := new(market.UpdateStallMetadataRequest)

	if err := ctx.ShouldBindJSON(req); err != nil {
		common.ReturnErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res, err := controller.stallClient.UpdateStallMetadata(req, common.GetMetadataFromContext(ctx))

	if err != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if res.Success {
		common.AsSuccessResponse(res.GetData().Stall, ctx)
	} else {
		common.AsErrorResponse(res.GetError(), ctx)
	}
}

// UpdateStallPosition
// @Router /api/v2/stalls/position [PUT]
// @Summary Update Stall Position (single)
// @Param _ body market.UpdateStallPositionRequest true "request body"
// @Tags Stall
// @Accept json
// @Produce json
// @Success 200 {object} market.Stall
// @Failure 400,401,500 {object} model.ErrorResponse
func (controller *StallController) UpdateStallPosition(ctx *gin.Context) {
	req := new(market.UpdateStallPositionRequest)

	if err := ctx.ShouldBindJSON(req); err != nil {
		common.ReturnErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res, err := controller.stallClient.UpdateStallPosition(req, common.GetMetadataFromContext(ctx))

	if err != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if res.Success {
		common.AsSuccessResponse(res.GetData().Stall, ctx)
	} else {
		common.AsErrorResponse(res.GetError(), ctx)
	}
}

// GetStall
// @Router /api/v2/stalls/:id [GET]
// @Summary Get Stall
// @Param id path string true "stall id"
// @Tags Stall
// @Accept json
// @Produce json
// @Success 200 {object} market.Stall
// @Failure 400,401,500 {object} model.ErrorResponse
func (controller *StallController) GetStall(ctx *gin.Context) {
	stallId := ctx.Param("id")

	draft := false

	if ctx.Query("draft") == "true" {
		draft = true
	}

	req := &market.GetStallRequest{
		StallId: stallId,
		Draft:   &draft,
	}

	res, err := controller.stallClient.GetStall(req, common.GetMetadataFromContext(ctx))

	if err != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if res.Success {
		stall := res.GetData().Stall
		//stall.MonthlyFee = controller.calculateMonthlyFee(stall, ctx)
		common.AsSuccessResponse(stall, ctx)
	} else {
		common.AsErrorResponse(res.GetError(), ctx)
	}
}

//func (controller *StallController) calculateMonthlyFee(stall *market.Stall, ctx *gin.Context) float64 {
//	if stall.GetMarketType() == market.MarketType_MARKET_TYPE_PUBLIC {
//		req := &rental.CalculateRateRequest{
//			MarketClass:    stall.GetMarketClass(),
//			StallClass:     stall.GetClazz(),
//			StallType:      stall.GetType(),
//			StallSectionId: stall.GetSection().GetSectionId(),
//			StallArea:      stall.GetArea(),
//			Includes:       []rental.FeeType{rental.FeeType_NSA_MONTHLY_FEE},
//		}
//
//		resp, err := controller.rateClient.CalculateApplicationRate(req, common.GetMetadataFromContext(ctx))
//
//		if err != nil {
//			log.Println(err.Error())
//			return 0
//		}
//
//		if !resp.GetSuccess() {
//			log.Println(resp.GetError().GetMessage())
//			return 0
//		}
//
//		return resp.GetData().GetMonthlyFee()
//	} else {
//		return 0
//	}
//}

//func (controller *StallController) getNewApplicationStallRate(ctx *gin.Context) float64 {
//	res, err := controller.rateClient.GetNewApplicationRate(common.GetMetadataFromContext(ctx))
//
//	if err != nil {
//		log.Println("Error when get new application stall rate", err)
//		return 0
//	} else if !res.GetSuccess() {
//		log.Println("Error when get new application stall rate", res.GetError().GetMessage())
//		return 0
//	} else {
//		return res.GetData().GetAmount()
//	}
//}

// GetPublishedStall
// @Router /api/v2/stalls/:id/published [GET]
// @Summary Get Published Stall
// @Param id path string true "stall id"
// @Tags Stall
// @Accept json
// @Produce json
// @Success 200 {object} market.Stall
// @Failure 400,401,500 {object} model.ErrorResponse
func (controller *StallController) GetPublishedStall(ctx *gin.Context) {
	stallID := ctx.Param("id")

	req := &common2.FindByIdRequest{
		Id: stallID,
	}

	res, err := controller.stallClient.GetPublishedStall(req, common.GetMetadataFromContext(ctx))

	if err != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if res.Success {
		stall := res.GetData().GetStall()
		//stall.MonthlyFee = controller.calculateMonthlyFee(stall, ctx)
		//if stall.OccupiedBy != "" {
		//	stallholderName, appErr := controller.obtainStallHolderName(stall, ctx)
		//	if appErr != nil {
		//		log.Println("Error when obtaining stall holder name for stall " + stall.StallId + ": " + appErr.Error())
		//	} else {
		//		stall.StallHolderName = stallholderName
		//	}
		//}
		common.AsSuccessResponse(stall, ctx)
	} else {
		common.AsErrorResponse(res.GetError(), ctx)
	}
}

//func (controller *StallController) obtainStallHolderName(stall *market.Stall, ctx *gin.Context) (string, error) {
//	res, err := controller.applicationClient.GetApplication(&common2.FindByIdRequest{
//		Id: stall.OccupiedBy,
//	}, common.GetMetadataFromContext(ctx))
//
//	if err != nil {
//		return "", err
//	}
//
//	if !res.Success {
//		return "", errors.New(res.GetError().GetMessage())
//	}
//
//	return res.GetData().GetApplication().GetOwner().GetFirstName() + " " + res.GetData().GetApplication().GetOwner().GetLastName(), nil
//}
