package main

import (
	"fmt"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp4", "localhost:8899")
	lis, _ := net.ListenTCP("tcp4", addr)
	for {
		conn, _ := lis.Accept()
		go func() {
			b := make([]byte, 1024)
			count, _ := conn.Read(b)
			fmt.Println("服务器接收到的消息为:", string(b[:count]))
			_, _ = conn.Write(append([]byte("server:"), b[:count]...))
			_ = conn.Close()
		}()
	}
}
