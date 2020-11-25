package main

import (
	"fmt"
	"runtime"
)

func main() {
	score := 65
	if score >= 60 {
		fmt.Println("That's fine")
	}
	if score < 60 {
		fmt.Println("unlucky")
	}

	score2 := 77
	if score2 >= 60 {
		if score2 >= 60 && score < 70 {
			fmt.Println("C")
		}
		if score2 >= 70 && score2 < 80 {
			fmt.Println("B")
		}
		if score2 >= 80 && score2 < 90 {
			fmt.Println("A")
		}
		if score2 >= 90 {
			fmt.Println("A+")
		}
	} else {
		fmt.Println("D")
	}

	score3 := 77
	if score3 >= 90 {
		fmt.Println("优秀")
	} else if score3 >= 80 {
		fmt.Println("良好")
	} else if score3 >= 70 {
		fmt.Println("中等")
	} else if score3 >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}

	switch num := 16; num {
	case 2:
		fmt.Println("Binary")
	case 8:
		fmt.Println("Octal")
	case 10:
		fmt.Println("Decimal")
	case 16:
		fmt.Println("Hexadecimal")
	default:
		fmt.Println("Wrong")
	}
	fmt.Println("Stop")

	month := 5
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println("31")
	case 2:
		fmt.Println(28 / 29)
	default:
		fmt.Println("30")
	}

	switch num1 := 1; num1 {
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
		fallthrough
	case 4:
		fmt.Println(4)
	default:
		fmt.Println("not match")
	}
	fmt.Println("Stop")

	switch num := 1; num {
	case 1:
		fmt.Println("1")
		break
		fmt.Println("break后面代码都不执行")
		fallthrough
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
		fallthrough
	case 4:
		fmt.Println("4")
	default:
		fmt.Println("不是1,2,3,4")
	}
	fmt.Println("程序结束")

	//	fmt.Println("执行程序")
	//	i := 6
	//	if i == 6 {
	//		goto Loop
	//	}
	//	fmt.Println("if下面输出")
	//Loop:
	//	fmt.Println("loop")
	//
	//	fmt.Println("执行程序")
	//	i2 := 6
	//	if i2 == 6 {
	//		goto Loop2
	//		goto Loop1
	//	}
	//	fmt.Println("if下面输出")
	//Loop2:
	//	fmt.Println("loop")
	//Loop1: //报错:label Loop1 defined and not used
	//	fmt.Println("Loop1")

	runtime.GOMAXPROCS(1)
	ch1 := make(chan int, 1)
	ch2 := make(chan string, 1)
	ch1 <- 1
	ch2 <- "hello"
	select {
	case value := <-ch1:
		fmt.Println(value)
	case value := <-ch2:
		fmt.Println(value)
	}

	ch := make(chan int)
	for i := 1; i <= 5; i++ {
		go func(arg int) {
			ch <- arg
		}(i)
	}
	//如果是一直接受消息,应该是死循环for{},下面代码中是明确知道消息个数
	for i := 1; i <= 5; i++ {
		select {
		case c := <-ch:
			fmt.Println("取出数据", c)
		default:
			//没有default会出现死锁
		}
	}
	fmt.Println("程序执行结束")
}
