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



# 一.线程休眠

* Go语言中main()函数为主线程(协程),程序是从上向下执行的
* 可以通过time包下的Sleep(n)让程序阻塞多少纳秒

```go
   fmt.Println("1")
   //单位是纳秒,表示阻塞多长时间
   //e9表示10的9次方
   time.Sleep(1e9)
   fmt.Println("2")
```

# 二.延迟执行

* 延迟指定时间后执行一次,但是需要注意在触发时程序没有结束

```go
  fmt.Println("开始")
   //2秒后执行匿名函数
   time.AfterFunc(2e9, func() {
      fmt.Println("延迟延迟触发")
   })
   time.Sleep(10e9)//一定要休眠,否则程序结束了
   fmt.Println("结束")
```
# 一.goroutine简介

* Golang中最迷人的一个优点就是从`语言层面`就支持并发
* 在Golang中的goroutine(`协程`)类似于其他语言的`线程`
* 并发和并行
  * 并行(parallelism)指不同的代码片段同时在不同的物理处理器上支持
  * 并发(concurrency)指同时管理多个事情,物理处理器上可能运行某个内容一半后就处理其他事情
  * 在一般看来并发的性能要好于并行.因为计算机的物理资源是固定的,较少的,而程序需要执行的内容是很多的.所以并发是”以较少的资源去去做更多事情”
* 几种主流并发模型
  * 多线程,每个线程只处理一个请求,只有请求结束后,对应的线程才会接收下一个请求.这种模式在高并发下,性能开销极大.
  * 基于回调的异步IO.在程序运行过程中可能产生大量回调导致维护成本加大,程序执行流程也不便于思维
  * 协程.不需要抢占式调用,可以有效提升线程任务的并发性,弥补了多线程模式的缺点;Golang在语言层面就支持,而其他语言很少支持
* goroutine的语法
  * 表达式可以是一条语句
  * 表达式也可以是`函数`,函数返回值即使有,也无效,当函数执行完成此goroutine自动结束

```go
	go 表达式
```

# 二. 代码示例

* 对比多次调用函数和使用goroutine的效果

```go
package main

import "fmt"
import "time"

func main() {
   //正常调用,输出3遍1 2 3 4 5(每个数字后换行)
   //for i:=1; i<=3; i++ {
   // go demo()
   //}

   /*
   添加go关键字后发现控制台什么也没有输出
   原因:把demo()设置到协程后没等到函数执行,主
   线程执行结束
    */
   for i := 1; i <= 3; i++ {
      go demo(i)
   }
}

func demo(index int) {
   for i := 1; i <= 5; i++ {
      fmt.Printf("第%d次执行,i的值为:%d\n", index, i)
   }
}
```

* 添加休眠等待goroutine执行结束
* 这种方式很大的问题就是休眠时间,如果休眠时间设置过小,可能goroutine并没有执行完成,如果休眠时间设置过大,影响程序执行执行.找到的本次执行的休眠时间,下次程序执行时这个休眠时间可能”过大”或”过小"
* 通过程序运行结果发现每次执行结果都不一定是一样的,因为每个demo()都是并发执行

```go
package main

import "fmt"
import "time"

func main() {
   //正常调用,输出3遍1 2 3 4 5(每个数字后换行)
   //for i:=1; i<=3; i++ {
   // go demo()
   //}

   /*
   添加go关键字后发现控制台什么也没有输出
   原因:把demo()设置到协程后没等到函数执行,主
   线程执行结束
    */
   for i := 1; i <= 3; i++ {
      go demo(i)
   }

   /*
   添加休眠,让主线程等待协程执行结束.
   具体休眠时间需要根据计算机性能去估计
   次数没有固定值
    */
   time.Sleep(3e9)
   fmt.Println("程序执行结束")
}

func demo(index int) {
   for i := 1; i <= 5; i++ {
      fmt.Printf("第%d次执行,i的值为:%d\n", index, i)
   }
}
```
# 一.WaitGroup简介

* Golang中sync包提供了基本同步基元,如互斥锁等.除了Once和WaitGroup类型.
* WaitGroup直译为等待组,其实就是`计数器`,只要计数器中有内容将一直阻塞
* 在Golang中WaitGroup存在于sync包中,在sync包中类型都是不应该被拷贝的.源码定义如下

```go
// A WaitGroup waits for a collection of goroutines to finish.
// The main goroutine calls Add to set the number of
// goroutines to wait for. Then each of the goroutines
// runs and calls Done when finished. At the same time,
// Wait can be used to block until all goroutines have finished.
//
// A WaitGroup must not be copied after first use.
type WaitGroup struct {
	noCopy noCopy

	// 64-bit value: high 32 bits are counter, low 32 bits are waiter count.
	// 64-bit atomic operations require 64-bit alignment, but 32-bit
	// compilers do not ensure it. So we allocate 12 bytes and then use
	// the aligned 8 bytes in them as state.
	state1 [12]byte
	sema   uint32
}
```

* Go语言标准库中WaitGroup只有三个方法
  * Add(delta int)表示向内部计数器添加增量(delta),其中参数delta可以是负数
  * Done()表示减少WaitGroup计数器的值,应当在程序最后执行.相当于Add(-1):blush:
  * Wait()表示阻塞直到WaitGroup计数器为0

