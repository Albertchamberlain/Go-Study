package main

import "net"

func main() {
	addr, _ := net.ResolveTCPAddr("tcp4", "localhost:8899")

	conn, _ := net.DialTCP("tcp", nil, addr)

	conn.Write([]byte("client send data"))

	conn.Close()

}
