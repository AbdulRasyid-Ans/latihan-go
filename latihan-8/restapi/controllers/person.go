package controllers

import (
	"net/http"
	"restapi/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type PersonController struct {
	db *gorm.DB
}

func NewPersonController(db *gorm.DB) *PersonController {
	return &PersonController{
		db: db,
	}
}

func (p *PersonController) CreatePerson(ctx *gin.Context) {
	var person models.Person

	err := ctx.ShouldBindJSON(&person)
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	err = p.db.Create(&person).Error
	if err != nil {
		badRequestResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusCreated, gin.H{
		"success": true,
		"message": "created success",
	})
}

func (p *PersonController) GetAllPeople(ctx *gin.Context) {
	limit := ctx.Query("limit")
	limitInt := 10

	if limit != "" {
		l, err := strconv.Atoi(limit)
		if err == nil {
			limitInt = l
		}
	}

	var people []models.Person
	var total int

	err := p.db.Limit(limitInt).Find(&people).Count(&total).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			notFoundResponse(ctx, err.Error())
			return
		}
		badRequestResponse(ctx, err.Error())
		return
	}

	writeJsonResponse(ctx, http.StatusOK, gin.H{
		"success": true,
		"payload": people,
		"query": map[string]interface{}{
			"limit": limitInt,
			"total": total,
		},
	})

}
