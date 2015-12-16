package ravel

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/martinlindhe/ravel/model"
)

func Migrate(db *gorm.DB) {

	fmt.Println("Running migrations ...")

	db.Set("gorm:table_options", "ENGINE=InnoDB")

	// XXX go thru a list of migrations and run them, somehow.. like in laravel

	// XXX auto-register all found models

	user_create_Up(db)

	// XXX can we introspect and find all structs in the model package, as strings?
	//db.AutoMigrate(&model.User{})

	// Feel free to change your struct, AutoMigrate will keep your database up-to-date.
	// AutoMigrate will ONLY add *new columns* and *new indexes*,
	// WON'T update current column's type or delete unused columns, to protect your data.
	// If the table is not existing, AutoMigrate will create the table automatically.

	fmt.Println("Migrations successful!")
}

func Rollback(db *gorm.DB) {
	// XXX function that ROLLS BACK all migrations

	fmt.Printf("Rolling back migrations... \n")
	user_create_Down(db)
}

func user_create_Up(db *gorm.DB) {
	db.CreateTable(&model.User{})
}

func user_create_Down(db *gorm.DB) {
	db.DropTable(&model.User{})
}
