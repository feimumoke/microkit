package main

import "fmt"

/**
slice
*/

func innerappend() {
	var runes []rune
	for _, r := range "Hello, World" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)
}

/**
每一次容量变化都会导致内存重新分配，因此每次扩容直接翻倍
*/
func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		z = x[:zlen]
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	return z
}

/**
输入和输出的slice是同一个底层数组
*/
func nonempty(strings []string) []string {
	i := 0
	for _, s := range strings {
		if s != "" {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

func nonempty2(strings []string) []string {
	out := strings[:0]
	for _, s := range strings {
		if s != "" {
			out = append(out, s)
		}
	}
	return out
}

func stack() {
	var stack []string
	stack = append(stack, "a")
	top := stack[len(stack)-1]
	fmt.Println("top:", top)
	stack = stack[:len(stack)-1]
}

/**
slice代表变长的序列，序列中每个元素的类型相同。一个slice是一个轻量级的数据结构，提供了访问数组子序列的功能。
slice底层引用着一个数组对象，一个slice由三部分组成：指针、长度、容量。指针指向第一个slice元素对应的底层数组
元素的地址，slice的第一个元素不一定是数组的第一个元素。长度对应slice中元素的数目，长度不能超过容量，容量一般是slice
开始的位置到数组的结尾位置。
arr[:]切片操作的是整个数组
*/

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
func main() {
	s := [...]int{0, 1, 2, 3, 4, 5}
	reverse(s[:])
	fmt.Println(s)
	stack()
	innerappend()
	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}
