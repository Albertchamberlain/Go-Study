package main

import "fmt"

func main() {
	var a, b, c int
	a, b, c = 1, 2, 3
	fmt.Println(a, b, c)

	var aa, bb, cc, dd = 1, 2, 3, false //注意顺序
	fmt.Println(aa, bb, cc, dd)

	var (
		a1 = 1
		b1 = true
		c1 = "测试"
	)
	fmt.Println(a1, b1, c1)

	var (
		a2 = 1
		b2 = true
		c2 = "测试"
	)
	//短变量操作多个值时只要保证里面至少有一个新变量
	b2, c2, d2 := false, "smallming", 3
	fmt.Println(a2, b2, c2, d2)

	var a5 bool = true
	var b5 bool = false
	var c5 = true
	d5 := false
	fmt.Println(a5, b5, c5, d5)

}
