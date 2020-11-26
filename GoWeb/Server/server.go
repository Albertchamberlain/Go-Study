package main

import (
	"fmt"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp4", "localhost:8899")

	lis, _ := net.ListenTCP("tcp4", addr)

	fmt.Println("Server running")

	//阻塞式
	conn, _ := lis.Accept()

	b := make([]byte, 1024)
	n, _ := conn.Read(b)

	fmt.Println("get data", string(b[:n]))
	conn.Write(append([]byte("Server"), b[:n]...))
	conn.Close()
}
