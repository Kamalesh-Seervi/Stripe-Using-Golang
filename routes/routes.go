package routes

import (
	"github.com/Kamalesh-Seervi/stripe-in-go/models"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Server() {
	router := gin.Default()

	router.Use(cors.Default())
	v1 := router.Group("/v1")
	{
		v1.GET("products", models.GetAllProducts)
		v1.POST("products", models.CreateProduct)
		v1.GET("config", Controllers.Config)
		v1.POST("create-payment-intent", Controllers.HandleCreatePaymentIntent)
	}
	router.Run(":8080")
}
