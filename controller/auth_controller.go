package controller

import (
	"gateway-service/client"
	"gateway-service/client/account"
	"gateway-service/common"
	"gateway-service/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type AuthController struct {
	authClient *client.AuthClient
}

func NewAuthController(
	accountClient *client.AuthClient,
) AuthController {
	return AuthController{
		authClient: accountClient,
	}
}

type TokenResponse struct {
	Message  string        `json:"message"`
	Response *TokenWrapper `json:"response"`
}

type TokenWrapper struct {
	StatusCode    string `json:"status_code"`
	StatusMessage string `json:"status_message"`
	AccessToken   string `json:"access_token"`
}

// Login
// @Summary Login
// @Description Get Access Token and Refresh Token
// @Tags Auth
// @Accept json
// @Param user body account.LoginRequest true "Login Information"
// @Success 200 {object} account.LoginResponse_Data
// @Failure 400,500 {object} model.ErrorResponse
// @Router /api/v2/login [POST]
func (controller *AuthController) Login(ctx *gin.Context) {
	req := new(account.LoginRequest)

	if err := ctx.ShouldBindJSON(req); err != nil {
		common.ReturnErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}
	log.Println("Login Request Info", req)
	if res, err := controller.authClient.Login(req); err != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	} else {
		if res.GetSuccess() {
			ctx.JSON(http.StatusOK, res.GetData())
		} else {
			ctx.JSON(http.StatusBadRequest, model.AsErrorResponse(res.GetError()))
		}
	}
}

// Logout
// @Summary Logout
// @Tags Auth
// @Accept json
// @Failure 400,500 {object} model.ErrorResponse
// @Router /api/v2/logout [POST]
func (controller *AuthController) Logout(ctx *gin.Context) {
	principal := common.GetPrincipalFromContext(ctx)

	if principal == nil {
		ctx.JSON(http.StatusOK, &model.SuccessResponse{
			Message: "Success",
		})
	}

	//controller.redisClient.Evict(constants.RedisPrincipalPrefix+principal.ExternalId, ctx)
	//controller.redisClient.Evict(constants.RedisCurrentUserPrefix+principal.ExternalId, ctx)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}
