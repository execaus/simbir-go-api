package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"simbir-go-api/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")

	account := api.Group("/Account")
	{
		account.GET("/Me", h.accountIdentity, h.GetAccount)
		account.POST("/SignUp", h.SignUp)
		account.POST("/SignIn", h.SignIn)
	}

	return router
}
