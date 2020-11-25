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