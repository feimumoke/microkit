package main

import "fmt"

/**
方法选择器会返回一个方法“值” -> 一个将方法绑定到特定接收器变量的函数。
这个函数之后可以不用再指定接收器，只需要传入函数的参数。

当T是一个类型时，方法表达式会写成T.f 或者 (*T).f,
会返回一个函数“值”，这种函数会将其第一个参数用作接收器。
*/
/*
func f() {
	p := Point{1, 2}
	q := Point{3, 4}
	disP := p.Distance

	fmt.Println(disP(q))

	distance := Point.Distance
	fmt.Println(distance(p, q))
}
*/

type Point3 struct {
	X, Y float64
}

func (p Point3) Add(q Point3) Point3 {
	return Point3{p.X + q.X, q.Y + p.Y}
}

func (p Point3) Sub(q Point3) Point3 {
	return Point3{p.X - q.X, p.Y - q.Y}
}

type Path2 []Point3

func (path Path2) TranslateBy(offset Point3, add bool) {
	var op func(p, q Point3) Point3
	if add {
		op = Point3.Add
	} else {
		op = Point3.Sub
	}
	for i := range path {
		path[i] = op(path[i], offset)
	}
}

func main() {
	pt := Path2{Point3{1, 2}, Point3{2, 3}}
	pt.TranslateBy(Point3{4, 5}, true)

	fmt.Println(pt)
}
