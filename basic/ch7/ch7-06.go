package main

import "fmt"

/**
goroutine 和线程
goroutine采用动态栈,一般只需要2KB，最大可伸缩到1GB，OS线程一般固定栈大小（2M),

OS线程会被内核调度，调度函数会挂起当前线程并将寄存器中的内容保存到内存中，
检查线程列表并从内存中恢复将要运行线程的寄存器信息，恢复该线程现场并运行。

Go运行时包含自己的调度，采用m:n调度，在n个操作系统线程上调度m个goroutine。
Go调度不需要硬件定时器产生中断，而是被Go语言“建筑“本身进行调度的。
不需要内核的上下文切换。

Go的调度器使用了一个叫做GOMAXPROCS的变量来决定会有多少个操作系统的线程同时执行Go的代码。
其默认的值是运行机器上的CPU的核心数，所以在一个有8个核心的机器上时，调度器一次会在8个OS线程上去调度GO代码。（GOMAXPROCS是前面说的m:n调度中的n）。

goroutine没有id号
*/

func main() {
	for {
		go fmt.Print(0)
		fmt.Print(1)

	}
}
