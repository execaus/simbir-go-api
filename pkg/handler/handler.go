package handler

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"simbir-go-api/pkg/service"
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
		account.PUT("/Update", h.accountIdentity, h.UpdateAccount)
		account.POST("/SignUp", h.SignUp)
		account.POST("/SignIn", h.SignIn)
		account.POST("/SignOut", h.accountIdentity, h.SignOut)
	}

	admin := api.Group("/Admin", h.onlyAdmin)
	{
		adminAccount := admin.Group("/Account")
		{
			adminAccount.GET("", h.AdminGetAccounts)
			adminAccount.GET("/:username", h.AdminGetAccount)
			adminAccount.POST("", h.AdminCreateAccount)
			adminAccount.PUT("/:username", h.AdminUpdateAccount)
			adminAccount.DELETE("/:username", h.AdminRemoveAccount)
		}
	}

	return router
}
