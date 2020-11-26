package main

import (
	"fmt"
	"net"
	"strconv"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp4", "localhost:8899")

	for i := 0; i < 5; i++ {
		conn, _ := net.DialTCP("tcp4", nil, addr)
		_, _ = conn.Write([]byte("客户端的消息" + strconv.Itoa(i)))
		b := make([]byte, 1024)
		count, _ := conn.Read(b)
		fmt.Println("服务器发送回来的消息为:", string(b[:count]))
		_ = conn.Close()
	}
}
