package appinit

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql",
		"root:tiger@tcp(localhost:3306)/gokit?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(50)
}

func GetDB() *gorm.DB {
	return db
}
