package src

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

var DBHelper *gorm.DB
var err error

func InitDB() {
	DBHelper, err = gorm.Open("mysql", "root:tiger@/gin?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		//fmt.Println("init db err:",err)
		//log.Fatal("DB init error",err)
		ShutDownServer(err)
		return
	}
	DBHelper.LogMode(true)
	DBHelper.DB().SetMaxIdleConns(10)
	DBHelper.DB().SetMaxOpenConns(100)
	DBHelper.DB().SetConnMaxLifetime(time.Hour)
}
