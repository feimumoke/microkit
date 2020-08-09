package main

import (
	"fmt"
	"os"
	"runtime"
)

/**
Painc
数组访问越界、空指针在运行时会引起panic异常
一般而言，panic会中断程序运行并执行goroutine的延时函数（defer）
然后程序崩溃并输出日志信息包括panic value和函数调用堆栈跟踪信息。
不是所有panic都来自运行时，直接调用内置的panic函数也会引起panic异常。

panic会引起程序的崩溃，因此一般用于严重错误
runtime包可以输出堆栈信息
*/

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x)
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}
func main() {
	printStack()
	f(3)
}
