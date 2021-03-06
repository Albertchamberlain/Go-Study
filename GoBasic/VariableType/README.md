# 一. 整型概述

* 在Go语言中可以进行**数学运算**的类型分为整型和浮点型
* 所有的整数数字存储到整型中就可以进行数学运算
  * 整型和整型运算的结果还是整型.(5/2=2)
  * 都是整型,但是int8和int16之间不可以进行运算,必须进行类型转换
* 整型分为有符号整型和无符号整型
  * 有符号整型有正数和负数.其二进制最高位表示符号,0为正数1为负数.int和intx为有符号整型

```
//int8 举例
0000 0010=2
1000 0010=-2
```

* 无符号整型只能取大于等于0的整数.其二进制最高位表示真实数字.unit和unitx为无符号整型

```
//uint8 举例
0000 0010=2
1000 0010=130
```

* 整型取值范围和作用
  * 有符号整数统一公式为:-2的n-1次幂到2的n-1次幂减一
  * 无符号整数统一公式为:0到2的n次幂减一

|     类型 | 取值范围                                                            |
|--------:|:-------------------------------------------------------------------|
|    int8 | [-128 , 127]                                                       |
|   int16 | [-32768 , 32767]                                                   |
|   int32 | [-2147483648 , 2147483647] Go语言中没有字符类型,所有字符都使用int32存储 |
|   int64 | [-9223372036854775808 , 9223372036854775807]                       |
|     int | 受限于计算机系统,系统是多少位,int为多少位                               |
|   uint8 | [0 , 255]                                                          |
|  uint16 | [0 , 65535]                                                        |
|  uint32 | [0 , 4294967295]                                                   |
|  uint64 | [0 , 18446744073709551615]                                         |
|    uint | 受限于计算机系统,系统是多少位,uint为多少位                              |
|    rune | 与int32类似,常用在获取值的Unicode码                                   |
|    byte | 与uint8类似.强调值为原始数据.一个字节占用8个二进制                       |
| uintptr | 大小不确定,类型取决于底层编程                                          |

# 二.类型转换

* Go语言是静态类型语言,并且不具备低精度向高精度自动转换功能,所以不同类型变量之间相互赋值需要进行类型转换.
* 例如:

```go
func main() {
	//声明3个类型变量
	var a int = 1
	var b int32 = 2
	var c int64 = 3
	fmt.Println(a, b, c)

	//把int32转换为int64
	a = int(b)
	fmt.Println(a, b)
	a = 1
	
	//把int64转换成int32
	b = int32(c)
	fmt.Println(b, c)
	b = 2
	
	//把int转换为int64
	c = int64(a)
	fmt.Println(a, c)
	c = 3
}
```

# 三.不同进制整数

* 支持八进制,十进制,十六进制数字创建整型,最后由系统转换为十进制
* 不支持二进制值

```go
func main() {
	//默认表示十进制
	d := 17
	
	//0开头表示八进制
	o := 021
	
	//0x开头表示十六进制
	x := 0x11
	
	//e2表示10的2次方
	e := 11e2
	
	//输出
	fmt.Println(d, o, x, e)
	
	//把变量d中内容转换为二进制
	b := fmt.Sprintf("%b", d)
	fmt.Println(b)
}
```

# 二. 浮点数概述

* 浮点类型用于存储带有小数点的数字
* 一个整数数值可以赋值给浮点类型但是一个整型变量不可以赋值给浮点类型
* 浮点数进行运算的结果是浮点数
* Go语言中浮点类型有两个
  * float32
  * float64

## 一.浮点数取值范围

* float32和float64取值范围

|   类型   |                    取值范围                     |
|:-------:|:----------------------------------------------:|
| float32 |  3.40282346638528859811704183484516925440e+38  |
| float64 | 1.797693134862315708145274237317043567981e+308 |

* 可以通过math或中常量快速获取浮点数的最大值

```go
package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.MaxFloat64)
}
```

## 二.浮点运算

* float32和float64之间不可以相互运算,需要进行类型转换

```go
func main() {
	var a float32 = 1.5
	var b float64 = 3.3         //默认是float64类型
	fmt.Println(a + float32(b)) //float64向float32转换
	fmt.Println(float64(a) + b) //float32向float64转换
}
```

* 建议使用float64,虽然占用空间多,但是float32在累计运算时可能出现误差
* 整型运算和浮点型运算结果类型为本身类型

```go
func main() {
	var a, b int = 3, 2
	var c, d float64 = 3, 2
	fmt.Println(a / b) //结果为int,舍去小数部分(向下取整)
	fmt.Println(c / d) //结果为float64
}
```


# 三. 布尔类型概述

* 布尔类型关键字**bool**
* 布尔类型可取值只有`两个`
  * true :代表真,表示成立,二进制表示时1表示真
  * false:代表假,表示不成立,二进制表示时0表示假
* 布尔类型不能与其他类型相互转换
* 布尔类型占用1个byte
* 布尔类型单独使用较少,多用在判断中


* 创建bool类型变量

```go
func main() {
	var a bool = true
	var b bool = false
	var c = true
	d := false
	fmt.Println(a, b, c, d)
}
```

* 使用unsafe包下的Sizeof()可以查看类型占用字节

```go
func main() {
	a := false
	fmt.Println(unsafe.Sizeof(a))
}
```

* 虽然bool类型占用一个byte,但是bool不能和byte或int8相互转换

```go
func main() {
	var a int8 = 1
	var b byte = 0
	var c bool = false
	fmt.Println(a, b, c)
	a = int8(c) //cannot convert c (type bool) to type int8
	b = byte(c) //cannot convert c (type bool) to type byte
	c = bool(a) //cannot convert a (type int8) to type bool
	c = bool(b) //cannot convert b (type byte) to type bool
	b = byte(a) //可以
}
```

* 布尔类型除了直接赋值true或false以外,还是可以表达式赋值,借助比较运算符、逻辑运算符等

```go
func main() {
	a := 5 > 3
	fmt.Println(a)      //输出:true
	fmt.Printf("%T", a) //输出:bool
}
```


# 字符型概述

* 字符型存放单个字母或单个文字
* Go语言不支持字符类型,在Go语言中所有字符值都转换为对应的编码表中int32值
* Go语言默认使用UTF-8编码


# 字符示例

* 示例

```go
func main() {
	//定义数字
	var i rune=`1 //0x5F20
	fmt.Println(i)

	fmt.Printf("%c\n",i)
	
	//获取转换后的内容
	c:=fmt.Sprintf("%c",i)
	fmt.Println(c)
}
```

* 也可以使用单引号表示一个字符类型,但是本质还是正数

```go
func main() {
	c := '侯'
	fmt.Println(c)      //24352
	fmt.Printf("%T", c) //int32
}
```




