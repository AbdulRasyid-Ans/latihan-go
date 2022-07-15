package db

import (
	"fmt"
	"latihan-7-gorm/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	POSTGRES_HOST = "localhost"
	POSTGRES_PORT = "5432"
	POSTGRES_DB   = "nama_database"
	POSTGRES_USER = "pg_username"
	POSTGRES_PASS = "pg_password"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		POSTGRES_HOST, POSTGRES_PORT, POSTGRES_USER, POSTGRES_PASS, POSTGRES_DB,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	db.AutoMigrate(models.User{}, models.Product{})

	return db
}
