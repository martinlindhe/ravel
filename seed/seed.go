package seed

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/martinlindhe/ravel/model"
	"github.com/martinlindhe/ravel/password"
)

func Seed(db *gorm.DB) {

	fmt.Printf("Running seeder ...\n")

	user := model.User{Name: "martin", Password: password.Make("password")}

	err := db.Create(&user)
	if err != nil {
		fmt.Println(err)
	}

}
