package ravel

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/martinlindhe/ravel/migrations"
)

func RunMigrations(db *gorm.DB) error {

	fmt.Println("Running migrations ...")

	migrations.Up0001(db)
	migrations.Up0002(db)

	fmt.Println("Migrations successful!")
	return nil
}

func RollbackMigrations(db *gorm.DB) {
	migrations.Down0002(db)
	migrations.Down0001(db)
}
