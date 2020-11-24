package main

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
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

	num1 := []int{1, 7, 5, 2, 6}
	sort.Ints(num1)

	fmt.Println(num1)

	sort.Sort(sort.Reverse(sort.IntSlice(num1))) //降序
	fmt.Println(num1)

	f := []float64{1.5, 7.2, 5.8, 2.3, 6.9}
	sort.Float64s(f) //升序
	fmt.Println(f)
	sort.Sort(sort.Reverse(sort.Float64Slice(f))) //降序
	fmt.Println(f)

	s12 := []string{"算", "法是灵魂", "a", "d", "程序", "的", "灵魂a"}
	sort.Sort(sort.StringSlice(s12)) //升序

	fmt.Println(sort.StringSlice(s12))
	fmt.Println(s12)
	//查找内容的索引,如果不存在,返回内容应该在升序排序切片的哪个位置插入
	fmt.Println(sort.SearchStrings(s12, "没错"))
	sort.Sort(sort.Reverse(sort.StringSlice(s12)))
	fmt.Println(s12)

	res := test1()
	fmt.Println(res()) //输出2
	fmt.Println(res()) //输出3
	fmt.Println(res()) //输出4

	f1 := test2()
	fmt.Println("f的地址", f1) //输出匿名函数地址
	fmt.Println("f:", f1()) //调用匿名函数输出2
	fmt.Println("f:", f1()) //调用匿名函数输出3
	k := test1()
	fmt.Println("k的地址", k)  //输出匿名函数地址,与f相等
	fmt.Println("k:", k())  //调用匿名函数输出2
	fmt.Println("f:", f1()) //输出:4
	fmt.Println("k:", k())  //输出:3

}

func test1() func() int {
	i := 1
	return func() int {
		i = i + 1
		return i
	}
}

func test2() func() int {
	i := 1
	return func() int {
		i++
		// 每调用一次test1()输出的地址不一样
		fmt.Println("i的地址:", &i)
		return i
	}
}
