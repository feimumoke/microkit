package main

import "sync"

/**
sync.Mutex互斥锁
可以用一个容量只有1的channel保证最多只有一个goroutine在同一时刻访问共享变量。
一个只能为1或者0的信号量称为二元信号量
*/
var (
	sema    = make(chan struct{}, 1)
	balance int
)

func Deposit2(amount int) {
	sema <- struct{}{} // acquire token
	balance = balance + amount
	<-sema //release token
}

func Balance2() int {
	sema <- struct{}{}
	b := balance
	<-sema
	return b
}

//sync包里的Mutex类型直接支持，Lock方法能获取token（锁）
//lock和unlock之间的代码中的内容goroutine可以随便读取或者修改，这个代码段叫做临界区。
//go里没有重入锁，没有办法对已经上锁的mutex再上锁--这将会导致死锁
var (
	mu       sync.Mutex
	balance3 int
)

func Deposit3(amount int) {
	mu.Lock() // acquire token
	balance = balance + amount
	mu.Unlock() //release token
}

func Balance3() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}

/**
锁总是会被保持并去做实际的操作，避免重入
*/

func Withdraw(amount int) bool {
	mu.Lock()
	defer mu.Unlock()
	deposit(-amount)
	if balance < 0 {
		deposit(amount)
		return false
	}
	return true
}
func Deposit4(amount int) {
	mu.Lock() // acquire token
	defer mu.Unlock()
	deposit(amount)
}

func Balance4() int {
	mu.Lock()
	defer mu.Unlock()
	return balance
}

func deposit(amount int) {
	balance += amount
}

/**
sync.RWMutex读写锁
RLock只能在临界区共享变量没有任何写入操作时可用
*/

var rw sync.RWMutex

func Balabce5() int {
	rw.RLock()
	defer rw.RUnlock()
	return balance
}
