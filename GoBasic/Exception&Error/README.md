# 一.断言

assert❗
* 只要实现了接口的全部方法认为这个类型属于接口类型,如果编写一个接口,这个接口中没有任何方法,这时认为所有类型都实现了这个接口(顶级).所以Go语言中`interface{}`代表任意类型
* 如果`interface{}`作为方法参数就可以`接收任意类型`,但是在程序中有时有需要知道这个参数到底是什么类型,这个时候就需要使用断言
* 断言使用时使用interface{}变量点括号,括号中判断是否属于的类型

```go
i interface{}
i.(Type)
```

* 断言的两大作用:
  * 判断是否是指定类型
  * 把interface{}转换为特定类型

# 二.代码示例

* 断言可以有一个返回值,如果判断结果是指定类型返回变量值,如果不是指定类型报错

```go
func demo(i interface{}){
	result:=i.(int)
	fmt.Println(result)
}

func main() {
	/*
	参数是456时,程序运行正常,输出:
		456
	参数是false时报错：
		panic: interface conversion: interface {} is bool, not int
	 */
	demo(456)
}
```

* 断言也可以有两个返回值,这时无论是否是指定类型都不报错.
  * 第一个参数:
    * 如果正确:返回值变量值
    * 如果错误:返回判断类型的默认值
  * 第二个参数:
    * 返回值为bool类型,true表示正确,false表示错误

```go
func demo(i interface{}) {
	result, ok := i.(int)
	fmt.Println(result, ok)
}

func main() {
	/*
	参数是456时,程序运行正常,输出:
		456	true
	参数是字符串"abc"时程序运行正常,输出:
		0 false
	 */
	demo("abc")
}
```


# 一. 错误

* 在程序执行过程中出现的不正常情况称为错误
* Go语言中使用builtin包下error接口作为错误类型,官方源码定义如下
  * 只包含了一个方法,方法返回值是string,表示错误信息

```go
// The error built-in interface type is the conventional interface for
// representing an error condition, with the nil value representing no error.
type error interface {
	Error() string
}
```

* Go语言中`错误`都作为方法/函数的`返回值`,因为Go语言认为使用其他语言类似`try...catch`这种方式会影响到程序结构
* 在Go语言标准库的`errors`包中提供了error接口的实现结构体errorString,并重写了error接口的Error()方法.额外还提供了快速创建错误的函数

```go
package errors

// New returns an error that formats as the given text.
func New(text string) error {
	return &errorString{text}
}

// errorString is a trivial implementation of error.
type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

```

* 如果错误信息由很多变量(小块)组成,可以借助`fmt.Errorf("verb",...)`完成错误信息格式化,因为底层还是errors.New()

```go
// Errorf formats according to a format specifier and returns the string
// as a value that satisfies error.
func Errorf(format string, a ...interface{}) error {
	return errors.New(Sprintf(format, a...))
}
```

# 二.自定义错误

* 使用Go语言标准库创建错误,并返回

```go
func demo(i, k int) (d int, e error) {
	if k == 0 {
		e = errors.New("初始不能为0")
		d=0
		return
	}
	d = i / k
	return
}

func main() {
	result,error:=demo(6,0)
	fmt.Println(result,error)
}
```

* 如果错误信息由多个内容组成,可以使用下面实现方式

```go
func demo(i, k int) (d int, e error) {
	if k == 0 {
		e = fmt.Errorf("%s%d和%d", "除数不能是0,两个参数分别是:", i, k)
		d = 0
		return
	}
	d = i / k
	return
}

func main() {
	result, error := demo(6, 0)
	fmt.Println(result, error)
}
```

# 三.Go语言中错误处理方式

* 可以忽略错误信息(不处理),使用占位符

```go
func demo(i, k int) (d int, e error) {
	if k == 0 {
		e = fmt.Errorf("%s%d和%d", "除数不能是0,两个参数分别是:", i, k)
		d = 0
		return
	}
	d = i / k
	return
}

func main() {
	result, _ := demo(6, 0)
	fmt.Println(result)
}
```

* 使用if处理错误,原则上每个错误都`应该`解决

```go
func demo(i, k int) (d int, e error) {
	if k == 0 {
		e = fmt.Errorf("%s%d和%d", "除数不能是0,两个参数分别是:", i, k)
		d = 0
		return
	}
	d = i / k
	return
}

func main() {
	result, error := demo(6, 0)
	if error != nil {
		fmt.Println("发生错误", error)
		return
	}
	fmt.Println("程序执行成功,结果为:", result)
}
```