package main

import "fmt"

/**
Channel
一个无缓存channel的发送操作会导致发送者阻塞，直到接收者接收。
基于channel发送消息有两个重要方面：
每个消息都有一个值，但有时候通讯的事实和发生的时刻同样重要。当强调通讯发生的时刻时，将其称为消息事件。
*/

/**
串联的channels（pipeline）
可以通过close函数来关闭channel减少不必要的等待。
当channel关闭之后再向其发送数据会导致panic。
*/

func f() {
	naturuals := make(chan int)
	squares := make(chan int)
	go func() {
		for x := 0; ; x++ {
			naturuals <- x
		}
	}()

	go func() {
		for {
			x := <-naturuals
			squares <- x * x
		}
	}()

	for {
		fmt.Println(<-squares)
	}
}

/**
 没有办法直接测试一个channel是否被关闭，但是接收操作有一个变体形式：
 它多接收一个结果，多接收的第二个结果是一个布尔值ok，
 ture表示成功从channels接收到值，false表示channels已经被关闭并且里面没有值可接收。
Go语言的range循环可直接在channels上面迭代。
 使用range循环是上面处理模式的简洁语法，它依次从channel接收数据，
 当channel被关闭并且没有值可接收时跳出循环。

其实你并不需要关闭每一个channel。只有当需要告诉接收者goroutine，
 所有的数据已经全部发送时才需要关闭channel。不管一个channel是否被关闭，当它没有被引用时将会被Go语言的垃圾自动回收器回收。
*/

func main1() {
	naturuals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x < 100; x++ {
			naturuals <- x
		}
		close(naturuals)
	}()

	go func() {
		for x := range naturuals {
			squares <- x * x
		}
		close(squares) //如果没有会报错
		/*
			x:=<-naturuals
			squares<-x*x
		*/
	}()

	for x := range squares {
		fmt.Println(x)
	}
}

/**
任何双向channel向单向channel变量的赋值操作都将导致该隐式转换,并没有反向转换的语法
*/

func counter(out chan<- int) {
	for x := 0; x < 100; x++ {
		out <- x
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for v := range in {
		out <- v * v
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}

func main() {
	naturals := make(chan int)
	squares := make(chan int)
	go counter(naturals)
	go squarer(squares, naturals)
	printer(squares)
}

/**
带缓存的channel
如果使用无缓存的channel，那么两个慢的goroutine将会因为没有人接收而被永远卡住。
这种情况称为goroutine泄漏，和垃圾变量不同，泄漏的goroutine并不会被自动回收，是一个BUG。

关于无缓存或带缓存channels之间的选择，或者是带缓存channels的容量大小的选择，都可能影响程序的正确性。
无缓存channel更强地保证了每个发送操作与相应的同步接收操作；但是对于带缓存channel，这些操作是解耦的。
同样，即使我们知道将要发送到一个channel的信息的数量上限，
创建一个对应容量大小的带缓存channel也是不现实的，
因为这要求在执行任何接收操作之前缓存所有已经发送的值。如果未能分配足够的缓存将导致程序死锁。
*/

func mirroredQuery() string {
	response := make(chan string, 3)
	go func() { response <- request("hello") }()
	go func() { response <- request("world") }()
	go func() { response <- request("nihao") }()
	return <-response //返回最快的
}

func request(hostname string) (response string) {
	return "echo " + hostname
}
