package main

import "fmt"

func main() {
	var slice []string
	// var array [5] string

	fmt.Println(slice == nil)
	fmt.Println("%p", slice)

	names := []string{"algorithm", "算法"}
	fmt.Println(names)

	namess := []string{"algorithm", "算法"}
	names1 := namess
	names1[0] = "法"
	fmt.Println(namess, names1)
	fmt.Printf("%p %p", namess, names1) //地址相同
}
