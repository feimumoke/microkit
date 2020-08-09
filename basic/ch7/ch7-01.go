package main

import (
	"fmt"
	"time"
)

/**
数据竞争
数据竞争会在两个以上的goroutine并发访问相同变量且至少一个为写操作时发生。
避免数据竞争的方法：
1、不要去写变量，可以在创建goroutine之前初始化
2、避免多个goroutine访问变量，其他的goroutine可以通过channel发送请求来查询或者更新变量。
-- 不要使用共享变量来通信，使用通信来共享数据

3、允许很多goroutine访问变量，但同一时刻最多只有一个访问
*/
var deposits = make(chan int)
var balances = make(chan int)

func Deposit(amount int) {
	deposits <- amount
}
func Balance() int {
	return <-balances
}

func bank() {
	var balance int
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case balances <- balance:
		}
	}
}

func main() {
	go bank()
	go Deposit(20)
	go Deposit(30)
	time.Sleep(100 * time.Millisecond)
	fmt.Println(Balance())

	bake := make(chan *Cake)
	ice := make(chan *Cake)

	go baker(bake)
	go icer(ice, bake)

	for cake := range ice {
		fmt.Println(*cake)
	}
}

//当变量无法绑定一个独立的goroutine时，串行绑定+channel传递信息仍然有效
type Cake struct {
	state string
}

func baker(cooked chan<- *Cake) {
	for {
		cake := new(Cake)
		cake.state = "cooked"
		cooked <- cake
	}
}

func icer(iced chan<- *Cake, cooked <-chan *Cake) {
	for cake := range cooked {
		cake.state = "iced"
		iced <- cake
	}
}
