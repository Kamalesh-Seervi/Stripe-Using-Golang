package main

import (
	"github.com/Kamalesh-Seervi/stripe-in-go/models"
	"github.com/Kamalesh-Seervi/stripe-in-go/routes"
)

func main() {
	models.Db_Setup()
	routes.Server()
}
