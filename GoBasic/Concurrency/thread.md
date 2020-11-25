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