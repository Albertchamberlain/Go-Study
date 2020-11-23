package main

import "fmt"

func main() {
	strings := [3]string{"算", "法"}
	for i := 0; i < len(strings); i++ {
		fmt.Println(strings[i])
	}

	string1 := [3]string{"学", "习"}
	for i, n := range string1 {
		fmt.Println(i, n)
	}

	for i := 0; i < 5; i++ {
		fmt.Println("开始")
		if i == 2 || i == 3 {
			continue
		}
		fmt.Println("结束")
	}

myfor:
	for k := 0; k < 2; k++ {
		for i := 0; i < 3; i++ {
			if i == 1 {
				continue myfor
			}
			fmt.Println(k, i, "结束")
		}
	}

myfor2:
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			if j == 1 {
				break myfor2
			}
			fmt.Println(i, j)
		}
	}

}
