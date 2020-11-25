# 一.defer使用
功能类似于`finally`
* Go语言中defer可以完成延迟功能,当前函数执行完成后执行defer功能
* defer最常用的就是关闭连接(数据库连接,文件等)可以打开连接后代码紧跟defer进行关闭,后面在执行其他功能
  * 在很多语言中要求必须按照顺序执行,也就是必须把关闭代码写在最后,但是经常会忘记关闭导致内存溢出,而Golang中defer很好的解决了这个问题.无论defer写到哪里都是最后执行

```go
func main() {
   fmt.Println("打开连接")
   defer func(){
      fmt.Println("关闭连接")
   }()
   fmt.Println("进行操作")
   //输出:打开连接 进行操作 关闭连接
}
```

# 二.多个defer

* 多重defer采用`栈`结构执行,先产生后执行
* 在很多代码结构中都可能出现产生多个对象,而程序希望这些对象倒序关闭,多个defer正好可以解决这个问题

```go
func main() {
   fmt.Println("打开连接A")
   defer func(){
      fmt.Println("关闭连接A")
   }()
   fmt.Println("打开连接B")
   defer func(){
      fmt.Println("关闭连接B")
   }()
   fmt.Println("进行操作")
   //输出:打开连接A 打开连接B 进行操作 关闭连接B 关闭连接A
}
```

# 三.defer和return结合

* defer与return同时存在时,要把return理解成两条执行结合(不是原子指令),一个指令是给返回值`赋值`,另一个指令返回`跳出函数`

* defer和return时整体执行顺序
  * 先给返回值赋值
  * 执行defer
  * 返回跳出函数

* 没有定义返回值接收变量,执行defer时返回值已经赋值

```go
func f() int{
	i:=0
	defer func(){
		i=i+2
	}()
	return i
}

func main() {
	fmt.Println(f())//输出:0
}
```

* 声明接收返回值变量,执行defer时修改了返回值内容.
  * 由于return后面没有内容,就无法给返回值赋值,所以执行defer时返回值才有内容

```go
func f() (i int){
	defer func(){
		i=i+2
	}()
	return
}
func main() {
	fmt.Println(f())//输出:2
}
```