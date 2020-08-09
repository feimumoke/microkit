package main

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
)

/**
函数值
函数被看作第一类值，函数跟其他值一样，拥有类型，零值是nil。
调用值为nil的函数值会引起painc
函数值之间是不能比较的，因此不能用函数值作为map的key
*/

func squire(n int) int {
	return n * n
}

// 函数值使得不仅可以通过数据来参数化函数，也可以通过行为
func add1(r rune) rune {
	return r + 1
}

// 使用函数值，可以将遍历节点的逻辑和操作节点的逻辑分离
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s>\n", depth*2, "", n.Data)
		depth++
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}
}

func main() {
	f := squire
	fmt.Println(f(3))
	fmt.Println(strings.Map(add1, "HAL-9000"))
	forEachNode(nil, startElement, endElement)
	var g func(int) int
	g(3)
}
