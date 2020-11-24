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
