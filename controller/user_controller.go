package controller

import (
	"gateway-service/client"
	"gateway-service/client/account"
	grpc "gateway-service/client/common"
	"gateway-service/common"
	"gateway-service/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
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

type PageUsers struct {
	Page          int32          `json:"page"`
	Size          int32          `json:"size"`
	TotalElements int64          `json:"total_elements"`
	TotalPages    int32          `json:"total_pages"`
	Items         []account.User `json:"items"`
}

// ListUsers
// @Summary List Users
// @Tags User
// @Accept json
// @Param size query integer false "Page size. Default to 20"
// @Param page query integer false "Page number. Default to 1"
// @Param sort query string false "Sort field. Accepts: [id, email, status, username]. Default to id."
// @Param direction query string false "Sort direction. Accepts: [asc, desc]. Default to asc"
// @Param search_term query string false "Filter by external id or email or first name or last name"
// @Param search_email_only query string false "Filter by email only. Ignored if search_term not given"
// @Param include_public query string false "Include public users in response. Default to false"
// @Param status query string false "Filter by status"
// @Param roles query string false "Filter by roles. Accepts a comma-separated UUIDs."
// @Success 200 {object} PageUsers
// @Failure 401,400,500 {object} model.ErrorResponse
// @Router /api/v2/users [GET]
func (controller *UserController) ListUsers(ctx *gin.Context) {
	req := common.AsPageRequest(ctx)

	searchTerm := ctx.Query("search_term")
	status := ctx.Query("status")
	roles := ctx.Query("roles")
	includePublic := ctx.Query("include_public")

	grpcReq := &account.ListUsersRequest{
		PageRequest: &grpc.PageRequest{
			Page:      req.Page,
			Size:      req.Size,
			Sort:      req.Sort,
			Direction: req.Direction,
		},
		SearchTerm: searchTerm,
	}

	if "0" == status {
		status := int32(0)
		grpcReq.Status = (*grpc.Status)(&status)
	} else if "1" == status {
		status := int32(1)
		grpcReq.Status = (*grpc.Status)(&status)
	}

	if includePublic == "true" {
		grpcReq.IncludePublic = true
	} else {
		grpcReq.IncludePublic = false
	}

	if ctx.Query("search_email_only") == "true" {
		grpcReq.SearchEmailOnly = true
	}

	if roles != "" {
		roleStrList := strings.Split(roles, ",")
		grpcReq.RoleIds = roleStrList
	}

	res, err := controller.userClient.ListUsers(grpcReq, common.GetMetadataFromContext(ctx))

	if err != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if res.Success {
		var items []*account.User

		for _, value := range res.GetData().GetItems() {
			var item account.User
			_ = value.UnmarshalTo(&item)
			items = append(items, &item)
		}

		ctx.JSON(http.StatusOK, common.AsPageResponse(res, items))
	} else {
		ctx.JSON(http.StatusBadRequest, model.AsErrorResponse(res.GetError()))
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

// GetCurrentUser
// @Summary Get current User profile
// @Tags User
// @Accept json
// @Success 200 {object} model.User
// @Failure 401,400,500 {object} model.ErrorResponse
// @Router /api/v2/users/current [GET]
func (controller *UserController) GetCurrentUser(ctx *gin.Context) {
	var user *model.User

	//principal := common.GetPrincipalFromContext(ctx)

	resp, err := controller.userClient.GetCurrentUser(common.GetMetadataFromContext(ctx))

	if err != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if !resp.Success {
		ctx.JSON(http.StatusBadRequest, model.AsErrorResponse(resp.GetError()))
		return
	}

	grpcUser := resp.GetData().GetUser()

	user = controller.buildGetUserResponse(grpcUser, ctx)

	ctx.JSON(http.StatusOK, user)
}

type ListUsers struct {
	TotalElements int64           `json:"total_elements"`
	Items         []*account.User `json:"items"`
}

// DeleteUser
// @Summary Delete User
// @Tags User
// @Accept json
// @Param id path string true "User generated ID"
// @Success 200 {object} common.NoContentResponse
// @Failure 401,400,500 {object} model.ErrorResponse
// @Router /api/v2/users/{id} [DELETE]
func (controller *UserController) DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := controller.userClient.DeleteUser(&grpc.FindByCodeRequest{Code: id}, common.GetMetadataFromContext(ctx))

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

// UpdateUser
// @Summary Update User
// @Tags User
// @Accept json
// @Param id path string true "User external ID"
// @Param data body account.UpsertUserRequest true "User Information"
// @Success 200 {object} common.NoContentResponse
// @Failure 401,400,500 {object} model.ErrorResponse
// @Router /api/v2/users/{id} [PUT]
func (controller *UserController) UpdateUser(ctx *gin.Context) {
	req := new(account.UpsertUserRequest)

	if err := ctx.ShouldBindJSON(req); err != nil {
		common.ReturnErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res, err := controller.userClient.UpdateUser(req, common.GetMetadataFromContext(ctx))

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

// GetUser
// @Summary Get User
// @Tags User
// @Accept json
// @Param id path string true "User external ID"
// @Success 200 {object} model.User
// @Failure 401,400,500 {object} model.ErrorResponse
// @Router /api/v2/users/{id} [GET]
func (controller *UserController) GetUser(ctx *gin.Context) {
	externalUserId := ctx.Param("id")

	userResp, userErr := controller.userClient.GetUser(&grpc.FindByCodeRequest{
		Code: externalUserId,
	}, common.GetMetadataFromContext(ctx))

	if userErr != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, userErr.Error())
		return
	}

	if !userResp.Success {
		ctx.JSON(http.StatusBadRequest, model.AsErrorResponse(userResp.GetError()))
		return
	}

	grpcUser := userResp.GetData().GetUser()

	user := controller.buildGetUserResponse(grpcUser, ctx)

	ctx.JSON(http.StatusOK, user)
}

// ListPublicUsersByEmail
// @Summary List Public Users by email
// @Tags User
// @Accept json
// @Param email query string true "User email"
// @Success 200 {object} ListUsers
// @Failure 401,400,500 {object} model.ErrorResponse
// @Router /api/v2/users/public [GET]
func (controller *UserController) ListPublicUsersByEmail(ctx *gin.Context) {
	email := ctx.Query("email")

	if email == "" {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, "No email was given")
		return
	}

	res, err := controller.userClient.ListUsersByEmail(&grpc.FindByCodeRequest{
		Code: email,
	}, common.GetMetadataFromContext(ctx))

	if err != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if !res.GetSuccess() {
		ctx.JSON(http.StatusBadRequest, model.AsErrorResponse(res.GetError()))
		return
	}

	if res.Success {
		var items []*account.User
		for _, value := range res.GetData().GetItems() {
			var item account.User
			_ = value.UnmarshalTo(&item)
			items = append(items, &item)
		}

		ctx.JSON(http.StatusOK, ListUsers{
			TotalElements: int64(len(items)),
			Items:         items,
		})
	} else {
		ctx.JSON(http.StatusBadRequest, model.AsErrorResponse(res.GetError()))
	}
}

func (controller *UserController) buildGetUserResponse(grpcUser *account.User, ctx *gin.Context) *model.User {
	user := &model.User{
		UserId:      grpcUser.UserId,
		FirstName:   grpcUser.FirstName,
		MiddleName:  grpcUser.MiddleName,
		LastName:    grpcUser.LastName,
		Email:       grpcUser.Email,
		Status:      grpcUser.Status,
		Roles:       grpcUser.Roles,
		RoleIds:     grpcUser.RoleIds,
		RoleCodes:   grpcUser.RoleCodes,
		Permissions: grpcUser.Permissions,
		MarketCodes: grpcUser.MarketCodes,
	}

	//isPublicUser := common.IsPublicUser(grpcUser.GetRoleCodes())

	//user.Markets = controller.getAssociatedMarkets(grpcUser.GetMarketCodes(), isPublicUser, ctx)

	if user.Roles == nil {
		user.Roles = make([]string, 0)
	}

	if user.RoleCodes == nil {
		user.RoleCodes = make([]string, 0)
	}

	if user.Markets == nil {
		user.Markets = make([]model.UserMarket, 0)
	}

	return user
}
