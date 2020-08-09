package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
)

/**
函数
函数的类型被称为函数的签名，如果两个函数的形式参数列表和返回值列表中的变量一一对应，
那么这两个函数被认为有相同的类型或者签名。形参和返回值的变量名不影响函数签名。
函数的形参和有名返回值作为函数最外层的局部变量，被存储在相同的词法块中。
实参通过值的方式传递，因此函数的形参是实参的拷贝，对形参的修改不会影响实参。
但是如果实参包括引用类型，如指针、slice、map、function、channel等类型，实参可能会因为函数的间接引用被修改。
没有函数体的声明，这表示该函数不是以Go实现的，这样的声明定义了函数的签名
*/

func findlink1() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}
func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

/**
outline有入栈操作，没有出栈，当其调用自身时，被调用者接收的是stack的拷贝。
修改的是stack的拷贝，可能会修改slice底层的数组，但并不会修改调用方的stack，
函数返回时调用方的stack与其调用自身之前完全一致。

go语言使用可变栈，大小按需增加。
*/

func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data)
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}
func main() {
	findlink1()
}
