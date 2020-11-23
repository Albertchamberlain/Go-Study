package main

import "fmt"

func main() {
	var a *int
	fmt.Println(a)
	fmt.Println(a == nil)

	a2 := 3
	fmt.Println(&a2)
	a2 = 4
	fmt.Println(&a2)

	b := a2
	b = 5
	fmt.Println(&b, &a) //两个值不相同
	fmt.Println(b, a)   ////输出:5 4

	a3 := 123
	var point *int
	point = &a3
	fmt.Println(point)

	*point = 3
	fmt.Println(*point, a3)

	var a4 *int
	fmt.Println(a4)
	fmt.Print(a == nil)

	a5 := new(int)
	fmt.Println(a5)
	*a5 = 123
	fmt.Print(*a5)
	fmt.Println(a5)
	fmt.Println(&a5)
	fmt.Println(&*a5)

	//var a6 *int //没有开辟内存空间
	//*a6 = 123
	//fmt.Println(*a6)

}
