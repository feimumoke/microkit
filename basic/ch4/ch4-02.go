package main

import "fmt"

/**
nil 是一个合法的接收器类型
方法也可以用nil指针作为其接收器

*/

// A nil *IntList represents the empty list
type IntList struct {
	Value int
	Tail  *IntList
}

func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}

/**
一般如果一个类有一个指针作为接收器的方法，那么所有的方法都必须有一个指针接收器。
只有类型和指向他们的指针才可能是出现在接收器声明里的两种接收器。
如果一个类型是一个指针的话，是不允许出现在接收器中的
*/

type P *int

/*
func (P)f()  {

}
*/

type Point2 struct {
	X, Y float64
}

func (p *Point2) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

func main() {
	p := Point2{1, 2}
	/**
	如果一个接收器是一个Ponit类型的变量，并且其方法需要一个Point指针作为接收器,
	编译器会隐式的用&p去调用ScaleBy这个方法。这种简写只适用于“变量”，包括struct
	里的字段，array以及slice里的元素。但是临时变量的内存地址无法获取，不能使用这种方法。
	*/
	p.ScaleBy(2)
	fmt.Println(p)
	//Point2{2,3}.ScaleBy(3) //cannot call pointer method on Point2 literal
}

/**
每一个合法表达式调用，需要满足以下三种中的一种：
1、接收器的实参和形参是相同类型，比如两者都是T或者*T
2、接收器实参是T，形参是*T,这种情况编译器会隐式的取变量地址
3、接收器实参是*T,形参是T，编译器会隐式解引用，取到指针指向的变量
*/
