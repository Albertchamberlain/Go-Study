package main

import (
	"fmt"
	"net"
)

func main() {
	addr, _ := net.ResolveTCPAddr("tcp4", "localhost:8899")

	conn, _ := net.DialTCP("tcp", nil, addr)

	conn.Write([]byte("Client send data"))
	b := make([]byte, 1024)
	count, _ := conn.Read(b)
	fmt.Println("data from Server", b[:count])
	conn.Close()

}
