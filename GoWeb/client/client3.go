package main

import (
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
)

type User struct {
	Username      string
	OtherUsername string
	Msg           string
	ServerMsg     string
}

var (
	user = new(User) //当前登录用户信息
	wg   sync.WaitGroup
)

func main() {
	wg.Add(1)
	fmt.Println("请输入您的账号:")
	fmt.Scanln(&user.Username)
	fmt.Println("请输入要给谁发送信息:")
	fmt.Scanln(&user.OtherUsername)

	addr, _ := net.ResolveTCPAddr("tcp4", "localhost:8899")
	conn, _ := net.DialTCP("tcp4", nil, addr)
	//发送消息
	go func() {
		fmt.Println("请输入您要发送的消息:(只提示一次)")
		for {
			fmt.Scanln(&user.Msg)
			if user.Msg == "exit" {
				conn.Close()
				wg.Done()
				os.Exit(0)
			}
			conn.Write([]byte(fmt.Sprintf("%s-%s-%s-%s", user.Username, user.OtherUsername, user.Msg, user.ServerMsg)))
		}
	}()
	//接收消息
	go func() {
		for {
			b := make([]byte, 1024)
			count, _ := conn.Read(b)
			array := strings.Split(string(b[:count]), "-")
			user2 := new(User)
			user2.Username = array[0]
			user2.OtherUsername = array[1]
			user2.Msg = array[2]
			user2.ServerMsg = array[3]
			if user2.ServerMsg != "" {
				fmt.Println("\t\t服务器的消息:", user2.ServerMsg)
			} else {
				fmt.Println("\t\t", user2.Username, ":", user2.Msg)
			}
		}
	}()
	wg.Wait()
}
