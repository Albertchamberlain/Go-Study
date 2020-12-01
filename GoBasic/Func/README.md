
# 一. 概述

* 讨论值传递和引用传递时,其实就是看值类型变量和引用类型变量作为函数参数时,修改`形参`是否会影响到`实参`
* go语言只有值传递
* 在Go语言中五个引用类型变量,其他都是值类型
  * slice
  * map
  * channel
  * interface
  * func()
* 引用类型作为参数时,称为`浅拷贝`,形参改变,实参数跟随变化.因为传递的是`地址`,`形参`和`实参`都指向同一块地址
* 值类型作为参数时,称为`深拷贝`,形参改变,实参不变,因为传递的是值的副本,形参会新开辟一块空间,与实参指向不同
* 如果希望值类型数据在修改形参时实参跟随变化,可以把参数设置为指针类型 (Go引用传递实现的方式)

# 二.代码演示

* 值类型作为参数代码演示

```go
package main

import "fmt"

func demo(i int, s string) {
	i = 5
	s = "改变"
}

func main() {
	i := 1
	s := "原值"
	demo(i, s)
	fmt.Println(i, s) //输出:1 原值
}
```

* 引用传递代码示例

```go
package main

import "fmt"

func demo(arg []int) {
   arg[len(arg)-1] = 110
}

func main() {
   s := []int{1, 2, 3}
   demo(s)
   fmt.Println(s) //输出:[1 2 110]
}
```

* 如果希望值类型实参跟随形参变化,可以把值类型指针作为参数

```go
package main

import "fmt"

//行参指针类型
func demo(i *int, s string) {
   //需要在变量前面带有*表示指针变量
   *i = 5
   s = "改变"
}

func main() {
   i := 1
   s := "原值"
   //注意此处第一个参数是i的地址,前面&
   //s保留为值类型
   demo(&i, s)
   fmt.Println(i, s) //输出:5 原值
}
```

# 一. 函数

```
func 函数名(参数列表) 返回值{
  //函数体
}
```

* 函数调用的语法

```go
返回值:=函数名(参数)
```


# 二. 无参数无返回值函数

* 函数可以`有参数`也可以`没有参数`,可以`有返回值`也可以`没有返回值`

```go
func main() {
	demo1()
}
func demo1(){
	fmt.Println("执行demo1函数")
}

```

# 三.有参数函数

* 函数的参数可以有多个,且每个参数类型都可以不同
* 参数表示调用函数方想要给函数内部传递的值,给函数使用的.
* 声明函数时的参数叫做形参数,调用函数时参数叫做实参

```go
func main() {
	i:=5
	s:="xiaopang"
	show(s,i)
}
func show(name string,age int){
	fmt.Println("姓名:",name,"年龄",age)
}
```

# 四.有返回值函数

* 函数的返回值是给`调用方`返回的数据,给调用方使用的.
* 具有返回值的函数,必须要有`return`

```go
func main() {
	a, b := 1, 2
	fmt.Println(add(a,b))
}
func add(c, d int) int {
	return c + d
}
```

* 也可以在`返回值类型`前面`添加变量`,return关键字后不写内容,表示`变量是什么返回值什么`

```
func main() {
	a, b := 1, 2
	fmt.Println(add2(a,b))
}

func add2(c, d int) (sum int) {
	sum = c + d
	return
}
```


# 多返回值函数

* 在Go语言中`每个函数`声明时都可以定义成`多返回值函数`
* Go语言中`所有的错误`都是通过`返回值`返回的
* 声明多返回值函数的语法

```go
func 函数名(参数列表) (返回值,返回值){
  //函数体
}
```

* 调用函数的语法

```go
变量,变量:=函数名(参数)
```

* 调用函数时如果不想接收可以使用`下划线`占位

```go
变量,_:=函数名(参数)
```

* 理论上函数返回值个数可以无限多个,但是一般不去定义特别多个返回值(用`结构体`代替多返回值)

# 二.代码演示

* 函数的返回值可以`不接收`,表示执行函数
* 函数的返回值如果接收,用于接收返回值的变量个数与返回值个数相同
* 不想接收的使用占位符(_)占位

```go
func main() {
	//不接收函数返回值
	demo()

	//每个返回值都接收
	a, b := demo()
	fmt.Println(a, b)

	//不希望接收的返回值使用下划线占位
	c, _ := demo()
	fmt.Println(c)
}

func demo() (string, int) {
	return "xiaopang", 17
}
```

* 多返回值函数也可以给返回值`定义变量`,return后面就不需要编写内容

```
func demo() (name string, age int) {
	name = "smallming"
	age = 17
	return
}
```

#  可变参数函数
* Go语言支持可变参数函数
* 可变参数指调用参数时,参数的个数可以是任意个
* 可变参数必须在参数列表最后的位置,在参数名和类型之间添加`三个点 ...`表示可变参数函数

```go
func 函数(参数,参数,名称 ... 类型 ){
	//函数体
}
```

* 输出语句就是可变参数函数,源码如下

```go
func Println(a ...interface{}) (n int, err error) {
	return Fprintln(os.Stdout, a...)
}
```

* 声明函数时,在函数体把可变参数当作切片使用即可

