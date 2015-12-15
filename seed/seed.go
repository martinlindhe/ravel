package seed

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/martinlindhe/ubique.se/model"
	"github.com/martinlindhe/ubique.se/password"
)

func Seed(db *gorm.DB) {

	fmt.Printf("Running seeder ...\n")

	user := model.User{Name: "martin", Password: password.Make("password")}

	err := db.Create(&user)
	if err != nil {
		fmt.Println(err)
	}

}
