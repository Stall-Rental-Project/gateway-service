package controller

import (
	"gateway-service/client"
	"gateway-service/client/account"
	"gateway-service/common"
	"gateway-service/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserController struct {
	userClient *client.UserClient
}

func NewUserController(
	userClient *client.UserClient) UserController {
	return UserController{
		userClient: userClient,
	}
}

// CreateUser
// @Summary Create User
// @Tags User
// @Accept json
// @Param data body account.UpsertUserRequest true "User Information"
// @Success 200 {object} common.NoContentResponse
// @Failure 401,400,500 {object} model.ErrorResponse
// @Router /api/v2/users [POST]
func (controller *UserController) CreateUser(ctx *gin.Context) {
	req := new(account.UpsertUserRequest)

	if err := ctx.ShouldBindJSON(req); err != nil {
		common.ReturnErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res, err := controller.userClient.CreateUser(req, common.GetMetadataFromContext(ctx))

	if err != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if res.Success {
		ctx.JSON(http.StatusOK, res)
	} else {
		ctx.JSON(http.StatusBadRequest, model.AsErrorResponse(res.GetError()))
	}
}
