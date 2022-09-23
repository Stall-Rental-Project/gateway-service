package middleware

import (
	"gateway-service/common"
	"gateway-service/common/constants"
	"gateway-service/model"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	UnauthorizedTokenMissing = "Token is missing"
	UnauthorizedTokenInvalid = "The access token is not valid"
)

type AdditionalClaims struct {
	Permissions []string `json:"permissions"`
	RoleIds     []string `json:"role_ids"`
}

func (middleware *Middleware) AuthMiddleware(ctx *gin.Context) {
	authorization := ctx.GetHeader("Authorization")

	if authorization == "" {
		abort(ctx, UnauthorizedTokenMissing)
		return
	}

	tokenString := strings.Split(authorization, " ")[1]

	if isInvalidToken(tokenString) {
		abort(ctx, UnauthorizedTokenInvalid)
		return
	}

	principal := common.ParseToken(tokenString)

	md := make(model.CacheableMetadata)
	if cached, err := middleware.redisClient.Get(constants.RedisPrincipalPrefix+principal.ExternalId, ctx); err != nil {
		authResp, authErr := middleware.obtainUserPermissions(principal)
		if authErr != nil {
			abort(ctx, authErr.Error())
			return
		}

		principal.Permissions = authResp.Permissions
		principal.RoleIds = authResp.RoleIds

		grpcToken, genErr := common.GenerateJwtToken(jwt.SigningMethodHS256, principal, constants.GrpcJwtSecret)
		if genErr != nil {
			abort(ctx, genErr.Error())
			return
		}

		md.Put("Authorization", constants.GrpcJwtHeaderPrefix+grpcToken)

		cacheErr := middleware.cachePrincipal(principal, md, ctx)
		if cacheErr != nil {
			abort(ctx, cacheErr.Error())
			return
		}
	} else {
		if cacheErr := json.Unmarshal([]byte(cached), &md); cacheErr != nil {
			log.Println("Failed when unmarshalling metadata", cacheErr)
			abort(ctx, "Unmarshalling error "+cacheErr.Error())
		}
	}

	ctx.Set("md", md.AsGrpcMetadata())
	ctx.Next()
}

//func (middleware *Middleware) obtainUserPermissions(principal *model.GrpcPrincipal) (*AdditionalClaims, error) {
//	resp, err := middleware.permissionClient.ListUserPermissions(&grpc.FindByIdRequest{
//		Id: principal.UserId,
//	})
//
//	if err != nil {
//		log.Println("Failed when retrieving user permission")
//		return nil, err
//	} else if !resp.Success {
//		log.Println("Failed when retrieving user permission")
//		return nil, errors.New(resp.GetError().GetMessage())
//	}
//
//	var permissions []string
//	for _, value := range resp.GetData().GetRolePermission().GetPermissions() {
//		permissions = append(permissions, value.Code)
//	}
//
//	return &AdditionalClaims{
//		Permissions: permissions,
//		RoleIds:     resp.GetData().GetRolePermission().GetRoleIds(),
//	}, nil
//}
//

func isInvalidToken(tokenString string) bool {
	if tokenString == "" || tokenString == "null" {
		log.Println("No token string provided")
		return false
	}

	var secret []byte
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	claims, ok := token.Claims.(jwt.MapClaims)
	valid := claims.VerifyExpiresAt(time.Now().Unix(), true)

	return !(valid && ok)
}

func abort(ctx *gin.Context, desc string) {
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, model.ErrorResponse{
		ErrorCode:        http.StatusText(http.StatusUnauthorized),
		ErrorDescription: desc,
	})
}
