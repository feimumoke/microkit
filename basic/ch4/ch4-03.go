package main

import "sync"

/**
通过嵌入结构体扩展类型
当编译器解析一个选择器到方法时，会首先去找直接定义在这个类型的方法
然后找内嵌字段引入的方法，如果选择器有二义性编译器报错。
比如同一级里有两个同名方法。
方法只能定义在命名类型或者指向类型方指针上定义。
由于内嵌，可以给匿名struct类型定义方法
*/

var cache = struct {
	sync.Mutex
	mapping map[string]string
}{mapping: make(map[string]string)}

func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}
