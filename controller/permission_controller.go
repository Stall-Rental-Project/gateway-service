package controller

import (
	"gateway-service/client"
	"gateway-service/client/account"
	"gateway-service/common"
	"gateway-service/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PermissionController struct {
	permissionClient *client.PermissionClient
}

func NewPermissionController(permissionClient *client.PermissionClient) PermissionController {
	return PermissionController{
		permissionClient: permissionClient,
	}
}

type Permission struct {
	PermissionId string `json:"permission_id"`
	Code         string `json:"code"`
	Name         string `json:"name"`
	Category     string `json:"category"`
}

type Permissions struct {
	Permissions []Permission `json:"permissions"`
}

// ListPermissions
// @Router /api/v2/permissions [GET]
// @Param roles query string false "Filter by roles. Accepts a comma-separated string."
// @Summary List all permissions
// @Tags Permissions
// @Accept json
// @Produce json
// @Success 200 {object} Permissions
// @Failure 400,401,500 {object} model.ErrorResponse
func (controller *PermissionController) ListPermissions(ctx *gin.Context) {
	roles := ctx.Query("roles")

	res, err := controller.permissionClient.ListPermissions(&account.ListPermissionsRequest{RoleCodes: roles})

	if err != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	} else if !res.Success {
		common.AsErrorResponse(res.GetError(), ctx)
		return
	}

	var items []account.Permission
	for _, value := range res.GetData().GetItems() {
		var item account.Permission
		_ = value.UnmarshalTo(&item)
		items = append(items, item)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"permissions": items,
	})
}

type ListPermissionCategories struct {
	Items         []account.PermissionCategory `json:"items"`
	TotalElements int64                        `json:"total_elements"`
}

// ListPermissionCategory
// @Summary List Permission Category
// @Description Get permission categories
// @Tags Role
// @Accept json
// @Success 200 {object} model.PageResponse
// @Failure 401,400,500 {object} model.ErrorResponse
// @Router /api/v2/permissions/categories [GET]
func (controller *PermissionController) ListPermissionCategory(ctx *gin.Context) {
	res, err := controller.permissionClient.ListPermissionCategories()
	if err != nil {
		common.ReturnErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	if res.GetSuccess() {
		var items []account.PermissionCategory
		for _, value := range res.GetData().GetItems() {
			var item account.PermissionCategory
			_ = value.UnmarshalTo(&item)
			items = append(items, item)
		}

		ctx.JSON(http.StatusOK, &ListPermissionCategories{
			Items:         items,
			TotalElements: res.GetData().GetTotalElements(),
		})
	} else {
		ctx.JSON(http.StatusBadRequest, model.AsErrorResponse(res.GetError()))
	}
}
