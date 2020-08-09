package main

import (
	"image"
	"sync"
)

/**
sync.Once惰性初始化

早期的loadIcons存在的bug：
多goroutine访问时会初始化多次或者icons为非空时还没初始化完成
*/

var icons map[string]image.Image

func loadIcons() {
	icons = map[string]image.Image{
		"spades.png":   loadIcon("spades.png"),
		"hearts.png":   loadIcon("hearts.png"),
		"diamonds.png": loadIcon("diamonds.png"),
		"clubs.png":    loadIcon("clubs.png"),
	}
}

func loadIcon(s string) image.Image {
	return image.Rectangle{image.Point{1, 2}, image.Point{3, 4}}
}

func Icons(name string) image.Image {
	if icons == nil {
		loadIcons()
	}
	return icons[name]
}

/**
保证所有goroutine能够观察到loadIcons效果的最简单方式是加mutex同步
*/

var mut sync.RWMutex

func Icon(name string) image.Image {
	mut.RLock()
	if icons != nil {
		icon := icons[name]
		mut.RUnlock()
		return icon
	}
	mut.RUnlock()
	mut.Lock()
	if icons == nil {
		loadIcons()
	}
	icon := icons[name]
	mut.Unlock()
	return icon
}

/**
上面的代码能够更好的并发，但是太复杂容易出错。
概念上来讲，一次性的初始化需要一个互斥量mutex和一个boolean变量来记录初始化是不是已经完成了；互斥量用来保护boolean变量和客户端数据结构。
每一次对Do()的调用都会锁定mutex,并会检查boolean变量.
mutex同步会保证loadIcons对内存产生的效果能够对所有goroutine可见
*/
var once sync.Once

func Icon3(name string) image.Image {
	once.Do(loadIcons)
	return icons[name]
}
