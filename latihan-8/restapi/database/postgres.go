package database

import (
	"fmt"
	"restapi/models"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const (
	POSTGRES_HOST   = "localhost"
	POSTGRES_PORT   = "5432"
	POSTGRES_USER   = "postgre_username"
	POSTGRES_PASS   = "postgre_password"
	POSTGRES_DBNAME = "project_database"
)

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		POSTGRES_HOST,
		POSTGRES_PORT,
		POSTGRES_USER,
		POSTGRES_PASS,
		POSTGRES_DBNAME,
	)

	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}

	db.Debug().AutoMigrate(models.Person{})

	return db
}
