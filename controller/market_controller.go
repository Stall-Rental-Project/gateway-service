package controller

import (
	"gateway-service/client"
	grpc "gateway-service/client/common"
	"gateway-service/client/market"
	"gateway-service/common"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type MarketController struct {
	marketClient *client.MarketClient
}

func NewMarketController(marketClient *client.MarketClient) MarketController {
	return MarketController{
		marketClient,
	}
}

type ListMarketsResponse struct {
	Page          int32           `json:"page"`
	Size          int32           `json:"size"`
	TotalElements int64           `json:"total_elements"`
	TotalPages    int32           `json:"total_pages"`
	Items         []market.Market `json:"items"`
}

// ListMarkets
// @Router /api/v2/markets [GET]
// @Summary List market
// @Description Returns the primary (current) version of markets
// @Param search_term query string false "search by name or location (district, ward, detail address)"
// @Param page query int false "query page, start from 1"
// @Param size query int false "query page size, start from 1"
// @Param sort query string false "sort field, default to name, accepts - name, location, types, statuses"
// @Param direction query string false "sort direction, default to desc, accepts - asc, desc"
// @Param all query bool false "fetch all. default to false"
// @Param types query string false "Comma-separated market's type. Accepts [1,2,3]"
// @Param statuses query string false "Comma-separated market's status. Accepts [1,2,3]"
// @Tags Market
// @Accept json
// @Produce json
// @Success 200 {object} ListMarketsResponse
// @Failure 400,401,500 {object} model.ErrorResponse
func (controller *MarketController) ListMarkets(ctx *gin.Context) {
	httpRequest := common.AsPageRequest(ctx)

	searchTerm := ctx.Request.URL.Query().Get("search_term")

	grpcRequest := &market.ListMarketsRequest{
		PageRequest: &grpc.PageRequest{
			Page:      httpRequest.Page,
			Size:      httpRequest.Size,
			Sort:      httpRequest.Sort,
			Direction: httpRequest.Direction,
		},
		SearchTerm: &searchTerm,
	}

	if searchTerm != "" {
		grpcRequest.SearchTerm = &searchTerm
	}

	marketTypeStr := ctx.Query("types")
	if marketTypeStr != "" {
		tmuStrList := strings.Split(marketTypeStr, ",")
		tmuNumList := make([]market.MarketType, len(tmuStrList))
		for i := 0; i < len(tmuStrList); i++ {
			tmuNum, err := strconv.ParseInt(tmuStrList[i], 10, 32)
			if err != nil {
				continue
			}
			result := int32(tmuNum)
			tmuNumList[i] = market.MarketType(result)
		}

		grpcRequest.Types = tmuNumList
	}

	marketStatusStr := ctx.Query("statuses")
	if marketStatusStr != "" {
		tmuStrList := strings.Split(marketStatusStr, ",")
		tmuNumList := make([]market.MarketStatus, len(tmuStrList))
		for i := 0; i < len(tmuStrList); i++ {
			tmuNum, err := strconv.ParseInt(tmuStrList[i], 10, 32)
			if err != nil {
				continue
			}
			result := int32(tmuNum)
			tmuNumList[i] = market.MarketStatus(result)
		}

		grpcRequest.Statuses = tmuNumList
	}

	if ctx.Query("all") != "" {
		fetchAll, castErr := common.ConvertToBool(ctx.Query("all"))
		if castErr != nil {
			log.Println("Cannot parse all")
			common.ReturnErrorResponse(ctx, http.StatusInternalServerError, castErr.Error())
			return
		}
		grpcRequest.All = fetchAll
	}

	res, err := controller.marketClient.ListMarkets(grpcRequest, common.GetMetadataFromContext(ctx))

	if err != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if res.Success {
		var items []market.Market
		for _, value := range res.GetData().GetItems() {
			var item market.Market
			_ = value.UnmarshalTo(&item)
			items = append(items, item)
		}

		ctx.JSON(http.StatusOK, &ListMarketsResponse{
			Page:          res.GetData().GetPage(),
			Size:          res.GetData().GetSize(),
			TotalElements: res.GetData().GetTotalElements(),
			TotalPages:    res.GetData().GetTotalPages(),
			Items:         items,
		})
	} else {
		common.AsErrorResponse(res.GetError(), ctx)
	}
}

type ListPublishedMarketsResponse struct {
	TotalElements int64            `json:"total_elements"`
	Items         []*market.Market `json:"items"`
}

// GetMarket
// @Router /api/v2/markets/{id} [GET]
// @Summary Get market (for both edit & review)
// @Description Returns the draft (if draft = true) or primary (otherwise) version of markets
// @Param draft query bool false "should get the draft version or not, default to false"
// @Param id path string true "ID"
// @Tags Market
// @Accept json
// @Produce json
// @Success 200 {object} market.GetMarketResponse_Data
// @Failure 400,401,500 {object} model.ErrorResponse
func (controller *MarketController) GetMarket(ctx *gin.Context) {
	marketID := ctx.Param("id")

	draft := false

	if ctx.Query("draft") == "true" {
		draft = true
	}

	req := &market.GetMarketRequest{
		MarketId: marketID,
		Draft:    &draft,
	}

	res, err := controller.marketClient.GetMarket(req, common.GetMetadataFromContext(ctx))

	if err != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if res.Success {
		ctx.JSON(http.StatusOK, res.GetData().GetMarket())
	} else {
		common.AsErrorResponse(res.GetError(), ctx)
	}
}

// CreateMarket
// @Router /api/v2/markets [POST]
// @Summary Create market
// @Description For step 1
// @Param _ body market.UpsertMarketRequest true "request body"
// @Tags Market
// @Accept json
// @Produce json
// @Success 200 {object} common.OnlyIdResponse_Data
// @Failure 400,401,500 {object} model.ErrorResponse
func (controller *MarketController) CreateMarket(ctx *gin.Context) {
	req := new(market.UpsertMarketRequest)

	if err := ctx.ShouldBindJSON(req); err != nil {
		common.ReturnErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res, err := controller.marketClient.CreateMarket(req, common.GetMetadataFromContext(ctx))

	if err != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if res.Success {
		common.AsSuccessResponse(res.GetData(), ctx)
	} else {
		common.AsErrorResponse(res.GetError(), ctx)
	}
}

// UpdateMarket
// @Router /api/v2/markets/{id} [PUT]
// @Summary Update Market
// @Param id path string true "ID"
// @Param _ body market.UpsertMarketRequest true "Market Metadata"
// @Tags Market
// @Accept json
// @Produce json
// @Success 200 {object} market.UpdateMarketResponse_Data
// @Failure 400,401,500 {object} model.ErrorResponse
func (controller *MarketController) UpdateMarket(ctx *gin.Context) {
	req := new(market.UpsertMarketRequest)
	id := ctx.Param("id")
	if err := ctx.ShouldBindJSON(req); err != nil {
		common.ReturnErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	req.MarketId = id
	res, err := controller.marketClient.UpdateMarket(req, common.GetMetadataFromContext(ctx))

	if err != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if res.Success {
		common.AsSuccessResponse(res.GetData(), ctx)
	} else {
		common.AsErrorResponse(res.GetError(), ctx)
	}
}

type PublishMarketRequest struct {
	MarketId string `json:"market_id"`
}
