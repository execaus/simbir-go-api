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

	transport := api.Group("/Transport")
	{
		transport.GET("/:id", h.GetTransport)
		transport.POST("", h.accountIdentity, h.CreateTransport)
		transport.PUT("/:id", h.accountIdentity, h.UpdateTransport)
		transport.DELETE("/:id", h.accountIdentity, h.DeleteTransport)
	}

	rent := api.Group("/Rent")
	{
		rent.GET("/Transport", h.GetRentTransport)
		rent.GET("/:id", h.accountIdentity, h.GetRent)
		rent.GET("/MyHistory", h.accountIdentity, h.GetRentMyHistory)
		rent.GET("/TransportHistory/:id", h.accountIdentity, h.GetRentTransportHistory)
		rent.POST("/New/:id", h.accountIdentity, h.CreateRent)
		rent.POST("/End/:id", h.accountIdentity, h.EndRent)
	}

	admin := api.Group("/Admin", h.onlyAdmin)
	{
		adminAccount := admin.Group("/Account")
		{
			adminAccount.GET("", h.AdminGetAccounts)
			adminAccount.GET("/:id", h.AdminGetAccount)
			adminAccount.POST("", h.AdminCreateAccount)
			adminAccount.PUT("/:id", h.AdminUpdateAccount)
			adminAccount.DELETE("/:id", h.AdminRemoveAccount)
		}

		adminTransport := admin.Group("/Transport")
		{
			adminTransport.GET("/", h.AdminGetTransports)
			adminTransport.GET("/:id", h.AdminGetTransport)
			adminTransport.POST("", h.AdminCreateTransport)
			adminTransport.PUT("/:id", h.AdminUpdateTransport)
			adminTransport.DELETE("/:id", h.AdminDeleteTransport)
		}

		adminRent := admin.Group("Rent")
		{
			adminRent.GET("/:id", h.GetAdminRent)
		}
		admin.GET("/UserHistory/:id", h.GetAdminUserRentHistory)
		admin.GET("/TransportHistory/:id", h.GetAdminTransportRentHistory)
	}

	return router
}
