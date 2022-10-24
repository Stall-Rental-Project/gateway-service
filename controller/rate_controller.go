package controller

import (
	"gateway-service/client"
	grpc "gateway-service/client/common"
	"gateway-service/client/rental"
	"gateway-service/common"
	"gateway-service/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
)

type RateController struct {
	rateClient *client.RateClient
}

func NewRateController(rateClient *client.RateClient) RateController {
	return RateController{
		rateClient: rateClient,
	}
}

type PageRates struct {
	Page          int32         `json:"page"`
	Size          int32         `json:"size"`
	TotalElements int64         `json:"total_elements"`
	TotalPages    int32         `json:"total_pages"`
	Items         []rental.Rate `json:"items"`
}

type Rate struct {
	RateId       string                    `json:"rate_id"`
	RateCode     string                    `json:"rate_code"`
	Status       grpc.Status               `json:"status"`
	Type         rental.RateType           `json:"type"`
	OtherRate    *rental.OtherRate         `json:"other_rate"`
	RentalRate   *rental.StallRentalRate   `json:"rental_rate"`
	RightsRate   *rental.StallRightsRate   `json:"rights_rate"`
	SecurityBond *rental.StallSecurityBond `json:"security_bond"`
}

// ListRates
// @Summary List Rates
// @Tags Rate
// @Accept json
// @Param size query integer false "Page size. Default to 20"
// @Param page query integer false "Page number. Default to 1"
// @Param sort query string false "Sort field. Accepts: [id, type, detail]. Default to id."
// @Param direction query string false "Sort direction. Accepts: [asc, desc]. Default to asc"
// @Param types query string false "Filter by types. Accept comma-separated numeric string"
// @Param details query string false "Filter by details. Accept comma-separated numeric string"
// @Success 200 {object} PageRates
// @Failure 401,400,500 {object} model.ErrorResponse
// @Router /api/v2/rates [GET]
func (controller *RateController) ListRates(ctx *gin.Context) {
	req := common.AsPageRequest(ctx)

	types := ctx.Query("types")
	details := ctx.Query("details")

	grpcReq := &rental.ListRatesRequest{
		PageRequest: &grpc.PageRequest{
			Page:      req.Page,
			Size:      req.Size,
			Sort:      req.Sort,
			Direction: req.Direction,
		},
	}

	if types != "" {
		typeStrList := strings.Split(types, ",")
		typeNumList := make([]rental.RateType, 0)
		for i := 0; i < len(typeStrList); i++ {
			tmuNum, err := strconv.ParseInt(typeStrList[i], 10, 32)
			if err != nil {
				continue
			}
			result := int32(tmuNum)
			typeNumList = append(typeNumList, rental.RateType(result))
		}
		grpcReq.Types = typeNumList
	}

	if details != "" {
		detailStrList := strings.Split(details, ",")
		detailNumList := make([]rental.OtherRateDetail, 0)
		for i := 0; i < len(detailStrList); i++ {
			tmuNum, err := strconv.ParseInt(detailStrList[i], 10, 32)
			if err != nil {
				continue
			}
			result := int32(tmuNum)
			detailNumList = append(detailNumList, rental.OtherRateDetail(result))
		}
		grpcReq.OtherRateDetails = detailNumList
	}

	res, err := controller.rateClient.ListRates(grpcReq, common.GetMetadataFromContext(ctx))

	if err != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if res.Success {
		var items []*rental.Rate

		for _, value := range res.GetData().GetItems() {
			var item rental.Rate
			_ = value.UnmarshalTo(&item)
			items = append(items, &item)
		}

		ctx.JSON(http.StatusOK, common.AsPageResponse(res, items))
	} else {
		ctx.JSON(http.StatusBadRequest, model.AsErrorResponse(res.GetError()))
	}
}

// GetRate
// @Summary Get Rate
// @Tags Rate
// @Accept json
// @Param id path string true "Rate generated ID"
// @Success 200 {object} Rate
// @Failure 401,400,500 {object} model.ErrorResponse
// @Router /api/v2/rates/{id} [GET]
func (controller *RateController) GetRate(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := controller.rateClient.GetRate(&rental.GetRateRequest{
		RateCode: id,
	}, common.GetMetadataFromContext(ctx))

	if err != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if res.Success {
		grpcRate := res.GetData().Rate
		rate := &Rate{
			RateId:   grpcRate.RateId,
			RateCode: grpcRate.RateCode,
			Status:   grpcRate.Status,
			Type:     grpcRate.Type,
		}

		if grpcRate.Type == rental.RateType_OTHER_RATES {
			rate.OtherRate = grpcRate.GetOtherRate()
		} else if grpcRate.Type == rental.RateType_STALL_RIGHTS_RATE {
			rate.RightsRate = grpcRate.GetRightsRate()
		} else if grpcRate.Type == rental.RateType_STALL_RENTAL_RATE {
			rate.RentalRate = grpcRate.GetRentalRate()
		}

		ctx.JSON(http.StatusOK, rate)
	} else {
		ctx.JSON(http.StatusBadRequest, model.AsErrorResponse(res.GetError()))
	}
}

// CreateRate
// @Summary Create Rate
// @Tags Rate
// @Accept json
// @Param data body rental.UpsertRateRequest true "Rate Information"
// @Success 200 {object} common.OnlyCodeResponse_Data
// @Failure 401,400,500 {object} model.ErrorResponse
// @Router /api/v2/rates [POST]
func (controller *RateController) CreateRate(ctx *gin.Context) {
	req := new(rental.UpsertRateRequest)

	if err := ctx.ShouldBindJSON(req); err != nil {
		common.ReturnErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res, err := controller.rateClient.CreateRate(req, common.GetMetadataFromContext(ctx))

	if err != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if res.Success {
		ctx.JSON(http.StatusOK, res.GetData())
	} else {
		ctx.JSON(http.StatusBadRequest, model.AsErrorResponse(res.GetError()))
	}
}

// UpdateRate
// @Summary Update Rate
// @Tags Rate
// @Accept json
// @Param id path string true "Rate generated ID"
// @Param data body rental.UpsertRateRequest true "Rate Information"
// @Success 200 {object} common.OnlyCodeResponse_Data
// @Failure 401,400,500 {object} model.ErrorResponse
// @Router /api/v2/rates/{id} [PUT]
func (controller *RateController) UpdateRate(ctx *gin.Context) {
	req := new(rental.UpsertRateRequest)

	id := ctx.Param("id")

	if err := ctx.ShouldBindJSON(req); err != nil {
		common.ReturnErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	req.RateId = &id

	res, err := controller.rateClient.UpdateRate(req, common.GetMetadataFromContext(ctx))

	if err != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if res.Success {
		ctx.JSON(http.StatusOK, res.GetData())
	} else {
		ctx.JSON(http.StatusBadRequest, model.AsErrorResponse(res.GetError()))
	}
}

// DeleteRate
// @Summary Delete Rate
// @Description Delete a Rate
// @Tags Rate
// @Accept json
// @Param id path string true "ID"
// @Success 200 {object} model.SuccessResponse
// @Failure 401,400,500 {object} model.ErrorResponse
// @Router /api/v2/rates/{id} [DELETE]
func (controller *RateController) DeleteRate(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := controller.rateClient.DeleteRate(&grpc.FindByIdRequest{Id: id}, common.GetMetadataFromContext(ctx))

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
