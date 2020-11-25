package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var wg sync.WaitGroup
var (
	//票数
	num = 100
	wgg sync.WaitGroup
	//互斥锁
	mu sync.Mutex
)

func sellTicker(i int) {
	defer wg.Done()
	for {
		//加锁,多个goroutine互斥
		mu.Lock()
		if num >= 1 {
			fmt.Println("第", i, "个窗口卖了", num)
			num = num - 1
		}
		//解锁
		mu.Unlock()

		if num <= 0 {
			break
		}
		//添加休眠,防止结果可能出现在一个goroutine中
		time.Sleep(time.Duration(rand.Int63n(1000) * 1e6))
	}
}

func main() {

	var rwm sync.RWMutex
	m := make(map[string]string)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(j int) {
			//没有锁在map时可能出现问题
			rwm.Lock()
			m["key"+strconv.Itoa(j)] = "value" + strconv.Itoa(j)
			fmt.Println(m)
			rwm.Unlock()
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("程序结束")

	//设置随机数种子
	rand.Seed(time.Now().UnixNano())
	//计算器的起始值和票数相同
	wgg.Add(4)
	go sellTicker(1)
	go sellTicker(2)
	go sellTicker(3)
	go sellTicker(4)
	wgg.Wait()

	fmt.Println("所有票卖完")

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
