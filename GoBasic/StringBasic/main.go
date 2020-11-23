package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	a := "a\tbc"
	b := `a\tbc`
	fmt.Println(a) //输出:a	abc
	fmt.Println(b) //输出a\tabc

	s := "11"
	/*
		第一个参数:需要转换的字符串变量
		第二个参数:这个数字是几进制,常用取值:2,8,10,16
		第三个参数:认为这个数字的整数类型.可取值:0,8,16,32,64.
		但是由于方法最后返回值是int64,所以第三个参数无论设置什么最终结果都是int64
	*/
	i, _ := strconv.ParseInt(s, 10, 8)
	fmt.Println(i)
	fmt.Printf("%T\n", i)

	//简单写法,相当于strconv.ParseInt(s,10,64)
	// auto to int
	i1, _ := strconv.Atoi(s)
	fmt.Println(i1)

	i3 := 11
	/*
		第一个参数:必须是int64类型
		第二个参数:进制数
	*/
	s3 := strconv.FormatInt(int64(i3), 10)
	fmt.Println(s3)        //输出:11
	fmt.Printf("%T\n", s3) //输出:string

	/*
			由于平时常用int,且使用短变量时整数默认是int类型
			所以下面方式较常用,把int转换为string
		int to string
	*/
	s1 := strconv.Itoa(i3)
	fmt.Println(s1)      //输出:11
	fmt.Printf("%T", s1) //输出:string

	f := 1.5
	/*
		把浮点型转换为字符串类型
		第一个参数:浮点型变量
		第二个参数:
			'f'（-ddd.dddd）
			'b'（-ddddp±ddd，指数为二进制）
			'e'（-d.dddde±dd，十进制指数）
			'E'（-d.ddddE±dd，十进制指数）
			'g'（指数很大时用'e'格式，否则'f'格式）
			'G'（指数很大时用'E'格式，否则'f'格式）
		第三个参数:小数点精度,精度不够使用0补全,超出精度四舍五入
		第四个参数:浮点型变量类型,64表示float64,32表示float32
	*/
	s5 := strconv.FormatFloat(f, 'g', 5, 64)

	fmt.Println(s5)

	s4 := "smallming小"
	a4 := s4[0]
	fmt.Println(a4)        //输出:115
	fmt.Printf("%T\n", a4) //输出uint8
	b4 := fmt.Sprintf("%c", a4)
	fmt.Printf("%T\n", b4) //输出:string
	fmt.Println(b4)        //输出s

	s6 := "smallming小"
	fmt.Println(len(s6)) //输出:12,字节长度
	fmt.Println(s6[1:4]) //输出:mal
	fmt.Println(s6[:2])  //输出:sm
	fmt.Println(s6[5:])  //输出:ming小
	// fmt.Println(s6[-1])

	s7 := "smallming小"
	s8 := []rune(s7)
	fmt.Println(len(s8))     //输出:10
	fmt.Println(s8[9])       //输出24352
	fmt.Println("%c", s8[9]) //输出:小

	//遍历字符串中内容
	for i, n := range s8 {
		fmt.Println(i, n)
	}

	s9 := "smallming"
	//第一次出现的索引
	fmt.Println(strings.Index(s9, "l"))
	//最后一次出现的索引
	fmt.Println(strings.LastIndex(s9, "l"))
	//是否以指定内容开头
	fmt.Println(strings.HasPrefix(s9, "small"))
	//是否以指定内容结尾
	fmt.Println(strings.HasSuffix(s9, "ming"))
	//是否包含指定字符串
	fmt.Println(strings.Contains(s9, "mi"))
	//全变小写
	fmt.Println(strings.ToLower(s9))
	//全变大写
	fmt.Println(strings.ToUpper(s9))
	//把字符串中前n个old子字符串替换成new字符串,如果n小于0表示全部替换.
	//如果n大于old个数也表示全部替换
	fmt.Println(strings.Replace(s9, "m", "k", -1))
	//把字符串重复count遍
	fmt.Println(strings.Repeat(s9, 2))
	//去掉字符串前后指定字符
	fmt.Println(strings.Trim(s9, " ")) //去空格可以使用strings.TrimSpace(s)
	//根据指定字符把字符串拆分成切片
	fmt.Println(strings.Split(s9, "m"))
	//使用指定分隔符把切片内容合并成字符串
	arr := []string{"small", "ming"}
	fmt.Println(strings.Join(arr, ""))

}
