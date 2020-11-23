package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	var t time.Time
	fmt.Println(t)

	t2 := time.Now()
	fmt.Println(t2)

	t3 := time.Now()
	t4 := time.Unix(0, t3.UnixNano())
	fmt.Println(t4.String())
	fmt.Println(t4)

	t5 := time.Date(2020, 11, 23, 7, 8, 9, 0, time.Local)
	fmt.Println(t5)

	t6 := time.Now()
	s := t6.Format("2020-11-23 7:04:05")
	fmt.Println(s)

	s2 := "2022-02-04 22:02:04"
	t7, err := time.Parse("2006-01-02 15:04:05", s2)
	fmt.Println(t7, err)

	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Int63n(10))

	var i, j float64 = 12.3, 9.6
	//向下取整,
	fmt.Println(math.Floor(i)) //输出:12
	//向上取整
	fmt.Println(math.Ceil(i)) //输出:13
	//绝对值
	fmt.Println(math.Abs(i)) //输出:12.3
	//返回值分别整数位和小数位,小数位可能出现误差
	num, decimal := math.Modf(i)
	fmt.Println(num, decimal)
	//返回两个变量中大的值
	fmt.Println(math.Max(i, j)) //输出:12.3
	//返回两个变量中小的值
	fmt.Println(math.Min(i, j)) //输出:9.6
	//x的y次方
	fmt.Println(math.Pow(3, 2)) //输出:输出9
	//四舍五入
	fmt.Println(math.Round(i)) //输出:12

}
