package controller

import (
	"gateway-service/client"
	"gateway-service/client/account"
	grpc "gateway-service/client/common"
	"gateway-service/common"
	"gateway-service/model"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type RoleController struct {
	roleClient *client.RoleClient
}

func NewRoleController(roleClient *client.RoleClient,
) RoleController {
	return RoleController{
		roleClient: roleClient,
	}
}

// CreateRole
// @Summary Create Role
// @Description Create new role
// @Tags Role
// @Accept json
// @Param data body account.UpsertRoleRequest true "Role Information"
// @Success 200 {object} common.NoContentResponse
// @Failure 401,400,500 {object} model.ErrorResponse
// @Router /api/v2/roles [POST]
func (controller *RoleController) CreateRole(ctx *gin.Context) {
	roleRequest := new(account.UpsertRoleRequest)
	if err := ctx.ShouldBindJSON(roleRequest); err != nil {
		common.ReturnErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	res, err := controller.roleClient.CreateRole(roleRequest, common.GetMetadataFromContext(ctx))
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

// UpdateRole
// @Summary Update Role
// @Description Update a role
// @Tags Role
// @Accept json
// @Param id path string true "ID"
// @Param data body account.UpsertRoleRequest true "Role Information"
// @Success 200 {object} common.NoContentResponse
// @Failure 401,400,500 {object} model.ErrorResponse
// @Router /api/v2/roles/{id} [PUT]
func (controller *RoleController) UpdateRole(ctx *gin.Context) {
	id := ctx.Param("id")

	req := new(account.UpsertRoleRequest)

	if err := ctx.ShouldBindJSON(req); err != nil {
		common.ReturnErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	req.RoleId = &id

	res, err := controller.roleClient.UpdateRole(req, common.GetMetadataFromContext(ctx))

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

// DeleteRole
// @Summary Delete Role
// @Description Delete a role
// @Tags Role
// @Accept json
// @Param id path string true "ID"
// @Success 200 {object} model.SuccessResponse
// @Failure 401,400,500 {object} model.ErrorResponse
// @Router /api/v2/roles/{id} [DELETE]
func (controller *RoleController) DeleteRole(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := controller.roleClient.DeleteRole(&grpc.FindByIdRequest{Id: id}, common.GetMetadataFromContext(ctx))

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

type PageRoles struct {
	Page          int32          `json:"page"`
	Size          int32          `json:"size"`
	TotalElements int64          `json:"total_elements"`
	TotalPages    int32          `json:"total_pages"`
	Items         []account.Role `json:"items"`
}

// ListRoles
// @Summary List Role
// @Tags Role
// @Accept json
// @Param size query integer false "Page size"
// @Param page query integer false "Page number"
// @Param sort query string false "Sort field"
// @Param direction query string false "Sort direction"
// @Param name query string false "Search By Name"
// @Param status query string false "Filter by status"
// @Param include_public query bool false "Should include public user role or not. Default to false"
// @Success 200 {object} PageRoles
// @Failure 401,400,500 {object} model.ErrorResponse
// @Router /api/v2/roles [GET]
func (controller *RoleController) ListRoles(ctx *gin.Context) {
	pageRequest := common.AsPageRequest(ctx)

	name := ctx.Query("name")
	status := ctx.Query("status")

	includePublicStr := ctx.Query("include_public")
	var includePublic bool
	if includePublicStr != "" {
		includePublic, _ = strconv.ParseBool(includePublicStr)
	} else {
		includePublic = false
	}

	req := &account.ListRoleRequest{
		PageRequest:   &pageRequest,
		Name:          name,
		IncludePublic: includePublic,
	}

	if status != "" {
		v, _ := strconv.ParseInt(status, 10, 32)
		vInt32 := int32(v)
		req.Status = (*grpc.Status)(&vInt32)
	}

	res, err := controller.roleClient.ListRoles(req, common.GetMetadataFromContext(ctx))

	if err != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if res.GetSuccess() {
		var items []*account.Role

		for _, value := range res.GetData().GetItems() {
			var item account.Role
			_ = value.UnmarshalTo(&item)
			items = append(items, &item)
		}

		ctx.JSON(http.StatusOK, common.AsPageResponse(res, items))
	} else {
		ctx.JSON(http.StatusBadRequest, model.AsErrorResponse(res.GetError()))
	}
}

// GetRole
// @Summary Get Role
// @Description Get a role
// @Tags Role
// @Accept json
// @Success 200 {object} account.Role
// @Failure 401,400,500 {object} model.ErrorResponse
// @Router /api/v2/roles/{id} [GET]
func (controller *RoleController) GetRole(ctx *gin.Context) {
	id := ctx.Param("id")

	res, err := controller.roleClient.GetRole(&grpc.FindByIdRequest{Id: id}, common.GetMetadataFromContext(ctx))

	if err != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	if res.GetSuccess() {
		ctx.JSON(http.StatusOK, res.GetData().GetRole())
	} else {
		ctx.JSON(http.StatusBadRequest, model.AsErrorResponse(res.GetError()))
	}
}
