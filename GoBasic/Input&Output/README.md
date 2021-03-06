* 以下为常用转义字符🗡

| verb     |          含义           |
|:---------|:----------------------:|
| %d       |        十进制整数        |
| %x,%X    | 大小写方式显示十六进制整数 |
| %o       |        八进制整数        |
| %b       |        二进制整数        |
| %f,%g,%e |         浮点数          |
| %t       |         布尔值          |
| %c       |          字符           |
| %s       |         字符串          |
| %q       |      带双引号字符串      |
| %v       |       内置格式内容       |
| %T       |          类型           |
| %p       |        内存地址         |
| %%       |          字符%          |
| \n       |          换行           |
| \t       |          缩进           |



# 输入输出


## 在Go语言中有多种输出方式,不同的输出适用场景不同.归纳起来三种,每种还分为3种方式(原内容,原内容+ln,原内容+f)

* PrintXX()
* FprintXX()
* SprintXX()

- FprintXX在Go Web中使用比较多,把内容写到响应流中.
* 函数参数中第一个参数是输出流,后面参数是内容,表示把内容写入到输出流中
* 第一个返回值表示输出内容长度(字节数),第二个返回值表示错误,如果没有错误取值nil
  * Fprintln()输出后会添加换行符,所以长度比内容多1个

```
* FprintXX()支持下面三种方式
  * os.Stdout 表示控制台输出流

func main() {
	fmt.Fprint(os.Stdout, "内容1")//向流中写入内容,多个内容之间没有空格
	fmt.Fprintln(os.Stdout, "内容2")//向流中写入内容后额外写入换行符,多个内容之间空格分割
	fmt.Fprintf(os.Stdout, "%s", "内容3")//根据verb格式向流中写入内容
}
```


- PrintXX支持下面三种方式

```
func main() {
	fmt.Println("内容","内容")//输出内容后换行
	fmt.Print("内容","内容")//输出内容后不换行
	fmt.Printf("verb","内容")//根据verb输出指定格式内容
}
```

- SPrintXX支持下面三种方式

```
func main() {
	fmt.Sprint("内容1", "内容12")
	fmt.Sprintln("内容2")
	fmt.Sprintf("%s", "内容3")
}
```

以Sprintln()举例,和Println()主要的`区别`是:

* Sprintln()把形成结果以字符串返回,并没有打印到控制台
* Println()把结果打印到控制台,返回内容长度和错误

## 输入

* 使用`Scanln(&变量名,&变量名)`的方式接收.
  * 输入的内容必须都在同一行
  * 每个内容之间使用空格分割
  * 回车换行后表示停止输入.
  * 如果希望接收3个值,而在控制台只输入2个值,回车后也停止接收
  * 如果希望接收2个值,而在控制台输入3个,回车后只能接收两个值

```go
package main

import "fmt"

func main() {
	var name, age string //声明两个字符串变量,变量在本章节后面讲解
	fmt.Print("请输入姓名和姓名:")
	fmt.Scanln(&name, &age) //此处&变量名是地址.指针地址在后面章节境界
	fmt.Println("接收到内容为:", name, age)
}
```

* 也可以使用`fmt.Scanf(verb,&变量)`按照特定的格式进行输入.
  * 下面例子演示的每次换行输入一个内容

```go
package main

import "fmt"

func main() {
	var a,b string
	fmt.Scanf("%s\n%s",&a,&b);
	fmt.Printf("%s\n%s",a,b)
}
```

* 需要注意,如果同行输入两个字符串,中间使用空格,否则编译器无法对输入内容拆分

```go
package main

import "fmt"

func main() {
	var a string
	var b string
	//输入时必须输入: aaa bbb
	//如果中间没有空格则把所有内容都赋值给了a
	fmt.Scanf("%s%s",&a,&b)
	fmt.Println(a,b)
}
```

