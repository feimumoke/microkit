package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

/**
Defer
defer常用在处理成对的操作，如打开、关闭、链接、断开链接、加锁，释放锁。
释放资源的defer应该放在请求资源语句后
*/

func ReadFile(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ioutil.ReadAll(file)
}

/**
defer 常用于记录何时进入和退出函数，defer语句后面要加上园括号
*/

func bigSlowOper() {
	defer trace("bigSlowOper")()
	time.Sleep(10 * time.Second)
}

func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func main() {
	bigSlowOper()
	fmt.Println(triple(8))
}

/**
被延时调用的函数可以修改函数返回值
*/

func triple(x int) (result int) {
	defer func() { result += x }()
	return x + x
}

/**
循环体中的defer不会立即执行
一种解决方法是将defer语句移到另外一个函数
*/

func operfile(filenames []string) error {
	for _, filename := range filenames {
		if err := doFile(filename); err != nil {
			return err
		}
	}
	return nil
}

func doFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	// process
	return nil
}
