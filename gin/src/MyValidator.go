package src

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"regexp"
)

func TopicUrl(fl validator.FieldLevel) bool {
	fmt.Printf("fl:%+v", fl)
	_, oks := fl.Top().Interface().(*TopicBox)
	_, ok := fl.Top().Interface().(*Topics)
	if oks || ok {
		s := fl.Field().String()
		fmt.Println("Field:", s)
		if m, _ := regexp.MatchString("^\\w{4,10}$", s); m {
			return true
		}
	}
	return false
}

func TopicsValidate(fl validator.FieldLevel) bool {
	fmt.Printf("fl:%+v", fl)
	topics, ok := fl.Top().Interface().(*TopicBox)
	if ok && topics.TopicListSize == len(topics.TopicList) {
		return true
	}
	return false
}
