package main

import (
	"errors"
	"fmt"
)

func demo(i interface{}) {
	result := i.(int)
	fmt.Println(result)
}
func demo3(i, k int) (d int, e error) {
	if k == 0 {
		e = errors.New("初始不能为0")
		d = 0
		return
	}
	d = i / k
	return
}

func demo4(i, k int) (d int, e error) {
	if k == 0 {
		e = fmt.Errorf("%s%d和%d", "除数不能是0,两个参数分别是:", i, k)
		d = 0
		return
	}
	d = i / k
	return
}
func demo2(i interface{}) {
	result, ok := i.(int)
	fmt.Println(result, ok)
}

func demo5(i, k int) (d int, e error) {
	if k == 0 {
		e = fmt.Errorf("%s%d和%d", "除数不能是0,两个参数分别是:", i, k)
		d = 0
		return
	}
	d = i / k
	return
}
func demo6(i, k int) (d int, e error) {
	if k == 0 {
		e = fmt.Errorf("%s%d和%d", "除数不能是0,两个参数分别是:", i, k)
		d = 0
		return
	}
	d = i / k
	return
}
func main() {

	result6, error6 := demo6(6, 0)
	if error6 != nil {
		fmt.Println("发生错误", error6)
		return
	}
	fmt.Println("程序执行成功,结果为:", result6)

	demo(345)
	demo2("qwe")

	result, error := demo3(6, 0)
	fmt.Println(result, error)

	result1, error1 := demo4(6, 0)
	fmt.Println(result1, error1)

	result2, _ := demo5(6, 0)
	fmt.Println(result2)

}
