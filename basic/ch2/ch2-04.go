package main

import "fmt"

/**
struct
go语言函数所有参数都是值拷贝传入的，因此大结构体通常传递指针
如果结构体的所有的成员都是可比较的，那么结构体也是可比较的
*/

//结构图嵌入
type Point struct {
	X, Y int
}

type Circle struct {
	Center Point
	Radius int
}

type Wheel struct {
	Circle Circle
	Spokes int //辐条
}

func createWheel() {
	var w Wheel
	w.Circle.Center.X = 8
	w.Circle.Center.Y = 8
	w.Circle.Radius = 5
	w.Spokes = 20
}

/**
go语言可以使用只声明匿名成员的形式，匿名成员的数据类型必须是命名的类型
或者指向一个命名的类型的指针。使用匿名嵌入可以直接访问叶子属性
*/

type Circle2 struct {
	Point
	Radius int
}

type Wheel2 struct {
	Circle2
	Spokes int
}

func createWheel2() {
	var w Wheel2
	w.X = 8
	w.Y = 8
	w.Radius = 5
	w.Spokes = 20
}

/**
结构体字面值必须遵循形状类型声明时的结构，所以我们只能用下面的两种语法
%v参数包含的#副词表示和go语言类似的语法打印值，对于结构体来说包含每个成员的名字
*/
func embed() {
	w := Wheel2{Circle2{Point{8, 8}, 5}, 20}
	w = Wheel2{
		Circle2: Circle2{Point: Point{X: 8, Y: 8}, Radius: 5},
		Spokes:  20,
	}
	fmt.Printf("%#v\n", w)
	w.X = 42
	fmt.Printf("%#v\n", w)
}

/**
一个结构体可能同时包含导出和未导出的成员。
一个命名为S的结构体类型不能呢个再包含S类型的成员，因为一个聚合的值不能包含它自身，该限制同样适用于数组。
但是可以包含*S类型的指针，这可以创建递归的数据结构
二插树插入排序：
*/
type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func main() {
	embed()
	nums := []int{10, 45, 4, -10, 500, 56, 0, 1000, -5, 32}
	Sort(nums)
	fmt.Printf("%v", nums)
}
