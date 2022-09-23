package server

import (
	"gateway-service/client"
	"gateway-service/config"
	"gateway-service/controller"
	_ "gateway-service/docs"
	"gateway-service/middleware"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
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

	userClient := client.NewUserClient(accountCC)

	cloudinaryClient := client.NewCloudinaryClient()

	authController := controller.NewAuthController(authClient)

	userController := controller.NewUserController(userClient)

	fileController := controller.NewFileController(cloudinaryClient)

	middlewares := middleware.NewMiddleware()

	auth := server.router.Group("/api/v2")
	auth.POST("/login", authController.Login)

	user := server.router.Group("/api/v2/users")
	user.Use(middlewares.GinMiddleware).Use(middlewares.AuthMiddleware)
	{
		user.POST("", userController.CreateUser)
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
