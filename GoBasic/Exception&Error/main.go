package main

import "fmt"

func demo(i interface{}) {
	result := i.(int)
	fmt.Println(result)
}

func demo2(i interface{}) {
	result, ok := i.(int)
	fmt.Println(result, ok)
}
func main() {
	demo(345)
	demo2("qwe")
}
