package src

import "time"

type Topics struct {
	TopicId         int    `json:"id" gorm:"PRIMARY_KEY"`
	TopicTitle      string `json:"title" binding:"min=4,max=20" `
	TopicShortTitle string `json:"stitle" binding:"required,nefield=TopicTitle"`
	UserIP          string `json:"ip" binding:"ipv4"`
	TopicScore      int    `json:"score" binding:"omitempty,gt=5"`
	TopicUrl string `json:"url" binding:"omitempty,topicurl"`
	TopicDate time.Time `json:"date" binding:"required"`
}

type TopicQuery struct {
	UserName string `json:"username" form:"username"`
	Page     int    `json:"page" form:"page" binding:"required"`
	PageSize int    `json:"pagesize" form:"pagesize"`
}

type TopicBox struct {
	TopicList     []Topics `json:"topics" binding:"gt=0,lt=3,topics,dive"`
	TopicListSize int     `json:"size"`
}

type TopicClass struct {
	ClassId int `gorm:"PRIMARY_KEY"`
	ClassName string
	ClassRemark string `gorm:"Column:class_remark"`
}

func CreateTopic(id int, title string) Topics {
	return Topics{TopicId: id, TopicTitle: title}
}
