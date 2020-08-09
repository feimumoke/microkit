package main

import "fmt"

/**
内存同步
在现代计算机中可能会有一堆处理器，每一个都会有其本地缓存（local cache）。
为了效率，对内存的写入一般会在每一个处理器中缓冲，并在必要时一起flush到主存。
这种情况下这些数据可能会以与当初goroutine写入顺序不同的顺序被提交到主存。
像channel通信或者互斥量操作这样的原语会使处理器将其聚集的写入flush并commit，
这样goroutine在某个时间点上的执行结果才能被其它处理器上运行的goroutine得到。


x:0 y:0
y:0 x:0

尽管goroutine A中一定需要观察到x=1执行成功之后才会去读取y，
但它没法确保自己观察得到goroutine B中对y的写入，所以A还可能会打印出y的一个旧版的值。
如果两个goroutine在不同的CPU上执行，每一个核心有自己的缓存，
这样一个goroutine的写入对于其它goroutine的Print，
在主存同步之前就是不可见的了。
*/

func main() {
	for i := 0; i < 10000; i++ {
		var x, y int
		go func() {
			x = 1
			fmt.Print("Y:", y, " ")
		}()
		go func() {
			y = 1
			fmt.Print("X:", x, " ")
		}()
	}
}
