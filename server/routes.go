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
	terminationClient := client.NewTerminationClient(rentalCC)
	leaseClient := client.NewLeaseClient(rentalCC)

	cloudinaryClient := client.NewCloudinaryClient()
	nsaClient := client.NewNSAClient(rentalCC)
	applicationClient := client.NewApplicationClient(rentalCC)

	rateClient := client.NewRateClient(rentalCC)
	locationClient := client.NewLocationClient(marketCC)
	marketClient := client.NewMarketClient(marketCC)
	floorClient := client.NewFloorClient(marketCC)
	stallClient := client.NewStallClient(marketCC)
	//controller
	authController := controller.NewAuthController(authClient)
	roleController := controller.NewRoleController(roleClient)
	userController := controller.NewUserController(userClient)
	permissionController := controller.NewPermissionController(permissionClient)

	locationController := controller.NewLocationController(locationClient)
	marketController := controller.NewMarketController(marketClient, nsaClient)
	floorController := controller.NewFloorController(floorClient, nsaClient)
	stallController := controller.NewStallController(stallClient, rateClient, nsaClient)

	rateController := controller.NewRateController(rateClient)
	nsaController := controller.NewNSAController(nsaClient, stallClient,
		marketClient, userClient, rateClient)
	applicationController := controller.NewApplicationController(applicationClient, stallClient)
	terminationController := controller.NewTerminationController(terminationClient)
	leaseController := controller.NewLeaseController(leaseClient, stallClient, marketClient, userClient, rateClient)

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
		market.POST("/:id/publish", common.HasAnyPermission([]string{
			constants.MarketApprovePublish,
		}), marketController.PublishMarket)

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
		market.DELETE("/:id", common.HasAnyPermission([]string{
			constants.MarketAddUpdate,
		}), marketController.DeleteMarket)
		market.GET("/:id/floors", common.HasAnyPermission([]string{
			constants.MarketView,
			constants.ApplicationSubmit,
			constants.ApplicationView,
		}), floorController.ListFloors)
		market.GET("/:id/stalls/count", marketController.CountStalls)
		market.GET("/:id/floors/published", floorController.ListPublishedFloors)
	}

	floor := server.router.Group("/api/v2/floors")
	floor.Use(middlewares.GinMiddleware).Use(middlewares.AuthMiddleware)
	{
		floor.POST("", common.HasAnyPermission([]string{
			constants.MarketAddUpdate,
		}), floorController.CreateFloor)
		floor.PUT("/:id", common.HasAnyPermission([]string{
			constants.MarketAddUpdate,
		}), floorController.UpdateFloor)
		floor.GET("/:id", common.HasAnyPermission([]string{
			constants.MarketView,
		}), floorController.GetFloor)
		floor.DELETE("/:id", common.HasAnyPermission([]string{
			constants.MarketDelete,
		}), floorController.DeleteFloor)
		floor.GET("/:id/published", floorController.GetPublishedFloor)

	}

	stall := server.router.Group("/api/v2/stalls")
	stall.Use(middlewares.GinMiddleware).Use(middlewares.AuthMiddleware)
	{
		stall.POST("", common.HasAnyPermission([]string{
			constants.MarketAddUpdate,
		}), stallController.CreateStall)
		stall.PUT("/:id/metadata", common.HasAnyPermission([]string{
			constants.MarketAddUpdate,
		}), stallController.UpdateStallMetadata)
		stall.PUT("/:id/position", common.HasAnyPermission([]string{
			constants.MarketAddUpdate,
		}), stallController.UpdateStallPosition)
		stall.GET("/:id", common.HasAnyPermission([]string{
			constants.MarketView,
			constants.ApplicationSubmit,
			constants.ApplicationView,
		}), stallController.GetStall)
		stall.DELETE("/:id", common.HasAnyPermission([]string{
			constants.MarketAddUpdate,
		}), stallController.DeleteStall)
		stall.GET("/:id/published", stallController.GetPublishedStall)
		stall.GET("/info", stallController.GetStallInfo)
	}

	application := server.router.Group("/api/v2/applications")
	application.Use(middlewares.GinMiddleware).Use(middlewares.AuthMiddleware)
	{
		application.GET("", common.HasAnyPermission([]string{
			constants.ApplicationView,
		}), applicationController.ListApplication)

		application.GET("/:id", common.HasAnyPermission([]string{
			constants.ApplicationView,
		}), nsaController.GetApplication)

		application.POST("", common.HasAnyPermission([]string{
			constants.ApplicationSubmit,
		}), nsaController.SubmitRequest)

		application.PUT("/:id/docs", common.HasAnyPermission([]string{
			constants.ApplicationSubmit,
		}), nsaController.SubmitDocs)

		application.PUT("/:id", common.HasAnyPermission([]string{
			constants.ApplicationSubmit,
		}), nsaController.UpdateApplication)
		application.PUT("/:id/payment", common.HasAnyPermission([]string{
			constants.ApplicationSubmit,
		}), nsaController.SubmitPayment)
		application.PUT("/:id/confirm", common.HasAnyPermission([]string{
			constants.ApplicationSubmit,
		}), nsaController.ConfirmApplication)

	}

	termination := server.router.Group("/api/v2/applications/:id/termination")
	termination.Use(middlewares.GinMiddleware).Use(middlewares.AuthMiddleware)
	{
		termination.GET("", terminationController.GetLeaseTermination)
		termination.POST("", common.HasAnyPermission([]string{
			constants.RequestTerminateLease,
			constants.TerminateLease,
		}), terminationController.SubmitLeaseTermination)
		termination.PUT("/:tid", common.HasAnyPermission([]string{
			constants.TerminateLease,
			constants.CancelTerminationRequest,
		}), terminationController.CancelLeaseTermination)
	}

	lease := server.router.Group("/api/v2/applications/in-lease")
	lease.Use(middlewares.GinMiddleware).Use(middlewares.AuthMiddleware)
	{
		lease.GET("", leaseController.ListLeases)
		lease.GET("/:id", leaseController.GetLease)
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