```go
type WaitGroup
  func (wg *WaitGroup) Add(delta int)
  func (wg *WaitGroup) Done()
  func (wg *WaitGroup) Wait()
```


# 二.代码示例

* 使用WaitGroup可以有效解决`goroutine`未执行完成`主协程`执行完成,导致程序结束,goroutine未执行问题

```go
package main

import (
   "fmt"
   "sync"
)

var wg sync.WaitGroup

func main() {

   for i := 1; i <= 3; i++ {
      wg.Add(1)
      go demo(i)
   }
   //阻塞,知道WaitGroup队列中所有任务执行结束时自动解除阻塞
   fmt.Println("开始阻塞")
   wg.Wait()
   fmt.Println("任务执行结束,解除阻塞")

}

func demo(index int) {
   for i := 1; i <= 5; i++ {
      fmt.Printf("第%d次执行,i的值为:%d\n", index, i)
   }
   wg.Done()
}
```

# 一.互斥锁🔒

* Go语言中多个协程操作一个变量时会出现冲突的问题
* go run -race 可以查看竞争
* 可以使用`sync.Mutex`对内容加锁
* 互斥锁的使用场景
  * 多个goroutine访问同一个函数(代码段)
  * 这个函数操作一个全局变量
  * 为了保证共享变量安全性,值合法性
* 使用互斥锁模拟`售票窗口`

```go
package main

import (
   "fmt"
   "sync"
   "time"
   "math/rand"
)

var (
   //票数
   num = 100
   wg  sync.WaitGroup
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
   //设置随机数种子
   rand.Seed(time.Now().UnixNano())
   //计算器的起始值和票数相同
   wg.Add(4)
   go sellTicker(1)
   go sellTicker(2)
   go sellTicker(3)
   go sellTicker(4)
   wg.Wait()

   fmt.Println("所有票卖完")
}
```

# 一. 死锁🔒

* 在主`goroutine`中向无缓存channel添加内容或在主goroutine中向channel添加内容且添加内容的个数已经大于channel缓存个数就会产生死锁

```
fatal error : all goroutines are asleep -deadlock!
```

* 死锁:在程序中多个进程(Golang中goroutine)由于相互竞争资源而产生的阻塞(等待)状态,而这种状态一直保持下去,此时称这个线程是死锁状态
* 在Golang中使用无缓存channel时一定要注意.以下是一个最简单的死锁程序
  * 主协程中有ch<-1,无缓存channel无论添加还是取出数据都会阻塞goroutine,当前程序无其他代码,主goroutine会一直被阻塞下去,此时主goroutine就是死锁状态

```go
func main() {
   ch := make(chan int)
   ch <- 1
}
```

* 而下面代码就不会产生死锁
  * 通过代码示例可以看出,在使用`无缓存channel`时,特别要注意的是在`主协程`中有操作channel代码

```go
package main

import (
   "time"
   "fmt"
)

func main() {
   ch := make(chan int)
   go func() {
      ch <- 1
      fmt.Println("执行goroutine")
   }()
   time.Sleep(5e9)
   fmt.Println("程序执行结束")
}
```


# 二. 有缓存通道

* 创建一个有缓存通道 :roller_coaster:

```go
func main() {
   ch := make(chan int, 3) //缓存大小3,里面消息个数小于等于3时都不会阻塞goroutine
   ch <- 1
   ch <- 2
   ch <- 3
   ch <- 4 //此行出现死锁,超过缓存大小数量
}
```

* 在Golang中有缓存channel的缓存大小是不能改变的,但是只要不超过缓存数量大小,都不会出现阻塞状态

```go
package main

import "fmt"

func main() {
   ch := make(chan int, 3) //缓存大小3,里面消息个数小于等于3时都不会阻塞goroutine
   ch <- 1
   fmt.Println(<-ch)
   ch <- 2
   fmt.Println(<-ch)
   ch <- 3
   ch <- 4
   fmt.Println(len(ch))//输出2,表示channel中有两个消息
   fmt.Println(cap(ch))//输出3,表示缓存大小总量为3
}
```


# RWMutex读写锁🔒

* RWMutex 源码如下

```go
type RWMutex struct {
	w           Mutex  // held if there are pending writers
	writerSem   uint32 // semaphore for writers to wait for completing readers
	readerSem   uint32 // semaphore for readers to wait for completing writers
	readerCount int32  // number of pending readers
	readerWait  int32  // number of departing readers
}
```

* Go语言标准库中API如下

```go
type RWMutex
  func (rw *RWMutex) Lock()//禁止其他协程读写
  func (rw *RWMutex) Unlock()
  func (rw *RWMutex) RLock()//禁止其他协程写入,只能读取
  func (rw *RWMutex) RUnlock()
  func (rw *RWMutex) RLocker() Locker
```

* Go语言中的`map`不是线程安全的,`多个goroutine`同时操作会出现错误.
* RWMutex可以添加多个读锁或一个写锁.读写锁不能同时存在.
  * map在并发下读写就需要结合读写锁完成
  * 互斥锁表示锁的代码同一时间只能有一个人goroutine运行,而读写锁表示在锁范围内数据的读写操作

```go
package main

import (
   "fmt"
   "sync"
   "strconv"
)

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
}
```

