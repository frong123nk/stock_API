package routers

import (
	"stock/controllers"
	"stock/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(router *gin.Engine) {

	authenAPI := router.Group("/api/v2")
	{
		authenAPI.POST("/login", controllers.Login)
		authenAPI.POST("/register", controllers.Register)
	}

	productAPI := router.Group("/api/v2")
	{
		productAPI.GET("/product", middleware.AuthMiddleware, controllers.GetProduct)
		productAPI.POST("/product", middleware.AuthMiddleware, controllers.CreateProduct)
		productAPI.PUT("/product", middleware.AuthMiddleware, controllers.EditProduct)
		productAPI.DELETE("/product", middleware.AuthMiddleware, controllers.DeleteProduct)
	}

	transactionAPI := router.Group("/api/v2")
	{
		transactionAPI.GET("/transaction", middleware.AuthMiddleware, controllers.GetTransaction)
		transactionAPI.POST("/transaction", middleware.AuthMiddleware, controllers.CreateTransaction)
	}
}
