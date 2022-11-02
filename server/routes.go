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
	rentalCC, err := grpc.Dial(config.Config.RentalClientUrl, transportOption)
	marketCC, err := grpc.Dial(config.Config.MarketClientUrl, transportOption)

	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}

	authClient := client.NewAuthClient(accountCC)
	roleClient := client.NewRoleClient(accountCC)
	permissionClient := client.NewPermissionClient(accountCC)
	userClient := client.NewUserClient(accountCC)

	cloudinaryClient := client.NewCloudinaryClient()

	rateClient := client.NewRateClient(rentalCC)
	locationClient := client.NewLocationClient(marketCC)
	marketClient := client.NewMarketClient(marketCC)
	floorClient := client.NewFloorClient(marketCC)
	//controller
	authController := controller.NewAuthController(authClient)
	roleController := controller.NewRoleController(roleClient)
	userController := controller.NewUserController(userClient)
	permissionController := controller.NewPermissionController(permissionClient)

	locationController := controller.NewLocationController(locationClient)
	marketController := controller.NewMarketController(marketClient)
	floorController := controller.NewFloorController(floorClient)

	rateController := controller.NewRateController(rateClient)

	fileController := controller.NewFileController(cloudinaryClient)

	middlewares := middleware.NewMiddleware(permissionClient)

	auth := server.router.Group("/api/v2")
	auth.POST("/login", authController.Login)

	user := server.router.Group("/api/v2/users")
	user.Use(middlewares.GinMiddleware).Use(middlewares.AuthMiddleware)
	{
		user.GET("", common.HasAnyPermission([]string{
			constants.UserManagement,
		}), userController.ListUsers)
		user.GET("/:id", common.HasAnyPermission([]string{
			constants.UserManagement,
		}), userController.GetUser)
		user.POST("", common.HasAnyPermission([]string{
			constants.UserManagement,
		}), userController.CreateUser)
		user.GET("/current", userController.GetCurrentUser)
		user.GET("/public", userController.ListPublicUsersByEmail)

		user.PUT("/:id", common.HasAnyPermission([]string{
			constants.UserManagement,
		}), userController.UpdateUser)
		user.DELETE("/:id", common.HasAllPermissions([]string{
			constants.UserManagement,
		}), userController.DeleteUser)
	}

	role := server.router.Group("/api/v2/roles")
	role.Use(middlewares.GinMiddleware).Use(middlewares.AuthMiddleware)
	{
		role.POST("", common.HasAnyPermission([]string{
			constants.RoleManagement,
		}), roleController.CreateRole)
		role.DELETE("/:id", common.HasAllPermissions([]string{
			constants.RoleManagement,
		}), roleController.DeleteRole)
		role.PUT("/:id", common.HasAnyPermission([]string{
			constants.RoleManagement,
		}), roleController.UpdateRole)
		role.GET("", common.HasAnyPermission([]string{
			constants.RoleManagement,
		}), roleController.ListRoles)
		role.GET("/:id", common.HasAnyPermission([]string{
			constants.RoleManagement,
		}), roleController.GetRole)
	}

	permission := server.router.Group("/api/v2/permissions")
	permission.Use(middlewares.GinMiddleware).Use(middlewares.AuthMiddleware)
	{
		permission.GET("", permissionController.ListPermissions)
		permission.GET("/categories", permissionController.ListPermissionCategory)
	}

	rate := server.router.Group("/api/v2/rates")
	rate.Use(middlewares.GinMiddleware).Use(middlewares.AuthMiddleware)
	{
		rate.GET("", common.HasAnyPermission([]string{
			constants.RateManagement,
		}), rateController.ListRates)
		rate.GET("/:id", common.HasAnyPermission([]string{
			constants.RateManagement,
		}), rateController.GetRate)
		rate.POST("", common.HasAnyPermission([]string{
			constants.RateManagement,
		}), rateController.CreateRate)
		rate.PUT("/:id", common.HasAnyPermission([]string{
			constants.RateManagement,
		}), rateController.UpdateRate)
		rate.DELETE("/:id", common.HasAnyPermission([]string{
			constants.RateManagement,
		}), rateController.DeleteRate)
	}

	location := server.router.Group("/api/v2/locations")
	location.Use(middlewares.GinMiddleware)
	{
		location.GET("/provinces", locationController.ListProvinces)
		location.GET("/cities", locationController.ListCities)
		location.GET("/wards", locationController.ListWards)
		location.GET("/query", locationController.GetLocation)
	}

	market := server.router.Group("/api/v2/markets")
	market.Use(middlewares.GinMiddleware).Use(middlewares.AuthMiddleware)
	{
		market.GET("", common.HasAnyPermission([]string{
			constants.MarketView,
		}), marketController.ListMarkets)

		market.GET("/published", marketController.ListPublishedMarkets)

		market.GET("/:id", common.HasAnyPermission([]string{
			constants.MarketView,
			constants.ApplicationSubmit,
			constants.ApplicationView,
		}), marketController.GetMarket)
		market.POST("", common.HasAnyPermission([]string{
			constants.MarketAddUpdate,
		}), marketController.CreateMarket)
		market.PUT("/:id", common.HasAnyPermission([]string{
			constants.MarketAddUpdate,
		}), marketController.UpdateMarket)
		market.GET("/:id/floors", common.HasAnyPermission([]string{
			constants.MarketView,
			constants.ApplicationSubmit,
			constants.ApplicationView,
		}), floorController.ListFloors)

	}

	floor := server.router.Group("/api/v2/floors")
	floor.Use(middlewares.GinMiddleware).Use(middlewares.AuthMiddleware)
	{
		floor.POST("", common.HasAnyPermission([]string{
			constants.MarketAddUpdate,
		}), floorController.CreateFloor)
		floor.PUT("", common.HasAnyPermission([]string{
			constants.MarketAddUpdate,
		}), floorController.UpdateFloor)
		floor.GET("/:id", common.HasAnyPermission([]string{
			constants.MarketView,
		}), floorController.GetFloor)
		floor.DELETE("", common.HasAnyPermission([]string{
			constants.MarketDelete,
		}), floorController.DeleteFloor)
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
