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
