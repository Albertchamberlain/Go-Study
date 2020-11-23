package main

import "fmt"

func main() {

	const (
		a1 = 1
		b1 = 2
		c1 = true
	)

	const (
		a2 = 1
		b2
		c2
	)
	fmt.Println(a2, b2, c2) //输出:1 1 1

	const a string = "smallming"
	const b = 123
	const c = 3*2 + 5 //不要指定类型
	const d = 1.5     //不要指定类型

	fmt.Printf("%T %T", c, d) //int float

	fmt.Println(c + d) //12.5

	//下面这种方式是错误的 ,因为i是变量
	//i := 3
	//const e = i*2 + 5 //const initializer i * 2 + 5 is not a constant

	//* 无论是否使用iota,一组常量中每个的iota值是固定的,iota按照顺序自增1
	//* 每组iota之间无影响
	const (
		a3 = iota
		b3
		c3
	)
	fmt.Println(a3, b3, c3) //输出: 0 1 2

	const (
		d3 = iota << 1
		e3
		f3
	)
	fmt.Println(d3, e3, f3) //输出:0 2 4

}
