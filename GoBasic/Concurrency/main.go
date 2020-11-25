package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	fmt.Println("1")
	//单位是纳秒,表示阻塞多长时间
	//e9表示10的9次方
	time.Sleep(1e9)
	fmt.Println("2")

	fmt.Println("开始")
	//2秒后执行匿名函数
	time.AfterFunc(2e9, func() {
		fmt.Println("延迟触发")
	})
	time.Sleep(10e9) //一定要休眠,否则程序结束了
	fmt.Println("结束")

	/*
	   添加go关键字后发现控制台什么也没有输出
	   原因:把demo()设置到协程后没等到函数执行,主
	   线程执行结束
	*/
	for i := 1; i <= 3; i++ {
		go demo(i)
	}

	for i := 1; i <= 3; i++ {
		go demo(i)
	}

	time.Sleep(3e9)
	fmt.Println("程序执行结束")

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go demo2(i)
	}

	fmt.Println("开始阻塞")
	wg.Wait()
	fmt.Println("任务执行结束,解除阻塞")
}
func demo2(index int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("第%d次执行,i的值为:%d\n", index, i)
	}
	wg.Done()
}

func demo(index int) {
	for i := 1; i <= 5; i++ {
		fmt.Printf("第%d次执行,i的值为:%d\n", index, i)
	}
}
