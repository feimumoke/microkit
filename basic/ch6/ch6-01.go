package main

import (
	"io"
	"log"
	"net"
	"time"
)

/**
并发Clock服务
handleConn函数会处理一个完整的客户端连接, 如果不用goroutine会一直处理完回到主函数继续处理下一个链接
*/

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
