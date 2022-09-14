package middleware

import (
	"github.com/gin-gonic/gin"
)

type Middleware struct {
}

func NewMiddleware() Middleware {
	return Middleware{}
}

func (middleware *Middleware) GinMiddleware(ctx *gin.Context) {
	ctx.Next()
}
