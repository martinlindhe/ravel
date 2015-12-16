package ravel

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/martinlindhe/ravel/model"
)

// Seed runs all seeds for the app
func Seed(db *gorm.DB) {

	fmt.Printf("Running seeders ...\n")

	seedUsers(db)
}

func seedUsers(db *gorm.DB) {
	user := model.User{Name: "martin", Password: HashPassword("password")}

	db.Create(&user)
}
