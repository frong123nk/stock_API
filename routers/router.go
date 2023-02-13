package routers

import (
	"stock/controllers"
	"stock/middleware"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func Setup(router *gin.Engine) {
	router.Use(CORSMiddleware())
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
