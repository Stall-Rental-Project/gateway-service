package controller

import (
	"gateway-service/client"
	"gateway-service/client/market"
	"gateway-service/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LocationController struct {
	locationClient *client.LocationClient
}

func NewLocationController(locationClient *client.LocationClient) LocationController {
	return LocationController{
		locationClient: locationClient,
	}
}

// ListProvinces
// @Router /api/v2/locations/provinces [GET]
// @Summary List provinces
// @Description List provinces
// @Param search_term query string false "search by name"
// @Tags Location
// @Accept json
// @Produce json
// @Success 200 {object} market.ListProvinceResponse_Data
// @Failure 400,401,500 {object} model.ErrorResponse
func (controller *LocationController) ListProvinces(ctx *gin.Context) {
	req := new(market.ListProvinceRequest)

	if err := ctx.ShouldBindJSON(req); err != nil {
		common.ReturnErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res, err := controller.locationClient.ListProvinces(req, common.GetMetadataFromContext(ctx))

	if err != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res.GetData())
}

// ListCities
// @Router /api/v2/locations/cities [GET]
// @Summary List provinces
// @Description List provinces
// @Param province query string true "province that cities belong to"
// @Param search_term query string false "search by name"
// @Tags Location
// @Accept json
// @Produce json
// @Success 200 {object} market.ListCityResponse_Data
// @Failure 400,401,500 {object} model.ErrorResponse
func (controller *LocationController) ListCities(ctx *gin.Context) {
	if ctx.Query("province") == "" {
		common.ReturnErrorResponse(ctx, http.StatusBadRequest, "missing required query - province")
	}

	req := &market.ListCityRequest{
		Province:   ctx.Query("province"),
		SearchTerm: ctx.Query("search_term"),
	}

	res, err := controller.locationClient.ListCities(req, common.GetMetadataFromContext(ctx))

	if err != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res.GetData())
}

// ListWards
// @Router /api/v2/locations/wards [GET]
// @Summary List wards
// @Description List wards
// @Param province query string true "province that wards belong to"
// @Param city query string true "city that wards belong to"
// @Param search_term query string false "search by name"
// @Tags Location
// @Accept json
// @Produce json
// @Success 200 {object} market.ListWardResponse_Data
// @Failure 400,401,500 {object} model.ErrorResponse
func (controller *LocationController) ListWards(ctx *gin.Context) {
	if ctx.Query("province") == "" {
		common.ReturnErrorResponse(ctx, http.StatusBadRequest, "missing required query - province")
		return
	} else if ctx.Query("city") == "" {
		common.ReturnErrorResponse(ctx, http.StatusBadRequest, "missing required query - city")
		return
	}

	req := &market.ListWardRequest{
		Province:   ctx.Query("province"),
		City:       ctx.Query("city"),
		SearchTerm: ctx.Query("search_term"),
	}

	res, err := controller.locationClient.ListWards(req, common.GetMetadataFromContext(ctx))

	if err != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res.GetData())
}

// GetLocation
// @Router /api/v2/locations/query [GET]
// @Summary Get specified location data
// @Description Query location by province, city & ward
// @Param province query string true "province that wards belong to"
// @Param city query string true "city that wards belong to"
// @Param ward query string true "ward specified"
// @Tags Location
// @Accept json
// @Produce json
// @Success 200 {object} market.GetLocationResponse_Data
// @Failure 400,401,500 {object} model.ErrorResponse
func (controller *LocationController) GetLocation(ctx *gin.Context) {
	if ctx.Query("province") == "" {
		common.ReturnErrorResponse(ctx, http.StatusBadRequest, "missing required query - province")
		return
	} else if ctx.Query("city") == "" {
		common.ReturnErrorResponse(ctx, http.StatusBadRequest, "missing required query - city")
		return
	} else if ctx.Query("barangay") == "" {
		common.ReturnErrorResponse(ctx, http.StatusBadRequest, "missing required query - ward")
		return
	}

	req := &market.GetLocationRequest{
		Province: ctx.Query("province"),
		City:     ctx.Query("city"),
		Ward:     ctx.Query("ward"),
	}

	res, err := controller.locationClient.GetLocation(req, common.GetMetadataFromContext(ctx))

	if err != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, res.GetData())
}
