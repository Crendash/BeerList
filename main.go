package main

import (
	"Beer-BackendV1/api"
	"Beer-BackendV1/controller"
	"Beer-BackendV1/database"
	"Beer-BackendV1/model"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	loadEnv()
	loadDatabase()
	serverApplication()
}

// loadDatabase connects to the Database and creates Tables accordingly to our model structs.
func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Entry{})
}

func indexView(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Headers", "access-control-allow-origin, access-control-allow-headers")
	c.JSON(http.StatusOK, gin.H{"message": "BEER BACKEND APP"})
}

// loadEnv loads the .env file
func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

// serverApplication runs and serves our application
func serverApplication() {
	router := gin.Default()

	// publicRoutes define the routes which you can access without being logged in
	router.GET("/", indexView)

	// Set routes for API
	// Update to POST, UPDATE, DELETE etc
	router.GET("/users", api.Users)
	router.GET("/beverages", controller.GetAllEntries)
	router.POST("/users/beverages", controller.GetAllEntriesOfUser)
	router.POST("/users/create", api.CreateUser)
	router.POST("/users/update-Balance", api.UpdateUserBalance)
	router.POST("/users/add-Beverage", controller.AddEntry)
	router.DELETE("/users/delete", api.DeleteUser)

	// starts the server on Port 8080
	router.Run(":8080")
	fmt.Println("Server running on Port 8080")
}
