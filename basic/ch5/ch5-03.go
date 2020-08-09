package main

import "io"

/**
通过类型断言询问行为
因为Write方法需要传入一个byte切片而我们希望写入的值是一个字符串，
所以我们需要使用[]byte(...)进行转换。这个转换分配内存并且做一个拷贝，
但是这个拷贝在转换后几乎立马就被丢弃掉。
让我们假装这是一个web服务器的核心部分并且我们的性能分析表示这个内存分配使服务器的速度变慢。
这里我们可以避免掉内存分配么？
*/

func writeHeader(w io.Writer, contentType string) error {
	if _, err := w.Write([]byte("Content-Type:")); err != nil {
		return err
	}
	if _, err := w.Write([]byte(contentType)); err != nil {
		return err
	}
	return nil
}

/**
这个io.Writer接口告诉我们关于w持有的具体类型的唯一东西：就是可以向它写入字节切片。
如果我们回顾net/http包中的内幕，我们知道在这个程序中的w变量持有的动态类型也有一个允许字符串高效写入的WriteString方法；这个方法会避免去分配一个临时的拷贝。
不能对任意io.Writer类型的变量w，假设它也拥有WriteString方法。但是我们可以定义一个只有这个方法的新接口并且使用类型断言来检测是否w的动态类型满足这个新接口。

这个例子的神奇之处在于，没有定义了WriteString方法的标准接口，
也没有指定它是一个所需行为的标准接口。
一个具体类型只会通过它的方法决定它是否满足stringWriter接口，
而不是任何它和这个接口类型所表达的关系。
它的意思就是上面的技术依赖于一个假设，这个假设就是：如果一个类型满足下面的这个接口，
然后WriteString(s)方法就必须和Write([]byte(s))有相同的效果。

interface {
	io.Writer
	WriteString(s string) (n int, err error)
}

*/

func writeString(w io.Writer, s string) (n int, err error) {
	type stringWriter interface {
		WriteString(string) (n int, err error)
	}
	if sw, ok := w.(stringWriter); ok {
		return sw.WriteString(s) //avoid a copy
	}
	return w.Write([]byte(s))
}
