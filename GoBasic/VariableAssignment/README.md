## 变量命名规则

* 以字母或下划线开头(Go语言中多不以_开头)

* 后面可以是任意数量的字符、数字和下划线

* 区分大小写

* 不能是关键字(关键字具备特定含义),下面是Go语言的关键字

  | 关键字如下 |             |        |           |        |
  |:---------:|:-----------:|:------:|:---------:|:------:|
  |   break   |   default   |  func  | interface | select |
  |   case    |    defer    |   go   |    map    | struct |
  |   chan    |    else     |  goto  |  package  | switch |
  |   const   | fallthrough |   if   |   range   |  type  |
  | continue  |     for     | import |  return   |  var   |

* 可以是保留字,但是建议不使用保留字做为变量名,下面是Go语言的保留字

  |   保留字如下   |            |           |         |         |
  |:-------------:|:----------:|:---------:|:-------:|:-------:|
  |     true      |   false    |   iota    |   nil   |   int   |
  |     int8      |   int16    |   int32   |  int64  |  unit   |
  |     unit8     |   unit16   |  unit32   | unitptr | float32 |
  |    float64    | complex128 | complex64 |  bool   |  byte   |
  |     rune      |   string   |   error   |  make   |   len   |
  |      cap      |    new     |  append   |  copy   |  close  |
  | deletecomplex |    real    |   imag    |  panic  |         |
  |    recover    |            |           |         |         |

* 在同一范围内不允许出现同名变量

* Go语言要求变量声明后至少使用一次(赋值不属于使用)

## 单个变量声明及赋值

* 先声明后赋值(声明后开辟内存,不同类型变量都有不同初值)

```go
//语法:
//1. 声明
var 变量名 类型
//2. 赋值
变量名=值

//示例:
var smallming string
smallming = "英文名"
```

* 声明并赋值(此方式不建议)

```go
//语法:
var 变量名 类型 = 值

//示例
var smallming string = "英文名"
```

* 声明并赋值(省略类型,变量类型取决于值的类型)

```go
//语法:
var 变量名 = 值

//示例:
var smallming = "英文名"
```

* 短变量(只能在函数内使用)

```go
//语法:
变量名 := 值

//示例:
smallming := "英文名"
```

# 五.声明多个变量和赋值

* 先声明后赋值


```go
func main() {
	var a, b, c int
	a, b, c = 1, 2, 3
	fmt.Println(a, b, c)
}
```

* 声明时赋值

```go
func main() {
	var a, b, c, d = 1, 2, 3, false
	fmt.Println(a, b, c, d)
}
```

* 声明并赋值,推荐方式

```go
func main() {
	var (
		a = 1
		b = true
		c = "测试"
	)
	fmt.Println(a, b, c)
}
```

* 使用短变量给多个变量赋值时,必须要保证至少有个变量是没有声明的

```go
func main() {
	var (
		a = 1
		b = true
		c = "测试"
	)
	//短变量操作多个值时只要保证里面至少有一个新变量
	b, c, d := false, "smallming", 3
	fmt.Println(a, b, c, d)
}
```


# 一.变量作用域

* 变量声明位置决定了变量的可访问范围(哪里能调用到变量)
* Go语言中变量的有效范围如下
  * 函数级别:变量声明在函数内部,只有在函数内部才能访问,称变量为局部变量
  * package 包级别,在当前包下都可以访问.称变量为全局变量.变量声明在函数外面
  * 应用级别,在整个应用下任何包内都可以访问.通过首字母大小写控制

# 二.局部变量

* 局部变量一定是在函数内部
* 在哪个{}内部声明,只能在哪个{}内部访问

```go
func test1() {
	i := 2 //从此处开始到test1结束}任何位置都能调用i
	if i>=2{
		j:=3
		fmt.Println(i+j)//此处可以访问i
	}
	fmt.Println(i)
	//fmt.Println(j)//此处不能调用j,超出声明j时{}外
}
func test2() {
	fmt.Println(i) //此处无法调用test1()中的i
}
```

# 三.全局变量

* 全局变量声明到函数外部,整个包都可以访问
* 如果全局变量首字母大写,跨包也可以访问.
* 声明全局变量时规范是

```go
var (
	变量名
	变量名=值
)
```

* 全局变量代码示例

```go
var (
	name = "xiaopang"
	age  = 17
)

func demo1() {
  	fmt.Println("名字:",name)
}

func demo2() {
  	fmt.Println("年龄:",age)
}
```

