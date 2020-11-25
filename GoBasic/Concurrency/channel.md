# 一. channel

* 线程通信在每个编程语言中都是重难点,在Golang中提供了语言级别的goroutine之间通信:channel
* channel不同的翻译资料叫法不一样.常见的几种叫法
  * 管道
  * 信道
  * 通道
* channel是进程内通信方式,每个channel只能传递一个类型的值.这个类型需要在声明channel时指定
* channel在Golang中主要的两个作用
  * 同步
  * 通信
* Go语言中channel的关键字是chan
* 声明channel的语法

```go
var 名称 chan 类型
var 名称 chan <- 类型 //只写
var 名称 <- chan 类型//只读
名称:=make(chan int) //无缓存channel
名称:=make(chan int,0)//无缓存channel
名称:=make(chan int,100)//有缓存channel
```

* 操作channel的语法:(假设定义一个channel名称为ch)

```go
ch <- 值 //向ch中添加一个值
<- ch //从ch中取出一个值
a:=<-ch //从ch中取出一个值并赋值给a
a,b:=<-ch//从ch中取出一个值赋值给a,如果ch已经关闭或ch中没有值,b为false
```


# 二. 代码示例

* 简单无缓存通道代码示例
  * 此代码中如果没有从channel中取值c,d=<-ch语句,程序结束时go func并没有执行
  * 下面代码示例演示了同步操作,类似与WaitGroup功能,保证程序结束时goroutine已经执行完成
  * 向goroutine中添加内容的代码会阻塞goroutine执行,所以要把ch<-1放入到goroutine有效代码最后一行
  * 无论是向channel存数据还是取数据都会阻塞
  * close(channel)关闭channel,关闭后只读不可写

```go
package main

import (
   "fmt"
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
}
```

* 使用channel实现goroutine之间通信
  * channel其实就是消息通信机制实现方案,在Golang中没有使用共享内存完成线程通信,而是使用channel实现goroutine之间通信.

```go
package main

import (
   "fmt"
)

func main() {
   //用于goroutine之间传递数据
   ch := make(chan string)
   //用于控制程序执行
   ch2 := make(chan string)
   go func() {
      fmt.Println("执行第一个goroutine,等待第二个goroutine传递数据")
      content := <-ch
      fmt.Println("接收到的数据为:", content)
      ch2 <- "第一个"
   }()
   go func() {
      fmt.Println("进入到第二个,开始传递数据")
      ch <- "内容随意"
      close(ch)
      fmt.Println("发送数据完成")
      ch2 <- "第二个"
   }()
   result1 := <-ch2
   fmt.Println(result1, "执行完成")
   result2 := <-ch2
   fmt.Println(result2, "执行完成")
   fmt.Println("程序执行结束")
}
```

* 可以使用for range获取channel中内容
  * 不需要确定channel中数据个数

```go
func main() {
   ch:=make(chan string)
   ch2:=make(chan int)
   go func() {
      for i:=97;i<97+26;i++{
         ch <- strconv.Itoa(i)
      }
      ch2<-1
   }()

   go func() {
      for c := range ch{
         fmt.Println("取出来的",c)
      }
   }()
   <-ch2
   fmt.Println("程序结束")
}
```

* channel是安全的.多个goroutine同时操作时,同一时间只能有一个goroutine存取数据

```go
package main

import (
   "time"
   "fmt"
)

func main() {
   ch := make(chan int)

   for i := 1; i < 5; i++ {
      go func(j int) {
         fmt.Println(j, "开始")
         ch <- j
         fmt.Println(j, "结束")
      }(i)
   }

   for j := 1; j < 5; j++ {
      time.Sleep(2 * time.Second)
      <-ch
   }
}
```