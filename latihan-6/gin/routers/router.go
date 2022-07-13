package routers

import (
	"latihan-6-gin/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.GET("/cars", controllers.GetListCar)
	router.POST("/cars", controllers.CreateCar)
	router.GET("/cars/:carId", controllers.GetCar)
	router.PUT("/cars/:carId", controllers.UpdateCar)
	router.DELETE("/cars/:carId", controllers.DeleteCar)

	return router
}
