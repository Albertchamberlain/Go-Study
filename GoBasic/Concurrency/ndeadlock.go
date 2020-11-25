package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		ch <- 1
		fmt.Println("执行goroutine")
	}()
	time.Sleep(5e9)
	fmt.Println("程序执行结束")

	chh := make(chan int, 3) //缓存大小3,里面消息个数小于等于3时都不会阻塞goroutine
	chh <- 1
	fmt.Println(<-chh)
	chh <- 2
	fmt.Println(<-chh)
	chh <- 3
	chh <- 4
	fmt.Println(len(chh)) //输出2,表示channel中有两个消息
	fmt.Println(cap(chh)) //输出3,表示缓存大小总量为3
}
