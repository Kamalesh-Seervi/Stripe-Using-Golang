package controllers

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/Kamalesh-Seervi/stripe-in-go/models"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/paymentintent"
)

func GetProducts(c *gin.Context) {
	products, err := models.GetAllProducts()
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "No products found"})
	} else {
		c.JSON(http.StatusOK, gin.H{"products": products})
	}
}

func CreateProducts(c *gin.Context) {
	var product models.Product
	c.BindJSON(&product)
	err := models.CreateProduct(product)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		models.Dbase.Last(&product)
		c.JSON(http.StatusOK, gin.H{"product": product})
	}
}

func Config(c *gin.Context) {
    // Fetch the Stripe public key from the environment variable
    publicKey := os.Getenv("STRIPE_PUBLISHABLE_KEY")

    if publicKey == "" {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "STRIPE_PUBLISHABLE_KEY not set"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"publicKey": publicKey})
}


func HandleCreatePaymentIntent(c *gin.Context) {
	var product models.Product
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error in Loading .env file")
	}

	// Initialize Stripe with your secret key
	stripeKey := os.Getenv("STRIPE_SECRET_KEY")
	stripe.Key = stripeKey

	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Printf("ShouldBindJSON: %v", err)
		return
	}

	// Replace this with the actual product ID you want to use
	productID := strconv.FormatUint(uint64(product.Id), 10)
	data, err := models.GetProductById(&product, productID)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	amount := int64(data.Price)
	currency := stripe.Currency(stripe.CurrencyUSD)

	params := &stripe.PaymentIntentParams{
		Amount:             stripe.Int64(amount),
		Currency:           stripe.String(string(currency)),
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}), // Specify the payment method type(s) you want to use
		// Other fields as needed for your PaymentIntent
	}

	pi, err := paymentintent.New(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		log.Printf("pi.New: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"clientSecret": pi.ClientSecret})
}
