package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//声明3个类型变量
	var a int = 1
	var b int32 = 2
	var c int64 = 3
	fmt.Println(a, b, c)

	//把int32转换为int64
	a = int(b)
	fmt.Println(a, b)
	a = 1

	//把int64转换成int32
	b = int32(c)
	fmt.Println(b, c)
	b = 2

	//把int转换为int64
	c = int64(a)
	fmt.Println(a, c)
	c = 3
	//默认表示十进制
	d := 17

	//0开头表示八进制
	o := 021

	//0x开头表示十六进制
	x := 0x11

	//e2表示10的2次方
	e := 11e2

	//输出
	fmt.Println(d, o, x, e)

	//把变量d中内容转换为二进制
	bb := fmt.Sprintf("%b", d)
	fmt.Println(bb)

	//定义数字
	var i rune = 1 //0x5F20
	fmt.Println(i)

	//输出汉字张
	fmt.Printf("%c\n", i)

	//获取转换后的内容
	c2 := fmt.Sprintf("%c", i)
	fmt.Println(c2)

	c3 := '算'
	fmt.Println(c)
	fmt.Printf("%T", c3) //int32

	var a3 float32 = 1.5
	var b3 float64 = 3.3          //默认是float64类型
	fmt.Println(a3 + float32(b3)) //float64向float32转换
	fmt.Println(float64(a3) + b3) //float32向float64转换

	var a4, b4 int = 3, 2
	var c4, d4 float64 = 3, 2
	fmt.Println(a4 / b4) //结果为int,舍去小数部分(向下取整)
	fmt.Println(c4 / d4) //结果为float64

	a6 := false
	fmt.Println(unsafe.Sizeof(a6))

	var a7 int8 = 1
	var b7 byte = 0
	var c7 bool = false
	fmt.Println(a7, b7, c7)
	/*
		以下不能转换
	*/
	//a = int8(c7) //cannot convert c (type bool) to type int8
	//b = byte(c7) //cannot convert c (type bool) to type byte
	//c = bool(a7) //cannot convert a (type int8) to type bool
	//c = bool(b7) //cannot convert b (type byte) to type bool

	b7 = byte(a7) //可以

}
