package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Email     string `gorm:"not null;unique;varchar(100)"`
	Products  []Product
	CreatedAt time.Time
	UpdatedAt time.Time
}
