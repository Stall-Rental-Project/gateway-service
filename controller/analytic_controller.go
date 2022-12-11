package controller

import (
	"gateway-service/client"
	"gateway-service/client/market"
	"gateway-service/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AnalyticController struct {
	marketAnalyticClient *client.MarketAnalyticClient
}

func NewAnalyticController(marketAnalyticClient *client.MarketAnalyticClient,

) AnalyticController {
	return AnalyticController{
		marketAnalyticClient: marketAnalyticClient,
	}
}

// GetMarketStallAnalytic
// @Router /api/v2/analytic/market/:code [GET]
// @Summary Get Market Stall Analytic
// @Param code path string true "Market Name"
// @Tags Analytic
// @Accept json
// @Produce json
// @Success 200 {object} market.GetMarketStallAnalyticsResponse_Data
// @Failure 400,401,500 {object} model.ErrorResponse
func (controller *AnalyticController) GetMarketStallAnalytic(ctx *gin.Context) {
	marketName := ctx.Param("code")

	req := &market.GetMarketStallAnalyticsRequest{
		MarketName: marketName,
	}

	res, err := controller.marketAnalyticClient.GetMarketStallAnalytics(req, common.GetMetadataFromContext(ctx))

	if err != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if res.Success {
		ctx.JSON(http.StatusOK, res.GetData().GetMarketVendorDetail())
	} else {
		common.AsErrorResponse(res.GetError(), ctx)
	}
}
