package server

import (
	"gateway-service/client"
	"gateway-service/common"
	"gateway-service/common/constants"
	"gateway-service/config"
	"gateway-service/controller"
	_ "gateway-service/docs"
	"gateway-service/middleware"
	swaggerFiles "github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

/*
 * Routing API Configuration
 */

func (server *Server) InitializeRoutes() {
	// Init gRPC clients
	transportOption := grpc.WithTransportCredentials(insecure.NewCredentials())
	accountCC, err := grpc.Dial(config.Config.AccountClientUrl, transportOption)
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}
	authClient := client.NewAuthClient(accountCC)
	permissionClient := client.NewPermissionClient(accountCC)

	userClient := client.NewUserClient(accountCC)

	cloudinaryClient := client.NewCloudinaryClient()

	authController := controller.NewAuthController(authClient)

	userController := controller.NewUserController(userClient)

	fileController := controller.NewFileController(cloudinaryClient)

	middlewares := middleware.NewMiddleware(permissionClient)

	auth := server.router.Group("/api/v2")
	auth.POST("/login", authController.Login)

	user := server.router.Group("/api/v2/users")
	user.Use(middlewares.GinMiddleware).Use(middlewares.AuthMiddleware)
	{
		user.POST("", common.HasAnyPermission([]string{
			constants.UserManagement,
		}), userController.CreateUser)
	}

	file := server.router.Group("/api/v2/files")
	file.Use(middlewares.GinMiddleware).Use(middlewares.AuthMiddleware)
	{
		file.POST("/upload", fileController.UploadAttachment)
	}
	//server.router.GET("/static", fileController.GetStaticFile)
	//server.router.GET("/assets/images/:path", fileController.GetFile)

	server.router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

}
