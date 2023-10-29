package handler

import (
	"github.com/execaus/exloggo"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"simbir-go-api/input_validator"
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

	setFieldValidator()

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
			adminRent.GET("/:id", h.AdminGetRent)
			adminRent.POST("", h.AdminCreateRent)
			adminRent.POST("/End/:id", h.AdminEndRent)
			adminRent.PUT("/:id", h.AdminUpdateRent)
		}
		admin.GET("/UserHistory/:id", h.AdminGetUserRentHistory)
		admin.GET("/TransportHistory/:id", h.AdminGetTransportRentHistory)
	}

	return router
}

func setFieldValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		if err := v.RegisterValidation("iso8601", input_validator.IsISO8601Date); err != nil {
			exloggo.Fatal(err.Error())
		}
	}
}
