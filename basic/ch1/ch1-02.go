package main

import (
	"fmt"
	"os"
	"strings"
)

/**
命令行参数
os包以跨平台的方式提供了一些与操作系统交互的函数和变量
程序的命令行参数可以通过os包的Args变量获取，os包外部采用os.Args访问该变量
os.Args是一个字符串的切片，os.Args[0]是命令本身的名字
*/

func unix_echo() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func unix_echo2() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func unix_echo3() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}

func main() {
	unix_echo()
	unix_echo2()
	unix_echo3()
}
