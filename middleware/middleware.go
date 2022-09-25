package middleware

import (
	"gateway-service/client"
	"github.com/gin-gonic/gin"
)

type Middleware struct {
	permissionClient *client.PermissionClient
}

func NewMiddleware(
	permissionClient *client.PermissionClient,
) Middleware {
	return Middleware{
		permissionClient: permissionClient,
	}
}

func (middleware *Middleware) GinMiddleware(ctx *gin.Context) {
	ctx.Next()
}
