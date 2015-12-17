package migrations

import "github.com/jinzhu/gorm"

type user1 struct {
	gorm.Model
	Name     string `sql:"size:40"`
	Password string `sql:"size:60"`
	Token    string `sql:"size:100"`
}

func Up0001(db *gorm.DB) {

	db.CreateTable(&user1{})
}

func Down0001(db *gorm.DB) {

	db.DropTable(&user1{})
}
