package main

import "fmt"

func main() {

	//二维数组的嵌套输出
	var a = [5][2]int{{0, 1}, {1, 2}, {2, 3}, {3, 4}, {4, 5}}
	fmt.Println(len(a))

	fmt.Println(len(a[0]))

	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[0]); j++ {
			fmt.Println(a[i][j])
		}
	}
}
