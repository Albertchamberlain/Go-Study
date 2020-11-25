package main

import (
	"fmt"
	"strconv"
)

func main() {
	ch := make(chan int)
	go func() {
		fmt.Println("进入goroutine")
		// 添加一个内容后控制台输出:1 true
		//ch<-1
		//关闭ch控制台输出:0 false
		close(ch)
	}()
	c, d := <-ch
	fmt.Println(c, d)
	fmt.Println("程序执行结束")

	//用于goroutine之间传递数据
	chh := make(chan string)
	//用于控制程序执行
	ch2 := make(chan string)
	go func() {
		fmt.Println("执行第一个goroutine,等待第二个goroutine传递数据")
		content := <-chh
		fmt.Println("接收到的数据为:", content)
		ch2 <- "第一个"
	}()
	go func() {
		fmt.Println("进入到第二个,开始传递数据")
		chh <- "214"
		close(chh)
		fmt.Println("发送数据完成")
		ch2 <- "第二个"
	}()
	result1 := <-ch2
	fmt.Println(result1, "执行完成")
	result2 := <-ch2
	fmt.Println(result2, "执行完成")
	fmt.Println("程序执行结束")

	chhh := make(chan string)
	ch22 := make(chan int)
	go func() {
		for i := 97; i < 97+26; i++ {
			chhh <- strconv.Itoa(i)
		}
		ch22 <- 1
	}()

	go func() {
		for c := range chhh {
			fmt.Println("取出来的", c)
		}
	}()
	<-ch22
	fmt.Println("程序结束")

}
