package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

/**
文本和HTML模板
一个模板是一个字符串或者一个文件，里面包含双花括号对象{{action}}
在一个action中，都有一个当前值的概念，对应点操作符。
{{range .Item}}和{{end}}对应一个循环action。
一个action中|操作符表示将前一个表达式的结果当中后一个函数的输入
*/

const templ = `{{.TotalCount}} issues:
{{range .Items}}--------------------------
Number: {{.Number}}
User:   {{.User.Login}}
Title:  {{.Title | printf "$.64s"}}
Age:    {{.CreatedAt | daysAgo}} days
{{end}}`

func daysAgo(t time.Time) int {
	return int(time.Since(t).Hours() / 24)
}

func output(data interface{}) {
	report := template.Must(template.New("report").
		Funcs(template.FuncMap{"daysAgo": daysAgo}).
		Parse(templ))
	if err := report.Execute(os.Stdout, data); err != nil {
		log.Fatal(err)
	}
}

/**
html模板使用和text模板相同的API和模板语言，但是增加了一个将字符串自动转义特性，可以避免HTML注入等问题
*/

func main() {

}
