package models

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Dbase *gorm.DB

type Product struct {
	Id          uint    `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

func Db_Setup() {
	// Load configuration from config.json
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// Read database connection details
	dbConfig := viper.Sub("DB")
	if dbConfig == nil {
		panic("DB configuration not found in config.json")
	}

	dbHost := dbConfig.GetString("Host")
	dbPort := dbConfig.GetString("Port")
	dbUser := dbConfig.GetString("User")
	dbPassword := dbConfig.GetString("Password")
	dbName := dbConfig.GetString("Name")

	conn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", dbHost, dbPort, dbUser, dbPassword, dbName)
	fmt.Println(conn)

	var err error

	Dbase, err = gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = Dbase.AutoMigrate(&Product{})
	if err != nil {
		fmt.Println(err)
	}
}
