package dbmodel

import "time"

type Users struct {
	UserId   int       `gorm:"column:user_id;AUTO_INCREMENT;PRIMARY_KEY"`
	UserName string    `gorm:"column:user_name;type:varchar(200);unique_index"`
	UserPwd  string    `gorm:"column:user_pwd;type:varchar(50)"`
	UserDate time.Time `gorm:"column:user_date"`
}
