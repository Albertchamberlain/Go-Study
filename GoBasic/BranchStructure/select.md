# 一. select简介

* Golang中select和switch结构特别像,但是select中case的条件只能是I/O
* select 的语法(condition是条件)

```go
select{
  case condition:
  case condition:
  default:
}
```

* select执行过程:
  * 每个case必须是一个`IO`操作
  * 哪个case可以执行就执行哪个
  * 多个case都可以执行,`随机`执行一个
  * 所有case都不能执行时,执行`default`
  * 所有case都不能执行,且没有default,将会`阻塞`
* 代码示例

```go
func main() {
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
}
```

* select多和for循环结合使用,下面例子演示出了一直在接收消息的例子

```go
package main

import (
	"fmt"
)

func main() {
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

```

* `break`可以对select生效,如果for中嵌套select,break选择`最近`结构