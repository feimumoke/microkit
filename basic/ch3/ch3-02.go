package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/**
错位处理策略
1) 传播错误   可以直接将错位返回或者构造一个新的错误信息
2) 重新尝试   如果错误是偶然的，或不可预知的问题导致。
3) 输出错位信息、结束程序    只在main中执行，对应函数库应该向上传播
4) 输出错误信息，继续执行
5) 忽略错误
*/

func eof() {
	in := bufio.NewReader(os.Stdin)
	for {
		r, _, err := in.ReadRune()
		if err == io.EOF {
			break //finished
		}
		if err != nil {
			fmt.Errorf("read failed: %v", err)
		}
	}
}