# 二.代码示例

* 声明可变参数声明与调用

```go
func main() {
	demo("看书", "写代码", "看小胖")
}

func demo(hover ... string) {
	for a, b := range hover {
		fmt.Println(a, b)
	}
}
```

* 可变参数必须存在其他参数后面,一个函数不能有多个可变参数.
  * 因为前面普通参数个数是确定的,编译器知道,哪个实参给哪个形参

```go
func main() {
	demo("张三", "看书", "写代码", "看小胖")
}

func demo(name string, hover ... string) {
	fmt.Println(name, "的爱好是")
	for a, b := range hover {
		fmt.Println(a, b)
	}
}
```


#  匿名函数  (其他语言中有些叫做lambda λ 表达式)
* 匿名函数就是没有名称的函数
* 正常函数可以通过名称多次调用,而匿名函数由于没有函数名,所以大部分情况都是在当前位置声明并立即调用(函数变量除外)
* 匿名函数声明完需要调用,在函数结束大括号后面紧跟小括号

```go
func (){
  
}()//括号表示调用
```

* 匿名函数都是声明在其他函数内部

* 无参数匿名函数

```go
  func main(){
     func(){
        fmt.Println("这是匿名函数")
     }()//括号表示调用
  }
```

* 有参数匿名函数

```go
func main() {
   func(s string) {
      fmt.Println(s, "这是匿名函数")
   }("传递参数") //调用时传递参数
}
```

* 有参数有返回值匿名函数

```go
func main() {
	r := func(s string) int {
		fmt.Println(s, "这是匿名函数")
		return 110
	}("传递参数") //调用时传递参数
	fmt.Println(r)
}
```
# 函数变量

* 在Go语言中函数也是`一种类型`,**函数有多少种形式,函数变量就有多少种写法**

```go
	var a func()           //无参数无返回值
	var b func(int)        //有一个int类型参数
	var c func(int) string //有一个int类型参数和string类型返回值
	fmt.Println(a, b, c)   //输出:<nil> <nil> <nil>
```

* 定义完函数变量后,可以使用匿名函数进行赋值.也可以使用已经定义好的函数进行赋值
* 函数变量定义以后与普通函数调用语法相同,变量名就是普通函数声明的函数名

```go
func main() {
	var a func()
	a = func() {
		fmt.Println("执行函数")
	}   //注意此处没有括号,有括号表示调用函数,变量a就表示接收函数返回值
	a() //调用函数
	
	b := func(s string) {
		fmt.Println("执行第二个函数")
	}
	b("参数")

	//使用定义好的函数
	d := c
	d()
	//函数名称c也是一个变量
	c()
}

func c() {
	fmt.Println("c函数")
}
```

* 函数类型变量是除了`slice`、`map`、`channel`、`interface`外第五种引用类型

```go
func main() {
	var a func()
	a = b
	a()
	var c func()
	c = a
	c()
	fmt.Printf("%p %p", a, c)//输出地址相同
}

func b() {
	fmt.Println("b")
}
```

# 五.函数作为参数或返回值

* 变量可以作为函数的`参数`或`返回值`类型.而函数既然可以当做变量看待,`函数变量`也可以当做函数的`参数`或`返回值`
* 函数作为参数时,类型写成`对应的类型`即可

```go
func main() {
	a(func(s string) {
		fmt.Println(s)
	})
}

func a(b func(s string)) {
	fmt.Println("a执行")
	b("传递给s的内容")
}
```

* 函数作为返回值

```
func main() {
	//此时result指向返回值函数.
	result := a()
	//调用函数,才能获取结果
	fmt.Println(result())
}

func a() func() int {
	return func() int {
		return 110
	}
}
```
* 闭包解决`局部变量`不能被外部访问
* 是把`函数`当作`返回值`的一种应用



* 总体思想为:**在函数内部定义局部变量,把另一个函数当作返回值,局部变量对于返回值函数就相当于全局变量,所以多次调用返回值函数局部变量的值跟随变化**

```
package main

import "fmt"

func main() {
	//res其实就是test1返回值函数,和之前匿名函数变量一个道理
	res := test1()
	fmt.Println(res()) //输出2
	fmt.Println(res()) //输出3
	fmt.Println(res()) //输出4
}

//注意此处,返回值类型是func int
func test1() func() int {
	i := 1
	return func() int {
		i = i + 1
		return i
	}
}
```

* 如果重新调用test1()会重新声明及赋值局部变量i

```
package main

import "fmt"

func main() {
	f := test1()
	fmt.Println("f的地址", f) //输出匿名函数地址
	fmt.Println("f:", f()) //调用匿名函数输出2
	fmt.Println("f:", f()) //调用匿名函数输出3
	k := test1()
	fmt.Println("k的地址", k) //输出匿名函数地址,与f相等
	fmt.Println("k:", k()) //调用匿名函数输出2
	fmt.Println("f:", f()) //输出:4
	fmt.Println("k:", k()) //输出:3
}

func test1() func() int {
	i := 1
	return func() int {
		i++
		// 每调用一次test1()输出的地址不一样
		fmt.Println("i的地址:", &i)
		return i
	}
}

```
