package main

import (
	"fmt"
	"log"
	"sync"
)

/**
并发的循环
没有什么直接的办法等待goroutine退出，但可以使用一个共享的channel发送事件。

注意我们将f的值作为一个显式的变量传给了函数，而不是在循环的闭包中声明：
这将导致f被更新
for _, f := range filenames {
	go func() {
		thumbnail.ImageFile(f) // NOTE: incorrect!
		// ...
	}()
}
*/

func makeThumbnails(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			fmt.Println("thumbnail file:", f)
			ch <- struct{}{}
		}(f)
	}

	for range filenames {
		<-ch
	}
}

/**
返回错位：
这个程序有一个微妙的bug。当它遇到第一个非nil的error时会直接将error返回到调用方，
使得没有一个goroutine去排空errors channel。
这样剩下的worker goroutine在向这个channel中发送值时，
都会永远地阻塞下去，并且永远都不会退出。
这种情况叫做goroutine泄露，可能会导致整个程序卡住或者跑出out of memory的错误。
最简单的解决办法就是用一个具有合适大小的buffered channel，这样这些worker goroutine向channel中发送错误时就不会被阻塞。
*/
func makeThumbnails2(filenames []string) error {
	ch := make(chan error)
	for _, f := range filenames {
		go func(f string) {
			_, err := fmt.Printf("thumbnail file:%v", f)
			ch <- err
		}(f)
	}

	for range filenames {
		if err := <-ch; err != nil {
			return err //导致goroutine泄漏
		}
	}
	return nil
}

/**
为了知道最后一个goroutine什么时候结束（最后一个结束并不一定是最后一个开始），
我们需要一个递增的计数器，在每一个goroutine启动时加一，在goroutine退出时减一。
这需要一种特殊的计数器，这个计数器需要在多个goroutine操作时做到安全并且提供在其减为零之前一直等待的一种方法。
这种计数类型被称为sync.WaitGroup

注意Add和Done方法的不对称。
Add是为计数器加一，必须在worker goroutine开始之前调用，而不是在goroutine中；
否则的话我们没办法确定Add是在"closer" goroutine调用Wait之前被调用。
并且Add还有一个参数，但Done却没有任何参数；其实它和Add(-1)是等价的。
当我们使用并发循环，但又不知道迭代次数时很通常而且很地道的写法。
*/

func makeThumbnails3(filenames []string) int {
	sizes := make(chan int)
	var wg sync.WaitGroup
	for _, f := range filenames {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			n, err := fmt.Printf("thumbnail file:%v\n", f)
			if err != nil {
				log.Println(err)
				return
			}
			sizes <- n
		}(f)
	}

	go func() {
		wg.Wait()
		close(sizes)
	}()
	var total int
	for size := range sizes {
		total += size
	}
	return total
}

func main() {
	fmt.Println(makeThumbnails3([]string{"a", "hello", "hhh"}))
}
