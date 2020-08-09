package main

import "math"

/**
方法声明
在函数声明时，在名字前面放一个变量，就是一个方法。
附加的参数叫做方法的接收器。由于方法和字段都是在同一命名空间，所以声明的方法
名称不能跟字段一样。
*/

type Point struct {
	X, Y float64
}

// 传统函数
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// Point的一个方法
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

/**
可以给同一包内的任意命名类型定义方法，只要这个命名类型的底层类型不是指针或者interface
*/

type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}
