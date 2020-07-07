package main

import (
	"github.com/go-acme/lego/v3/log"
	"github.com/go-playground/validator/v10"
	"zhuhui.com/microkit/micro/applib"
)

type Users struct {
	Username string   `validate:"required,min=6,max=20" vmsg:"用户名必须大于6位"`
	Userpwd  string   `validate:"required,min=6,max=18" vmsg:"用户密码必须大于6位"`
	Testname string   `validate:"abc" vmsg:"rule not ok"`
	UserTags []string `validate:"required,min=1,max=5,unique,dive,utag" vmsg:"Tag not valid"`
}

func main() {
	tags := []string{"a", "gffggds", "ad", "xa", "as"}
	users := &Users{Username: "ngs8an", Userpwd: "129983", Testname: "dd111111ab", UserTags: tags}
	validate := validator.New()
	//_ = validate.RegisterValidation("abc", func(fl validator.FieldLevel) bool {
	//	fmt.Println(fl.Field().String())
	//	matched, _ := regexp.MatchString("[a-zA-Z]\\w{5,19}", fl.Field().String())
	//	return matched
	//}, false)
	applib.AddRegexTag("abc", "[a-zA-Z]\\w{5,19}", validate)
	applib.AddRegexTag("utag", "^[a-zA-Z]{1,4}$", validate)
	err := applib.ValidErrMsg(users, validate.Struct(users))
	if err != nil {
		log.Fatal(err)
		//if errs, ok := err.(validator.ValidationErrors); ok {
		//	for _, e := range errs {
		//		fmt.Println(e.Value())
		//		fmt.Println(e.Field())
		//		fmt.Println(e.Tag())
		//		applib.GetValidMsg(&users, e.Field())
		//		fmt.Println()
		//	}
		//}
	}
}
