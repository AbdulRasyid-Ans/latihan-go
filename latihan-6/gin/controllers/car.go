package controllers

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type Car struct {
	CarId string `json:"car_id"`
	Brand string `json:"brand"`
	Model string `json:"model"`
	Price int    `json:"price"`
}

var Cars []Car

func CreateCar(ctx *gin.Context) {
	var newCar Car
	var newCarId string

	err := ctx.ShouldBindJSON(&newCar)
	if err != nil {
		writeJsonResponse(ctx, http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	if len(Cars) > 0 {
		lastCarId := Cars[len(Cars)-1].CarId
		lastIdNumber, _ := strconv.Atoi(strings.Split(lastCarId, "-")[1])
		newCarId = fmt.Sprintf("car-%d", lastIdNumber+1)
	} else {
		newCarId = fmt.Sprintf("car-%d", len(Cars)+1)
	}

	newCar.CarId = newCarId
	Cars = append(Cars, newCar)

	writeJsonResponse(ctx, http.StatusCreated, gin.H{
		"status": "success",
		"car":    newCar,
	})
}

func GetListCar(ctx *gin.Context) {
	if len(Cars) == 0 {
		writeJsonResponse(ctx, http.StatusNotFound, gin.H{
			"error": "no cars data",
		})
		return
	}
	writeJsonResponse(ctx, http.StatusOK, gin.H{
		"status": "success",
		"data":   Cars,
	})
}

func GetCar(ctx *gin.Context) {
	carId := ctx.Param("carId")
	found := false
	var carData Car

	for i, car := range Cars {
		if carId == car.CarId {
			found = true
			carData = Cars[i]
			break
		}
	}

	if !found {
		writeJsonResponse(ctx, http.StatusNotFound, gin.H{
			"error": "car not found",
		})
		return
	}
	writeJsonResponse(ctx, http.StatusOK, gin.H{
		"status": "success",
		"data":   carData,
	})
}

func UpdateCar(ctx *gin.Context) {
	carId := ctx.Param("carId")
	found := false
	var updatedCar Car

	err := ctx.ShouldBindJSON(&updatedCar)
	if err != nil {
		writeJsonResponse(ctx, http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	for i, car := range Cars {
		if carId == car.CarId {
			updatedValue := reflect.ValueOf(&updatedCar).Elem()
			currentValue := reflect.ValueOf(&Cars[i]).Elem()

			fmt.Println(updatedCar)
			found = true
			for i := 0; i < updatedValue.NumField(); i++ {
				isFieldEmpty := updatedValue.Field(i).IsZero()
				if isFieldEmpty {
					updatedValue.Field(i).Set(currentValue.Field(i))
				}
			}
			Cars[i] = updatedCar
			break
		}
	}

	if !found {
		writeJsonResponse(ctx, http.StatusNotFound, gin.H{
			"error": "car not found",
		})
		return
	}
	writeJsonResponse(ctx, http.StatusOK, gin.H{
		"status": "success",
		"data":   updatedCar,
	})
}

func DeleteCar(ctx *gin.Context) {
	carId := ctx.Param("carId")
	found := false
	carIdx := 0

	for i, car := range Cars {
		if carId == car.CarId {
			found = true
			carIdx = i
			break
		}
	}

	if !found {
		writeJsonResponse(ctx, http.StatusNotFound, gin.H{
			"error": "car not found",
		})
		return
	}

	var newCars []Car
	newCars = append(newCars, Cars[:carIdx]...)
	newCars = append(newCars, Cars[carIdx+1:]...)

	Cars = newCars

	writeJsonResponse(ctx, http.StatusOK, gin.H{
		"status": "deleted",
		"data":   newCars,
	})
}

func writeJsonResponse(ctx *gin.Context, status int, payload interface{}) {
	ctx.JSON(status, payload)
}
