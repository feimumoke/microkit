package main

import (
	"fmt"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"os"
	"sort"
)

/**
匿名函数
拥有函数名的函数只能在包级语法块中被声明，通过函数字面量可以绕过这一限制。
函数字面量的语法和函数声明类似，去被在于func关键字后面没有函数名。
函数值字面量是一种表达式，它的值被称为匿名函数。它运行我们在使用时再定义。
通过这种方式定义的函数可以访问完整的词法环境，即函数中定义的内部函数可以引用该函数的变量。

函数值不仅是一段代码，还记录了状态，它属于引用类型。匿名函数可以和函数中存在变量引用。
*/

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

/**
变量的声明周期不由它的作用域决定，squares函数返回之后，变量x仍然隐式的存在在f中。
*/
func main() {
	f := squares
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	links, _ := Extract("http://www.baidu.com")
	fmt.Println(links)
	breadthFirst(crawl, []string{"http://www.baidu.com"})
}

/**
当匿名函数需要被递归调用时，必须首先声明一个变量，再将匿名函数赋值给这个变量
如果不分成两步，函数字面量无法与visitAll绑定，无法递归调用
*/

func toposort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}
	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

func Extract(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := resp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	eachNode(doc, visitNode, nil)
	return links, nil
}
func eachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		eachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

/**
网页抓取的核心在于如何遍历图，toposort中展示了如何用深度优先遍历图，
下面介绍深度广度有限
*/

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

func crawl(url string) []string {
	fmt.Println(url)
	list, err := Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

/**
捕获迭代变量
*/

//创建一些目录，再将目录删除。
/**
for循环引入新的词法块，循环变量dir在这个词法块中被声明，该循环中所有函数值都共享相同的循环变量。
注意，函数值中记录的是循环变量的内存地址，而不是循环变量某一时刻的值。后续的迭代会不断更新d的值，
当删除操作进行时，for循环已经完成，d保存的是最后一次迭代的值。
通常为了解决这个问题会引入一个与循环变量同名的局部变量，作为循环变量的副本。
对于循环变量i也有这个问题

*/
func operDir() {
	var rmdirs []func()
	var tempDir = []string{"/tmp/a", "/tmp/b"}
	for _, dir := range tempDir {
		dir := dir //不可忽略
		os.MkdirAll(dir, 0755)
		rmdirs = append(rmdirs, func() {
			os.RemoveAll(dir)
		})
	}

	// do some work
	for _, rmdir := range rmdirs {
		rmdir()
	}
}
