package main

import (
	"fmt"
	"os"
	"time"
)

/**
基于select的多路复用
select的每一个case代表一个通信操作，在某个channel上发送或者接收。
select会等待case中有能够执行的case时去执行。
当条件满足时，select才会去通信并执行case之后的语句；这时候其它通信是不会执行的。一个没有任何case的select语句写作select{}，会永远地等待下去。
*/

func f1() {
	fmt.Println("launch rocket, Press return to about in 10 seconds")
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	select {
	case <-time.After(10 * time.Second):
	//Do nothing
	case <-abort:
		fmt.Println("about launch")
		return
	}
	fmt.Println("lanuch !!!")
}

func f2() {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case x := <-ch:
			fmt.Println(x)
		case ch <- i:
		}
	}
}

/**
如果多个case同时就绪时，select会随机地选择一个执行，
这样来保证每一个channel都有平等的被select的机会。

Tick函数挺方便，但是只有当程序整个生命周期都需要这个时间时我们使用它才比较合适。
否则的话，我们应该使用下面的这种模式：

ticker := time.NewTicker(1 * time.Second)
<-ticker.C    // receive from the ticker's channel
ticker.Stop() // cause the ticker's goroutine to terminate
*/

func f3() {
	fmt.Println("launch rocket, Press return to about in 10 seconds")
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()
	tick := time.Tick(1 * time.Second)
	for count := 10; count > 0; count-- {
		fmt.Println(count)
		select {
		case <-tick:
		//Do nothing
		case <-abort:
			fmt.Println("about launch")
			return
		}
	}
	fmt.Println("lanuch !!!")
}

func main() {
	f3()
}
