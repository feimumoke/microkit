package main

import (
	"fmt"
	"golang.org/x/net/html"
)

/**
Recover
如果在defered函数中调用了内置函数recover，并且定义该defer语句的函数
发生了panic异常，recover会使程序从panic中恢复，且返回painc value。
导致panic异常的函数不会继续运行，但能正常返回。未发生painc时调用recover，
recover会返回nil。

不加区分的恢复panic异常是不可取的，因为panic之后无法保证包级变量的状态
仍然和预期一致。如数据结构的一次更新没有完整完成，文件或者网络链接没有关闭，
获得的锁没有释放等。

不应该恢复其他包引起的panic。公有API应该将函数运行失败作为error返回，而不是panic。
同样，也不应该恢复他人开发函数引起的panic。
例外就是net包中handler引起的panic异常web服务器会调用recover。
*/

func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}
	defer func() {
		switch p := recover(); p {
		case nil:
		case bailout{}:
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p)
		}
	}()

	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" && n.FirstChild != nil {
			if title != "" {
				panic(bailout{})
			}
			title = n.FirstChild.Data
		}
	}, nil)
	if title == "" {
		return "", fmt.Errorf("no title element")
	}
	return title, nil
}
