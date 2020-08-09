package main

import (
	"bytes"
	"io"
)

/**
接口

一个包内nil指针的接口不是nil接口
main调用f时传入的是 *bytes.Buffer的指针，所以out的动态值nil。
然而，其动态类型是 *bytes.Buffer，意思就是out是一个包含空指针的非空接口。
解决方法是将main中的buf类型改为io.Writer
*/

const debug = false

func main() {
	//var buf *bytes.Buffer
	var buf io.Writer
	if debug {
		buf = new(bytes.Buffer)
	}
	f(buf)

	//...
}

func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("done!\n")) //invalid memory address or nil pointer dereference
	}
}
