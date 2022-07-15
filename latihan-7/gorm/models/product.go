package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"varchar(100)"`
	Brand     string `gorm:"varchar(100)"`
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *Product) BeforeCreate(tx *gorm.DB) (err error) {
	var user User
	fmt.Println(">> Before create Product")
	if len(p.Name) <= 5 {
		err = fmt.Errorf("length of name must be greater than 5")
	}

	errUser := tx.First(&user, "id=?", p.UserID).Error
	if errUser != nil {
		err = fmt.Errorf("User ID %d not found", p.UserID)
	}

	return
}

func (p *Product) AfterCreate(tx *gorm.DB) (err error) {
	fmt.Println(">> After create Product")
	return
}

func (p *Product) BeforeUpdate(tx *gorm.DB) (err error) {
	fmt.Println(">> Before Update Product")
	return
}

func (p *Product) AfterUpdate(tx *gorm.DB) (err error) {
	fmt.Println(">> After Update Product")
	return
}

func (p *Product) BeforeDelete(tx *gorm.DB) (err error) {
	fmt.Println(">> Before Delete Product")
	return
}

func (p *Product) AfterDelete(tx *gorm.DB) (err error) {
	fmt.Println(">> After Delete Product")
	return
}
