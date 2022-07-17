package main

import (
	"restapi/controllers"
	"restapi/database"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db := database.ConnectDB()

	personController := controllers.NewPersonController(db)

	router.POST("/person", personController.CreatePerson)
	router.GET("/person", personController.GetAllPeople)

	PORT := ":4444"
	router.Run(PORT)
}
