package ravel

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"

	"github.com/jinzhu/gorm"
	"github.com/smallfish/simpleyaml"
)

func Migrate(db *gorm.DB) error {

	fmt.Println("Running migrations ...")

	db.Set("gorm:table_options", "ENGINE=InnoDB")

	migrationsPath := "./db/migrations"

	fileList, err := ioutil.ReadDir(migrationsPath)
	if err != nil {
		return err
	}

	for _, fi := range fileList {
		ext := path.Ext(fi.Name())
		if ext != ".yml" {
			log.Printf("[migrate] Unrecognized extension %s, skipping %s\n", ext, fi.Name())
			continue
		}

		fullPath := path.Join(migrationsPath, fi.Name())

		err = runMigration(fullPath)
		if err != nil {
			panic(err)
		}
	}

	//	user_create_Up(db)

	fmt.Println("Migrations successful!")
	return nil
}

func runMigration(fileName string) error {

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	// read yml
	fmt.Println(data)
	yaml, err := simpleyaml.NewYaml(data)
	if err != nil {
		return err
	}

	upStep, err := yaml.Get("up").GetIndex(0).String()
	if err != nil {
		return err
	}

	fmt.Println(upStep)

	// XXXX read up on gorm migration features....

	return nil
}

func Rollback(db *gorm.DB) {
	// XXX function that ROLLS BACK all migrations

	fmt.Printf("Rolling back migrations... \n")
	//user_create_Down(db)
}

/*
func user_create_Up(db *gorm.DB) {
	db.CreateTable(&model.User{})
}

func user_create_Down(db *gorm.DB) {
	db.DropTable(&model.User{})
}
*/
