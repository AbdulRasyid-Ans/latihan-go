package main

import (
	"fmt"
	"latihan-7-gorm/db"
	"latihan-7-gorm/models"
	"log"
	"strings"

	"gorm.io/gorm"
)

func main() {
	db := db.ConnectDB()
	if db != nil {
		log.Println("db connected")
	}

	createUser("admin@abdulrasyid.my.id", db)
	// getListUsers(db)
	// createProduct(db)
	// getUserWithProducts(db, 1)
	// updateProduct(db)
	// deleteProduct(db, 3)
}

func createUser(email string, db *gorm.DB) {
	user := models.User{
		Email: email,
	}

	err := db.Create(&user).Error
	if err != nil {
		panic(err.Error())
	}

	log.Println("created", email, "success")

}

func getListUsers(db *gorm.DB) {
	var users []models.User

	err := db.Find(&users).Error
	if err != nil {
		panic(err.Error())
	}

	for _, user := range users {
		log.Println("email :", user.Email)
		log.Println(strings.Repeat("-", 20))
	}
}

func getUserWithProducts(db *gorm.DB, id uint) {
	var user models.User

	err := db.Preload("Products").First(&user, "id=?", id).Error
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%+v\n", user)
}

func createProduct(db *gorm.DB) {
	product := models.Product{
		Name:   "Laptop",
		Brand:  "Asus",
		UserID: 3,
	}

	err := db.Create(&product).Error
	if err != nil {
		panic(err.Error())
	}

	log.Println("create product succcess")
}

func updateProduct(db *gorm.DB) {
	var product models.Product

	err := db.Model(&product).Where("id = ?", 2).Updates(models.Product{UserID: 1, Name: "Update"}).Error
	if err != nil {
		panic(err.Error())
	}

	log.Printf("Update product: %+v\n", product)
}

func deleteProduct(db *gorm.DB, id uint) {
	var product models.Product

	err := db.Where("id = ?", id).Delete(&product).Error
	if err != nil {
		panic(err.Error())
	}

	log.Printf("Delete product ID: %d\n", id)
}
