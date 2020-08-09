package main

import "fmt"

/**
数组
数组长度是数组类型的组成部分，因此数组长度需要在编译期间确定。因此数组用的很少。
*/
var a [3]int
var q [3]int = [3]int{1, 2, 3}

func main() {
	fmt.Println(a[0])
	for i, v := range a {
		fmt.Printf("%d %d\n", i, v)
	}
	fmt.Println(q[2])
	// 数组长度根据初始化值确定
	p := [...]int{1, 2, 3, 4}
	fmt.Println(p[3])

	//----------------
	//指定一个索引和对应值列表的方法
	type Currency int
	const (
		USD Currency = iota
		EUR
		GBP
		RMB
	)
	symbol := [...]string{USD: "$", EUR: "@", GBP: "&", RMB: "#"}
	fmt.Println(RMB, symbol[RMB])
}

/**
当调用一个函数时，函数的每个调用参数将会被赋值给函数内部的参数变量，所以函数参数变量接收的是一个
复制的副本，并不是原始的调用的变量。因为函数参数传递的机制导致传递大的数组类型是低效的，且修改的只发生在复制的数组上。
可以显式的传入一个指针，可以将修改反馈到调用者。
*/

func zero(ptr *[32]byte) {
	for i := range ptr {
		ptr[i] = 0
	}
}
