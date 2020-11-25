# 一.日志简介


* 有三种级别日志输出
  * Print() 输出日志信息
  * Panic()  打印日志信息,并触发panic,日志信息为Panic信息
  * Fatal()  打印日志信息后调用os.Exit(1)
* 所有日志信息打印时都带有时间,且颜色为红色
* 每种级别日志打印都提供了三个函数
  * Println()
  * Print()
  * Printf()
* 日志文件扩展名为log

# 二.普通日志信息打印

* 官方源码如下

```go
func Println(v ...interface{}) {
	std.Output(2, fmt.Sprintln(v...))
}
```

* 直接使用log包调用Println()即可

```go
log.Println("打印日志信息")
```

# 三.Panic日志信息打印

* 通过源码可以看出在日志信息打印后调用了panic()函数,且日志信息为panic信息

```go
// Panicln is equivalent to Println() followed by a call to panic().
func Panicln(v ...interface{}) {
	s := fmt.Sprintln(v...)
	std.Output(2, s)
	panic(s)
}
```

* 执行后输出日志信息,同时也会触发panic

```go
log.Panicln("打印日志信息")
```

# 四.致命日志信息

* 打印日志后,终止程序

```go
// Fatal is equivalent to Print() followed by a call to os.Exit(1).
func Fatal(v ...interface{}) {
	std.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}
```

* 执行日志打印后,程序被终止

```go
log.Fatal("打印日志信息")
```

# 五.打印日志信息到文件中

* Go语言标准库支持输出日志信息到文件中.
* 输出日志时的几种状态

```go
const (
	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	Ltime                         // the time in the local time zone: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
)
```

* 代码如下

```go
	f, _ := os.OpenFile("D:/golog.log", os.O_APPEND|os.O_CREATE, 07777)
	defer f.Close()
	logger := log.New(f, "[info]\t", log.Ltime)
	logger.Println("输出日志信息")
```