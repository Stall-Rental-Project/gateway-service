package common

import (
	grpc "gateway-service/client/common"
	"gateway-service/model"
	"github.com/gin-gonic/gin"
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
