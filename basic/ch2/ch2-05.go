package main

import (
	"encoding/json"
	"fmt"
	"log"
)

/**
JSON
一个json对象是一个字符串到值的映射，json对象类型可以用于编码Go语言的map类型
一个结构体成员Tag是在编译阶段关联到该成员的元信息字符串。
结构体成员Tag可以是任意的字符串面值，但通常是一系列用空格分割的key:"value"键值对。
omitempty选项表示该成员为空或者零值（此处为false）时该JSON对象不包含此成员。
*/
type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Avata", Year: 2012, Color: true, Actors: []string{"Salam", "Alice"}},
	{Title: "Wukongzhuan", Year: 2018, Color: true, Actors: []string{"Pengyuyan", "Yuwenle"}},
	{Title: "ShaPoLang", Year: 2006, Color: false, Actors: []string{"zhenzidan", "Hongjinbao"}},
}

/**
slice转为Json的过程称为编组(marshaling)
Marshal函数返回一个编码后的字节slice，且没有缩进
*/

func marshaling() {
	data, err := json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
	var titles []struct {
		Title string
		Year  int `json:"released"`
	}
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshling failed: %s", err)
	}
	fmt.Println(titles)
}

//整齐输出

func marshalIndent() {
	data, err := json.MarshalIndent(movies, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshaling failed: %s", err)
	}
	fmt.Printf("%s\n", data)
}

func main() {
	marshaling()
	marshalIndent()
}
