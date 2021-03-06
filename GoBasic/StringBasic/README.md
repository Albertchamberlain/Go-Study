# 一.字符串概述

* 字符串是一段不可变的字符序列.内容是任意内容,可以是一段文字也可以是一串数字,但是字符串类型数字不能进行数学运算,必须转换成整型或浮点型
* 字符串类型关键字:string
* 创建字符串类型变量

```go
var s string = "smallming"
s1 := "smallming"
```

* 字符串类型的值使用双引号""扩上,内容支持转义字符串.两侧使用反单引号时原格式输出

```go
func main() {
	a := "a\tbc"
	b := `a\tbc`
	fmt.Println(a) //输出:a	abc
	fmt.Println(b) //输出a\tabc
}
```


# 二.字符串和数值转换

* 包strconv提供了字符串和其他类型相互转换的函数,下面以字符串和数值类型转换为例
* int和string相互转换

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
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
	i1, _ := strconv.Atoi(s)
	fmt.Println(i1)
}
```

* Int转换成string

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 11
	/*
	第一个参数:必须是int64类型
	第二个参数:进制数
	 */
	s := strconv.FormatInt(int64(i), 10)
	fmt.Println(s)        //输出:11
	fmt.Printf("%T\n", s) //输出:string

	/*
	由于平时常用int,且使用短变量时整数默认是int类型
	所以下面方式较常用,把int转换为string
	 */
	s1 := strconv.Itoa(i)
	fmt.Println(s1)      //输出:11
	fmt.Printf("%T", s1) //输出:string
}
```

* string转换为floatXX类型

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "1.5"
	/*
	把字符串转换为指定类型
	第一个参数:字符串
	第二个参数:可取值为32和64,分别表示float32和float64
	返回值是float64
	 */
	f, _ := strconv.ParseFloat(s, 64)
	fmt.Println(f)
	fmt.Printf("%T", f)
}
```

* floatXX转换为string类型

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
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
	s := strconv.FormatFloat(f, 'g', 5, 64)

	fmt.Println(s)
}

```

# 三,字符串截取

## 一.字符串截取

* 可以使用**len(字符串变量)**获取字符串的`字节`长度,其中英文占1个字节长度,中文占用3个字节长度
* 可以使用**变量名[n]**获取到字符串第n+1个字节,返回这个字节对应的Unicode码值(uint8类型).注意n的取值范围是[0,长度)

```go
func main() {
	s := "smallming小"
	a := s[0]
	fmt.Println(a)        //输出:115
	fmt.Printf("%T\n", a) //输出uint8
	b := fmt.Sprintf("%c", a)
	fmt.Printf("%T\n", b) //输出:string
	fmt.Println(b)        //输出s
}
```

* 可以使用变量名[n:m]取出`大于等于n小于m`的字符序列
  * n和m都可以省略,省略时认为n为0,m为长度
  * 因为中文占用三个字节,如果没有把中文完整取出,会出现乱码

```go
func main() {
	s := "smallming小"
	fmt.Println(len(s)) //输出:12,字节长度
	fmt.Println(s[1:4]) //输出:mal
	fmt.Println(s[:2])  //输出:sm
	fmt.Println(s[5:])  //输出:ming小
}
```

* 可以通过把字符串转换为切片获取`长度`,并获取里面内容.
  也可以直接使用for循环结合range获取

```go
func main() {
	s := "smallming小"
	s1 := []rune(s)
	fmt.Println(len(s1))    //输出:10
	fmt.Println(s1[9])      //输出24352
	fmt.Printf("%c", s1[9]) //输出:小

	//遍历字符串中内容
	for i, n := range s {
		fmt.Println(i, n)
	}
}
```

## 二.常用函数

* 在strings包中提供了字符串常用的函数
* 常用函数整理如下

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "smallming"
	//第一次出现的索引
	fmt.Println(strings.Index(s, "l"))
	//最后一次出现的索引
	fmt.Println(strings.LastIndex(s, "l"))
	//是否以指定内容开头
	fmt.Println(strings.HasPrefix(s, "small"))
	//是否以指定内容结尾
	fmt.Println(strings.HasSuffix(s, "ming"))
	//是否包含指定字符串
	fmt.Println(strings.Contains(s, "mi"))
	//全变小写
	fmt.Println(strings.ToLower(s))
	//全大写
	fmt.Println(strings.ToUpper(s))
	//把字符串中前n个old子字符串替换成new字符串,如果n小于0表示全部替换.
	//如果n大于old个数也表示全部替换
	fmt.Println(strings.Replace(s, "m", "k", -1))
	//把字符串重复count遍
	fmt.Println(strings.Repeat(s, 2))
	//去掉字符串前后指定字符
	fmt.Println(strings.Trim(s, " ")) //去空格可以使用strings.TrimSpace(s)
	//根据指定字符把字符串拆分成切片
	fmt.Println(strings.Split(s, "m"))
	//使用指定分隔符把切片内容合并成字符串
	arr := []string{"small", "ming"}
	fmt.Println(strings.Join(arr, ""))
}
```

