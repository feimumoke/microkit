package main

import (
	"bufio"
	"fmt"
	"os"
)

/**
map
delete可以删除元素,删除不存在的元素也是安全的
如果一个查找失败将返回value类型对应的零值
map中的元素并不是一个变量，不能对她进行取址
map遍历的顺序是不一定的，常见的顺序遍历是先将key排序之后再根据key取value
map类型的零值是nil

map和slice一样，不能进行相等比较，唯一的例外是和nil进行比较。
要判断两个map相等必须通过循环实现
*/

func basemap() {
	ages := map[string]int{"zhangsan": 20}
	ages["lisi"] = 30
	delete(ages, "zhangsan")
	delete(ages, "haha")
	ages["wangwu"] = ages["wangwu"] + 1
	ages["zhaoliu"]++
	fmt.Println("wangwu:", ages["wangwu"])
	fmt.Println("zhaoliu:", ages["zhaoliu"])
	fmt.Println("zhangqi:", ages["zhangqi"]) //默认取零值

	age, ok := ages["wangba"]
	if !ok {
		//不存在wangba
		ages["wangba"] = 40
	} else {
		fmt.Println("wangba:", age)
	}

	//nil 值的map不能进行插入
	var zero map[string]string
	fmt.Println("zero for map:", zero)
	fmt.Println(zero == nil)
}

/**
不能简单的用xv!=y[k]判断，这样会导致零值的map判断错误，如:
equal(map[string]int{"A":0}, map[string]int{"B":32})
*/
func equal(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k, xv := range x {
		if yv, ok := y[k]; !ok || yv != xv {
			return false
		}
	}
	return true
}

/**
go语言没有set类型，但是map中的key也是不相同的，可以用map实现set的功能
*/

func dedup() {
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
			fmt.Println(line)
		}
	}
	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}

/**
有时候需要map的key或者value是slice类型，但是map的key必须是可比较类型，可以通过两个步骤绕过：
1、定义一个辅助函数k，将slice转化为map对应的string类型的key，确保只有x和y相等时k(x)==k(y)才成立
2、创建一个key为string类型的map，对map操作时先用k函数将slice转换为string
eg:记录提交相同字符串列表的次数
*/

var m = make(map[string]int)

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}

func Add(list []string) {
	m[k(list)]++
}
func Count(list []string) int {
	return m[k(list)]
}

/**
map的value可以说复合类型
*/
var graph = make(map[string]map[string]bool)

func addEdge(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}
func hasEdge(from, to string) bool {
	return graph[from][to]
}
func main() {
	basemap()
	dedup()
}
