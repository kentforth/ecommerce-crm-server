package main

import (
	"ecommerce-crm-server/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	gin.SetMode(gin.ReleaseMode)

	app := gin.New()

	router := app.Group("/")
	routes.AddRoutes(router)

	//utils.ConnectDatabase()

	app.Run(":" + os.Getenv("PORT"))

}
