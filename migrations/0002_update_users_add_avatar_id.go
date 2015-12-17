package migrations

import "github.com/jinzhu/gorm"

type user2 struct {
	gorm.Model
	Name     string `sql:"size:40"`
	Password string `sql:"size:60"`
	Token    string `sql:"size:100"`
	AvatarID uint
}

func Up0002(db *gorm.DB) {

	// XXX how to add column with gorm?
	//db.Table(&user2{}).Add
}

func Down0002(db *gorm.DB) {

	db.Model(&user2{}).DropColumn("AvatarID")
}
