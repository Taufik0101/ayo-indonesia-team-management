package routes

import (
	"gin-ayo/controller"
	"gin-ayo/middleware"
	repositories "gin-ayo/repository"
	"gin-ayo/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

type route struct {
	db *gorm.DB
}

func (r *route) SetupRoutes() *gin.Engine {
	router := gin.Default()

	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, PATCH")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Access-Control-Allow-Origin, Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			log.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	})

	r.Handler(router)

	return router
}

type RouteInterface interface {
	SetupRoutes() *gin.Engine
}

func NewRoute(
	db *gorm.DB,
) RouteInterface {
	return &route{
		db: db,
	}
}

func (r *route) Handler(router *gin.Engine) {
	// Repository
	userRepository := repositories.NewUserRepository(r.db)
	teamRepository := repositories.NewTeamRepository(r.db)
	playerRepository := repositories.NewPlayerRepository(r.db)
	scheduleRepository := repositories.NewScheduleRepository(r.db)
	resultRepository := repositories.NewResultRepository(r.db)
	detailResultRepository := repositories.NewDetailResultRepository(r.db)

	// Service
	authService := service.NewAuthService(userRepository)
	uploadService := service.NewUploadService()
	teamService := service.NewTeamService(teamRepository)
	playerService := service.NewPlayerService(playerRepository)
	scheduleService := service.NewScheduleService(scheduleRepository, resultRepository, playerRepository)
	resultService := service.NewResultService(resultRepository, detailResultRepository, scheduleRepository)

	// Controller
	authController := controller.NewAuthController(authService)
	uploadController := controller.NewUploadController(uploadService)
	teamController := controller.NewTeamController(teamService)
	playerController := controller.NewPlayerController(playerService)
	scheduleController := controller.NewScheduleController(scheduleService)
	resultController := controller.NewResultController(resultService)

	// Routes
	auth := router.Group("/auth")
	{
		auth.POST("/login", authController.Login)
	}

	routeMiddleware := router.Group("/api", middleware.AuthMiddleware())

	routeMiddleware.POST("/upload", uploadController.UploadFile)

	teamRoute := routeMiddleware.Group("/team")
	{
		teamRoute.POST("/create", teamController.Create)
		teamRoute.POST("/update", teamController.Update)
		teamRoute.DELETE("/delete/:id", teamController.Delete)
	}

	playerRoute := routeMiddleware.Group("/player")
	{
		playerRoute.POST("/create", playerController.Create)
		playerRoute.POST("/update", playerController.Update)
		playerRoute.DELETE("/delete/:id", playerController.Delete)
	}

	scheduleRoute := routeMiddleware.Group("/schedule")
	{
		scheduleRoute.POST("/create", scheduleController.Create)
		scheduleRoute.POST("/update", scheduleController.Update)
		scheduleRoute.DELETE("/delete/:id", scheduleController.Delete)
		scheduleRoute.GET("/detail/:id", scheduleController.Detail)
	}

	resultRoute := routeMiddleware.Group("/result")
	{
		resultRoute.POST("/create", resultController.Create)
	}
}
