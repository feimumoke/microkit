package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

/**
channel 的广播来控制并发的退出

一种可能的手段是向abort的channel里发送和goroutine数目一样多的事件来退出它们。
如果这些goroutine中已经有一些自己退出了，那么会导致我们的channel里的事件数比goroutine还多，
这样导致我们的发送直接被阻塞。另一方面，如果这些goroutine又生成了其它的goroutine，
我们的channel里的数目又太少了，所以有些goroutine可能会无法接收到退出消息。
一般情况下我们是很难知道在某一个时刻具体有多少个goroutine在运行着的。
另外，当一个goroutine从abort channel中接收到一个值的时候，他会消费掉这个值，
这样其它的goroutine就没法看到这条信息。
为了能够达到我们退出goroutine的目的，我们需要更靠谱的策略，
来通过一个channel把消息广播出去，这样goroutine们能够看到这条事件消息，
并且在事件完成之后，可以知道这件事已经发生过了。

回忆一下我们关闭了一个channel并且被消费掉了所有已发送的值，
操作channel之后的代码可以立即被执行，并且会产生零值。
我们可以将这个机制扩展一下，来作为我们的广播机制：
不要向channel发送值，而是用关闭一个channel来进行广播。
*/
var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		fmt.Println("done")
		return true
	default:
		fmt.Println("false")
		return false
	}
}
func f05(wg sync.WaitGroup) {
	defer wg.Done()
	if cancelled() {
		return
	}
	os.Stdin.Read(make([]byte, 1))
	close(done)
	fmt.Println("close")
	time.Sleep(10 * time.Second)
}

func main() {
	end := make(chan bool)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		f05(wg)
		end <- true
	}()
	go func() {
		if cancelled() {
			end <- true
		}

	}()
	go wg.Wait()
	<-end
}
