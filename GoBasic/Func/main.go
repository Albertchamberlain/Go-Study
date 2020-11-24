package main

import "fmt"

func main() {
	demo()
	i := 5
	s := "xiaopang"
	show(s, i)

	a, b := 1, 2
	fmt.Println(add(a, b))

	a1, b1 := 1, 2
	fmt.Println(add2(a1, b1))

	demo5()

	//每个返回值都接收
	a2, b2 := demo5()
	fmt.Println(a2, b2)

	//不希望接收的返回值使用下划线占位
	c, _ := demo5()
	fmt.Println(c)

	demo7("看书", "写代码", "看小胖")
	demo8("张三", "看书", "写代码", "看小胖")

	func() {
		fmt.Println("这是匿名函数")
	}() //括号表示调用

	func(s string) {
		fmt.Println(s, "这是匿名函数")
	}("传递参数") //调用时传递参数

	r := func(s string) int {
		fmt.Println(s, "这是匿名函数")
		return 110
	}("传递参数") //调用时传递参数
	fmt.Println(r)

	var aa1 func()
	aa1 = func() {
		fmt.Println("执行函数")
	} //注意此处没有括号,有括号表示调用函数,变量a就表示接收函数返回值
	aa1() //调用函数

	bb1 := func(s string) {
		fmt.Println("执行第二个函数")
	}
	bb1("参数")

	var a12 func()           //无参数无返回值
	var b12 func(int)        //有一个int类型参数
	var c15 func(int) string //有一个int类型参数和string类型返回值
	fmt.Println(a12, b12, c15)
	//使用定义好的函数
	d2 := c15
	d2(14)
	//函数名称c也是一个变量
	//c()

	var a17 func()

	a17 = b17
	a17()
	var c17 func()
	c17 = a17
	c17()
	fmt.Printf("%p %p", a17, c17) //输出地址相同
	a18(func(s string) {
		fmt.Println(s)
	})

	//此时result指向返回值函数.
	result := a19()
	//调用函数,才能获取结果
	fmt.Println(result())
}

func a19() func() int {
	return func() int {
		return 110
	}
}

func a18(b18 func(s string)) {
	fmt.Println("a执行")
	b18("传递给s的内容")
}

func b17() {
	fmt.Println("b")
}

func demo() {
	fmt.Println("okk")
}

func show(name string, age int) {
	fmt.Println("姓名:", name, "年龄", age)
}

func add(c, d int) int {
	return c + d
}

func add2(c, d int) (sum int) {
	sum = c + d
	return
}

func demo5() (string, int) {
	return "xiaopang", 17
}

func demo6() (name string, age int) {
	name = "smallming"
	age = 17
	return
}

func demo7(hover ...string) {
	for a, b := range hover {
		fmt.Println(a, b)
	}
}
func demo8(name string, hover ...string) {
	fmt.Println(name, "的爱好是")
	for a, b := range hover {
		fmt.Println(a, b)
	}
}
