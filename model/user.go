package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name     string `sql:"size:40"`
	Password string `sql:"size:60"`
	Token    string `sql:"size:100"`
}
