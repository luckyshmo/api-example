package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/luckyshmo/api-example/pkg/service"

	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
	"github.com/swaggo/gin-swagger/swaggerFiles"
	ginprometheus "github.com/zsais/go-gin-prometheus"

	_ "github.com/luckyshmo/api-example/docs"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	//Open-Api endpoints documentation using GIN swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	p := ginprometheus.NewPrometheus("gin")
	p.Use(router)

	api := router.Group("/api", h.limit) //JWT Auth
	{
		data := api.Group("/data")
		{
			data.POST("/", h.receiveData)
		}

		auth := api.Group("/auth")
		{
			auth.POST("/sign-up", h.signUp)
			auth.POST("/sign-in", h.signIn)
		}

		users := api.Group("/user", h.userIdentity)
		{
			users.GET("/", h.getUserList)
			users.GET("/:id", h.getUser)
		}
	}

	return router
}
