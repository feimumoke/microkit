package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

/**
并发的非阻塞缓存
*/

type Memo struct {
	f     Func
	cache map[string]result
}
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

//并发不安全
func (memo *Memo) Get(key string) (interface{}, error) {
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err
}

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

func single() {
	m := New(httpGetBody)
	urls := []string{"http://www.baidu.com", "https://mp.weixin.qq.com/s/oexktPKDULqcZQeplrFunQ"}
	for _, url := range urls {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
	}
}

func concurrency() {
	m := New(httpGetBody)
	var n sync.WaitGroup
	urls := []string{"http://www.baidu.com", "http://www.baidu.com", "https://github.com/AobingJava/JavaFamily", "https://www.ai66.cc/dongzuopian/", "https://mp.weixin.qq.com/s/oexktPKDULqcZQeplrFunQ"}
	for _, url := range urls {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
			n.Done()
		}(url)
	}
	n.Wait()
}

/**
最简单的cache并发是基于监控的同步，只需要给Memo加上mutex，
Get开始时获取锁，return时释放锁。
但这一点改变丧失了并发性能优点。
*/
type Memo2 struct {
	f     Func
	mu    sync.Mutex
	cache map[string]result
}

func (memo *Memo2) Get2(key string) (interface{}, error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	memo.mu.Unlock()
	return res.value, res.err
}

//提升性能，但会有同时获取相同url的Get操作
func (memo *Memo2) Get3(key string) (interface{}, error) {
	memo.mu.Lock()
	res, ok := memo.cache[key]
	memo.mu.Unlock()
	if !ok {
		res.value, res.err = memo.f(key)
		//多个goroutine会计算f并更新map
		memo.mu.Lock()
		memo.cache[key] = res
		memo.mu.Unlock()
	}
	return res.value, res.err
}

//避免多余工作版本（并发、不重复、无阻塞）
type entry struct {
	res   result
	ready chan struct{}
}

type Memo3 struct {
	f     Func
	mu    sync.Mutex
	cache map[string]*entry
}

func New3(f Func) *Memo3 {
	return &Memo3{f: f, cache: make(map[string]*entry)}
}
func (memo *Memo3) Get(key string) (value interface{}, err error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()
		e.res.value, e.res.err = memo.f(key)
		close(e.ready) //广播条件已经就绪
	} else {
		memo.mu.Unlock()
		<-e.ready //等待就绪条件（f的执行）
	}
	return e.res.value, e.res.err
}

/**
无锁版本
*/

type request struct {
	key      string
	response chan<- result
}
type Memo4 struct {
	requests chan request
}

func New4(f Func) *Memo4 {
	memo := &Memo4{requests: make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo4) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo4) Close() {
	close(memo.requests)
}

func (memo *Memo4) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		e := cache[req.key]
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key)
		}
		go e.deliver(req.response)
	}
}
func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready) //广播
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready
	response <- e.res
}

func nolock() {
	m := New4(httpGetBody)
	urls := []string{"http://www.baidu.com", "http://www.baidu.com", "https://github.com/AobingJava/JavaFamily", "https://www.ai66.cc/dongzuopian/", "https://mp.weixin.qq.com/s/oexktPKDULqcZQeplrFunQ"}
	for _, url := range urls {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("%s, %s, %d bytes\n", url, time.Since(start), len(value.([]byte)))
	}
}
func main() {
	//concurrency()
	nolock()
}
