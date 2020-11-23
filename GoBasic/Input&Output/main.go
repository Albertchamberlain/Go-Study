package main

import "fmt"

func main() {
	fmt.Printf("%s", "内容")  //输出
	fmt.Scanf("%s", "接收变量") //输入

	var name, age string //声明两个字符串变量,变量在本章节后面讲解
	fmt.Print("请输入姓名和姓名:")
	fmt.Scanln(&name, &age) //此处&变量名是地址.指针地址在后面章节境界
	fmt.Println("接收到内容为:", name, age)

	var a, b string
	fmt.Scanf("%s\n%s", &a, &b)
	fmt.Printf("%s\n%s", a, b)

	var aa string
	var bb string
	//输入时必须输入: aaa bbb
	//如果中间没有空格则把所有内容都赋值给了a
	fmt.Scanf("%s%s", &aa, &bb)
	fmt.Println(a, b)
}
