package main

import "fmt"

func main() {
	score := 65
	if score >= 60 {
		fmt.Println("That's fine")
	}
	if score < 60 {
		fmt.Println("unlucky")
	}

	score2 := 77
	if score2 >= 60 {
		if score2 >= 60 && score < 70 {
			fmt.Println("C")
		}
		if score2 >= 70 && score2 < 80 {
			fmt.Println("B")
		}
		if score2 >= 80 && score2 < 90 {
			fmt.Println("A")
		}
		if score2 >= 90 {
			fmt.Println("A+")
		}
	} else {
		fmt.Println("D")
	}

	score3 := 77
	if score3 >= 90 {
		fmt.Println("优秀")
	} else if score3 >= 80 {
		fmt.Println("良好")
	} else if score3 >= 70 {
		fmt.Println("中等")
	} else if score3 >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}

	switch num := 16; num {
	case 2:
		fmt.Println("Binary")
	case 8:
		fmt.Println("Octal")
	case 10:
		fmt.Println("Decimal")
	case 16:
		fmt.Println("Hexadecimal")
	default:
		fmt.Println("Wrong")
	}
	fmt.Println("Stop")

	month := 5
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		fmt.Println("31")
	case 2:
		fmt.Println(28 / 29)
	default:
		fmt.Println("30")
	}

	switch num1 := 1; num1 {
	case 1:
		fmt.Println("1")
		fallthrough
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
		fallthrough
	case 4:
		fmt.Println(4)
	default:
		fmt.Println("not match")
	}
	fmt.Println("Stop")

	switch num := 1; num {
	case 1:
		fmt.Println("1")
		break
		fmt.Println("break后面代码都不执行")
		fallthrough
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
		fallthrough
	case 4:
		fmt.Println("4")
	default:
		fmt.Println("不是1,2,3,4")
	}
	fmt.Println("程序结束")

	fmt.Println("执行程序")
	i := 6
	if i == 6 {
		goto Loop
	}
	fmt.Println("if下面输出")
Loop:
	fmt.Println("loop")

	fmt.Println("执行程序")
	i2 := 6
	if i2 == 6 {
		goto Loop2
		goto Loop1
	}
	fmt.Println("if下面输出")
Loop2:
	fmt.Println("loop")
Loop1: //报错:label Loop1 defined and not used
	fmt.Println("Loop1")

}
