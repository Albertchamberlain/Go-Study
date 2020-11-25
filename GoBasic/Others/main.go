package main

import "fmt"

func main() {
	fmt.Println(f()) //输出:0
}

func f() int {
	i := 0
	defer func() {
		i = i + 2
	}()
	return i
}
