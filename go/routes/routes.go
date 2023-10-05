package routes

import (
	"github.com/Kamalesh-Seervi/stripe-in-go/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Server() *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}                      // You can adjust this to allow specific origins
	config.AllowMethods = []string{"GET", "POST", "OPTIONS"} // You can specify which HTTP methods are allowed
	router.Use(cors.New(config))
	v1 := router.Group("/v1")
	{
		v1.GET("products", controllers.GetProducts)
		v1.POST("products", controllers.CreateProducts)
		v1.GET("config", controllers.Config)
		v1.POST("create-payment-intent", controllers.HandleCreatePaymentIntent)
	}
	router.Run(":8080")
	return router
}
