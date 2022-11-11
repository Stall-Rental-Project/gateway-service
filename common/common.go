package common

import (
	grpc "gateway-service/client/common"
	"gateway-service/model"
	"github.com/gin-gonic/gin"
	"github.com/iancoleman/strcase"
	"net/http"
	"strconv"
)

func ReturnErrorResponse(ctx *gin.Context, code int, desc string) {
	codeStr := ""

	switch code {
	case 400:
		codeStr = grpc.ErrorCode_BAD_REQUEST.String()
		break
	default:
		codeStr = grpc.ErrorCode_INTERNAL_SERVER_ERROR.String()
		break
	}

	ctx.JSON(code, model.ErrorResponse{
		ErrorCode:        codeStr,
		ErrorDescription: desc,
	})
}

func AsPageRequest(ctx *gin.Context) grpc.PageRequest {
	// Initializing default
	size := 20
	page := 1
	sort := ""
	direction := "asc"
	query := ctx.Request.URL.Query()
	for key, value := range query {
		queryValue := value[len(value)-1]
		switch key {
		case "size":
			size, _ = strconv.Atoi(queryValue)
			break
		case "page":
			page, _ = strconv.Atoi(queryValue)
			break
		case "sort":
			sort = strcase.ToLowerCamel(queryValue)
			break
		case "direction":
			direction = queryValue
			break
		}
	}
	return grpc.PageRequest{
		Size:      int32(size),
		Page:      int32(page),
		Sort:      sort,
		Direction: direction,
	}
}

func AsErrorResponse(error *grpc.Error, ctx *gin.Context) {
	if error == nil {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{
			ErrorCode:        grpc.ErrorCode_INTERNAL_SERVER_ERROR.String(),
			ErrorDescription: "Unknown error",
		})
		return
	}

	if len(error.Details) > 0 {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{
			ErrorCode:        error.Code.String(),
			ErrorDescription: error.Message,
			Errors:           error.Details,
			Exception:        error.Exception,
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{
			ErrorCode:        error.Code.String(),
			ErrorDescription: error.Message,
			Errors:           error.Details,
			Exception:        error.Exception,
		})
	}
}

func AsPageResponse(pageResponse *grpc.PageResponse, items interface{}) (result model.PageResponse) {
	result.Page = pageResponse.GetData().Page
	result.Size = pageResponse.GetData().Size
	result.TotalPages = pageResponse.GetData().TotalPages
	result.TotalElements = pageResponse.GetData().TotalElements
	result.Items = items
	return
}

func AsSuccessResponse(data interface{}, ctx *gin.Context) {
	ctx.JSON(http.StatusOK, data)
}
func AsLegacyErrorResponse(error *grpc.ErrorResponse, ctx *gin.Context) {
	if len(error.Errors) > 0 {
		ctx.JSON(http.StatusBadRequest, model.ErrorResponse{
			ErrorCode:        grpc.ErrorCode_BAD_REQUEST.String(),
			ErrorDescription: error.ErrorDescription,
			Errors:           error.Errors,
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, model.ErrorResponse{
			ErrorCode:        grpc.ErrorCode_INTERNAL_SERVER_ERROR.String(),
			ErrorDescription: error.ErrorDescription,
			Errors:           error.Errors,
		})
	}
}
