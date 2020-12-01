# 一.defer使用
功能类似于`finally`
* Go语言中defer可以完成延迟功能,当前函数执行完成后执行defer功能
* defer最常用的就是关闭连接(数据库连接,文件等)可以打开连接后代码紧跟defer进行关闭,后面在执行其他功能
  * 在很多语言中要求必须按照顺序执行,也就是必须把关闭代码写在最后,但是经常会忘记关闭导致内存溢出,而Golang中defer很好的解决了这个问题.无论defer写到哪里都是最后执行

```go
func main() {
   fmt.Println("打开连接")
   defer func(){
      fmt.Println("关闭连接")
   }()
   fmt.Println("进行操作")
   //输出:打开连接 进行操作 关闭连接
}
```

# 二.多个defer

* 多重defer采用`栈`结构执行,先产生后执行
* 在很多代码结构中都可能出现产生多个对象,而程序希望这些对象倒序关闭,多个defer正好可以解决这个问题

```go
func main() {
   fmt.Println("打开连接A")
   defer func(){
      fmt.Println("关闭连接A")
   }()
   fmt.Println("打开连接B")
   defer func(){
      fmt.Println("关闭连接B")
   }()
   fmt.Println("进行操作")
   //输出:打开连接A 打开连接B 进行操作 关闭连接B 关闭连接A
}
```

# 三.defer和return结合

* defer与return同时存在时,要把return理解成两条执行结合(不是原子指令),一个指令是给返回值`赋值`,另一个指令返回`跳出函数`

* defer和return时整体执行顺序
  * 先给返回值赋值
  * 执行defer
  * 返回跳出函数

* 没有定义返回值接收变量,执行defer时返回值已经赋值

```go
func f() int{
	i:=0
	defer func(){
		i=i+2
	}()
	return i
}

func main() {
	fmt.Println(f())//输出:0
}
```

* 声明接收返回值变量,执行defer时修改了返回值内容.
  * 由于return后面没有内容,就无法给返回值赋值,所以执行defer时返回值才有内容

```go
func f() (i int){
	defer func(){
		i=i+2
	}()
	return
}
func main() {
	fmt.Println(f())//输出:2
}
```


# 一.日志简介


* 有三种级别日志输出
  * Print() 输出日志信息
  * Panic()  打印日志信息,并触发panic,日志信息为Panic信息
  * Fatal()  打印日志信息后调用os.Exit(1)
* 所有日志信息打印时都带有时间,且颜色为红色
* 每种级别日志打印都提供了三个函数
  * Println()
  * Print()
  * Printf()
* 日志文件扩展名为log

# 二.普通日志信息打印

* 官方源码如下

```go
func Println(v ...interface{}) {
	std.Output(2, fmt.Sprintln(v...))
}
```

* 直接使用log包调用Println()即可

```go
log.Println("打印日志信息")
```

# 三.Panic日志信息打印

* 通过源码可以看出在日志信息打印后调用了panic()函数,且日志信息为panic信息

```go
// Panicln is equivalent to Println() followed by a call to panic().
func Panicln(v ...interface{}) {
	s := fmt.Sprintln(v...)
	std.Output(2, s)
	panic(s)
}
```

* 执行后输出日志信息,同时也会触发panic

```go
log.Panicln("打印日志信息")
```

# 四.致命日志信息

* 打印日志后,终止程序

```go
// Fatal is equivalent to Print() followed by a call to os.Exit(1).
func Fatal(v ...interface{}) {
	std.Output(2, fmt.Sprint(v...))
	os.Exit(1)
}
```

* 执行日志打印后,程序被终止

```go
log.Fatal("打印日志信息")
```

# 五.打印日志信息到文件中

* Go语言标准库支持输出日志信息到文件中.
* 输出日志时的几种状态

```go
const (
	Ldate         = 1 << iota     // the date in the local time zone: 2009/01/23
	Ltime                         // the time in the local time zone: 01:23:23
	Lmicroseconds                 // microsecond resolution: 01:23:23.123123.  assumes Ltime.
	Llongfile                     // full file name and line number: /a/b/c/d.go:23
	Lshortfile                    // final file name element and line number: d.go:23. overrides Llongfile
	LUTC                          // if Ldate or Ltime is set, use UTC rather than the local time zone
	LstdFlags     = Ldate | Ltime // initial values for the standard logger
)
```

* 代码如下

```go
	f, _ := os.OpenFile("D:/golog.log", os.O_APPEND|os.O_CREATE, 07777)
	defer f.Close()
	logger := log.New(f, "[info]\t", log.Ltime)
	logger.Println("输出日志信息")
```


# 一.os包结构介绍

* Go语言标准库中os包提供了`不依赖平台`的操作系统接口
* 设计为Unix风格的，而错误处理是go风格的,失败的调用会返回`错误值`而非错误码。通常错误值里包含更多信息
* os包及子包功能

```
-- os 包
  --os/exec 包,负责执行外部命令.
  --os/signal对输入信息的访问
  --os/user 通过名称或ID	查询用户账户
```

* 在os/user中提供了User结构体,表示操作系统用户
  * Uid 用户id
  * Gid 所属组id
  * Username 用户名
  * Name 所属组名
  * HomeDir 用户对应文件夹路径

* 在os/user中的Group表示用户所属组
  * Gid 组的id
  * Name 组的名称

```go
// Group represents a grouping of users.
//
// On POSIX systems Gid contains a decimal number representing the group ID.
type Group struct {
	Gid  string // group ID
	Name string // group name
}
```

* 整个os/user包中内容比较少,提供了两个错误类型和获取当前用户,查找用户

```go
type UnknownUserError
  func (e UnknownUserError) Error() string
type UnknownUserIdError
  func (e UnknownUserIdError) Error() string
type User
  func Current() (*User, error)
  func Lookup(username string) (*User, error)
  func LookupId(uid string) (*User, error)
```


## 代码示例

* 可以获取当前用户或查找用户后获取用户信息

```go
   //获取当前登录用户
   //u,_:=user.Current()
   /*
   Lookup()参数是用户名,按照用户名查找指定用户对象
   注意:必须使用完整名称不可以只写zhang
    */
   u, _ := user.Lookup(``)
   fmt.Println(u.Name)
   fmt.Println(u.Gid)
   fmt.Println(u.HomeDir)
   fmt.Println(u.Uid)
   fmt.Println(u.Username)
```

# 一. os文件相关内容介绍

* 使用os包中内容进行操作系统`文件或目录`
* File结构体表示`操作系统文件(夹)`

```go
// File represents an open file descriptor.
type File struct {
	*file // os specific
}
```

```go
// file is the real representation of *File.
// The extra level of indirection ensures that no clients of os
// can overwrite this data, which could cause the finalizer
// to close the wrong file descriptor.
type file struct {
	pfd     poll.FD
	name    string
	dirinfo *dirInfo // nil unless directory being read
}
```

* 操作系统的文件都是有权限控制的,包含可读,可写等,在os包中FileMode表示文件权限,本质是uint32,可取值都以常量形式提供

```go
// A FileMode represents a file's mode and permission bits.
// The bits have the same definition on all systems, so that
// information about files can be moved from one system
// to another portably. Not all bits apply to all systems.
// The only required bit is ModeDir for directories.
type FileMode uint32
```

```go
// The defined file mode bits are the most significant bits of the FileMode.
// The nine least-significant bits are the standard Unix rwxrwxrwx permissions.
// The values of these bits should be considered part of the public API and
// may be used in wire protocols or disk representations: they must not be
// changed, although new bits might be added.
const (
	// The single letters are the abbreviations
	// used by the String method's formatting.
	ModeDir        FileMode = 1 << (32 - 1 - iota) // d: is a directory
	ModeAppend                                     // a: append-only
	ModeExclusive                                  // l: exclusive use
	ModeTemporary                                  // T: temporary file; Plan 9 only
	ModeSymlink                                    // L: symbolic link
	ModeDevice                                     // D: device file
	ModeNamedPipe                                  // p: named pipe (FIFO)
	ModeSocket                                     // S: Unix domain socket
	ModeSetuid                                     // u: setuid
	ModeSetgid                                     // g: setgid
	ModeCharDevice                          // c: Unix character device, when ModeDevice is set
	ModeSticky                                     // t: sticky

	// Mask for the type bits. For regular files, none will be set.
	ModeType = ModeDir | ModeSymlink | ModeNamedPipe | ModeSocket | ModeDevice

	ModePerm FileMode = 0777 // Unix permission bits
)
```

* FIleInfo是一个interface表示文件的信息

```go
// A FileInfo describes a file and is returned by Stat and Lstat.
type FileInfo interface {
	Name() string       // base name of the file
	Size() int64        // length in bytes for regular files; system-dependent for others
	Mode() FileMode     // file mode bits
	ModTime() time.Time // modification time
	IsDir() bool        // abbreviation for Mode().IsDir()
	Sys() interface{}   // underlying data source (can return nil)
}
```

## 资源路径

* 在获取系统资源时资源路径分为`相对路径`和`绝对路径`
* 相对路径:在Go语言中相对路径用于是GOPATH,也就是项目的根目录
* 绝对路径:磁盘根目录开始表示资源详细路径的描述

## 代码示例

* Go语言标准库中提供了两种创建文件夹的方式


```go
	/*
	要求文件夹不存在且父目录必须存在,才能创建
	 */
	//error := os.Mkdir("D:/godir", os.ModeDir)
	//if error != nil {
	//	fmt.Println("文件夹创建失败",error)
	//	return
	//}
	//fmt.Println("文件夹创建成功")


	/*
	如果文件夹已经存在,不报错,保留原文件夹
	如果父目录不存在帮助创建
	 */
	error := os.MkdirAll("D:/godir/a/b", os.ModeDir)
	if error != nil {
		fmt.Println("文件夹创建失败",error)
		return
	}
	fmt.Println("文件夹创建成功")
```

* 创建空文件

```go
	/*
	创建文件时要求文件目录必须已经存在
	如果文件已经存在则会创建一个空文件覆盖之前的文件
	 */
	file, err := os.Create("D:/godir/test.txt")
	if err != nil {
		fmt.Println("文件创建失败,", err)
		return
	}
	fmt.Println("文件创建成功",file.Name())
```

* 重命名文件或文件夹

```go
	/*
	第一个参数:原文件夹名称,要求此路径是必须存在的
	第二个参数:新文件夹名称
	 */
	err := os.Rename("D:/godir", "D:/godir1")
	if err != nil {
		fmt.Println("重命名文件夹失败,", err)
		return
	}
	fmt.Println("文件夹重命名成功")

	/*
	重命名文件和重命名文件夹用法相同
	 */
	err = os.Rename("D:/godir1/test.txt", "D:/godir1/test1.txt")
	if err != nil {
		fmt.Println("重命名文件失败,", err)
		return
	}
	fmt.Println("文件重命名成功")
```

* 获取文件(夹)信息

```go
	f, err := os.Open("D:/godir1/test1.txt")
	defer f.Close() //文件打开后要关闭,释放资源
	if err != nil {
		fmt.Println("打开文件失败", err)
		return
	}
	fileInfo, err := f.Stat()
	if err != nil {
		fmt.Println("获取文件信息失败", err)
		return
	}
	fmt.Println(fileInfo.Name())    //文件名
	fmt.Println(fileInfo.IsDir())   //是否是文件夹,返回bool,true表示文件夹,false表示文件
	fmt.Println(fileInfo.Mode())    //文件权限
	fmt.Println(fileInfo.ModTime()) //修改时间
	fmt.Println(fileInfo.Size())    //文件大小
```

* 删除文件或文件夹

```go
	/*
	删除的内容只能是一个文件或空文件夹且必须存在
	 */
	//err := os.Remove("D:/godir1/a")
	//if err != nil {
	//	fmt.Println("文件删除失败", err)
	//	return
	//}
	//fmt.Println("删除成功")

	/*
	只要文件夹存在,删除文件夹.
	无论文件夹是否有内容都会删除
	如果删除目标是文件,则删除文件
	 */
	err := os.RemoveAll("D:/godir1/a.txt")
	if err != nil {
		fmt.Println("删除失败", err)
		return
	}
	fmt.Println("删除成功")
```

# 一. panic

* panic是`builtin`中函数

```go

func panic(v interface{})
```

* panic有点类似与其他编程语言的`throw`,抛出异常.当执行到panic后终止剩余代码执行.并打印错误栈信息

```go
func main() {
   fmt.Println("1")
   panic("panic执行了,哈哈")
   fmt.Println("2")
}
```

* 执行结果

```
1
panic: panic执行了,哈哈

goroutine 1 [running]:
main.main()
	D:/GoPro/GoBasic/Others/main.go:10 +0x14c
```

* 注意panic不是立即停止程序(`os.Exit(0)`),defer还是执行的.


```go
func main() {
   defer func(){
      fmt.Println("defer执行")
   }()
   fmt.Println("1")
   panic("panic执行了,哈哈")
   fmt.Println("2")
}
```

# 一.recover

* recover()表示恢复程序的panic(),让程序正常运行
* recover()是和panic(v)一样都是builtin中函数,可以接收panic的信息,恢复程序的正常运行

```go

func recover() interface{}
```

* recover()一般用在`defer内部`,如果没有`panic`信息返回nil,如果有panic,recover会把panic状态`取消`

```go
func main() {
	defer func() {
		if error:=recover();error!=nil{
			fmt.Println("出现了panic,使用reover获取信息:",error)
		}
	}()
	fmt.Println("11111111111")
	panic("出现panic")
	fmt.Println("22222222222")
}
```

* 输出

```
11111111111
出现了panic,使用reover获取信息: 出现panic
```

# 二.函数调用过程中panic和recover()

* recover()只能恢复`当前函数级`或`当前函数调用函数`中的panic(),恢复后调用当前级别函数结束,但是调用此函数的函数可以继续执行.
* panic会一直`向上传递`,如果没有recover()则表示终止程序,但是碰见了recover(),recover()所在级别函数表示没有panic,panic就不会向上传递

```go
func demo1(){
	fmt.Println("demo1上半部分")
	demo2()
	fmt.Println("demo1下半部分")
}
func demo2(){
	defer func() {
		recover()//此处进行恢复
	}()
	fmt.Println("demo2上半部分")
	demo3()
	fmt.Println("demo2下半部分")
}
func demo3(){
	fmt.Println("demo3上半部分")
	panic("在demo3出现了panic")
	fmt.Println("demo3下半部分")
}
func main() {
	fmt.Println("程序开始")
	demo1()
	fmt.Println("程序结束")
}
```


# 一.反射介绍
在Java语言中很多框架的实现机制都与反射有关

* 在Go语言标准库中reflect包提供了运行时反射,程序运行过程中`动态操作`结构体
* 当变量存储结构体属性名称,想要对结构体这个属性赋值或查看时,就可以使用反射.
* 反射还可以用作判断变量类型
* 整个reflect包中最重要的两个类型
  * reflect.Type 类型
  * reflect.Value 值
* 获取到Type和Value的函数
  * reflect.TypeOf(interface{}) 返回Type
  * reflect.ValueOf(interface{}) 返回值Value


# 二.代码示例

* 判断变量类型

```go
   a:=1.5
   fmt.Println(reflect.TypeOf(a))
```

* 获取结构体属性的值

```go
type People struct {
   Id   int
   Name string
}

func main() {
   fmt.Println("asdf")

   peo := People{1, "张三"}

   //获取peo的值
   v := reflect.ValueOf(peo)
   //获取属性个数,如果v不是结构体类型panic
   fmt.Println(v.NumField())

   //获取第0个属性,id,并转换为int64类型
   fmt.Println(v.Field(0).Int())
   //获取第1个属性,转换换为string类型
   fmt.Println(v.Field(1).String())

   //根据名字获取类型,并把类型名称转换为string类型
   idValue := v.FieldByName("Id")
   fmt.Println(idValue.Kind().String())

}
```

* 设置结构体属性的值时要传递结构体指针,否者无法获取设置的结构体对象
  * 反射直射结构体属性时,要求属性名首字母必须大写,否则无法设置

```go
package main

import (
   "fmt"
   "reflect"
)

type People struct {
   Id   int
   Name string
}

func main() {
   fmt.Println("asdf")
   peo := People{1, "张三"}

   /*
   反射时获取peo的地址.
   Elem()获取指针指向地址的封装.
   地址的值必须调用Elem()才可以继续操作
    */
   v := reflect.ValueOf(&peo).Elem()

   fmt.Println(v.FieldByName("Id").CanSet())
   v.FieldByName("Id").SetInt(123)
   v.FieldByName("Name").SetString("李四")
   fmt.Println(peo)
}
```

* 结构体支持标记(tag),标记通常都是通过反射技术获取到.结构体标记语法

```
type 结构体名称 struct{
  属性名 类型 `key:"Value"`
}
```

* 获取结构体标记(tag)

```go
type People struct {
	Name    string `xml:"name"`
	Address string `xml:"address"`
}

func main() {
	t:=reflect.TypeOf(People{})
	name,_:=t.FieldByName("Name")
	fmt.Println(name.Tag)//获取完整标记
	fmt.Println(name.Tag.Get("xml"))//获取标记中xml对应内容
}
```


# 一.Go语言标准库提供的API

* 在encoding/xml包下提供了对XML`序列化和反序列化`的API
* 使用Unmarshal可以直接把XML字节切片数据转换为结构体
* 转换时按照特定的转换规则进行转换,且数据类型可以自动转换

```
* 如果结构体字段的类型为字符串或者[]byte，且标签为",innerxml"，
  Unmarshal函数直接将对应原始XML文本写入该字段，其余规则仍适用。
* 如果结构体字段类型为xml.Name且名为XMLName，Unmarshal会将元素名写入该字段
* 如果字段XMLName的标签的格式为"name"或"namespace-URL name"，
  XML元素必须有给定的名字（以及可选的名字空间），否则Unmarshal会返回错误。
* 如果XML元素的属性的名字匹配某个标签",attr"为字段的字段名，或者匹配某个标签为"name,attr"
  的字段的标签名，Unmarshal会将该属性的值写入该字段。
* 如果XML元素包含字符数据，该数据会存入结构体中第一个具有标签",chardata"的字段中，
  该字段可以是字符串类型或者[]byte类型。如果没有这样的字段，字符数据会丢弃。
* 如果XML元素包含注释，该数据会存入结构体中第一个具有标签",comment"的字段中，
  该字段可以是字符串类型或者[]byte类型。如果没有这样的字段，字符数据会丢弃。
* 如果XML元素包含一个子元素，其名称匹配格式为"a"或"a>b>c"的标签的前缀，反序列化会深入
  XML结构中寻找具有指定名称的元素，并将最后端的元素映射到该标签所在的结构体字段。
  以">"开始的标签等价于以字段名开始并紧跟着">" 的标签。
* 如果XML元素包含一个子元素，其名称匹配某个结构体类型字段的XMLName字段的标签名，
  且该结构体字段本身没有显式指定标签名，Unmarshal会将该元素映射到该字段。
* 如果XML元素的包含一个子元素，其名称匹配够格结构体字段的字段名，且该字段没有任何模式选项
  （",attr"、",chardata"等），Unmarshal会将该元素映射到该字段。
* 如果XML元素包含的某个子元素不匹配以上任一条，而存在某个字段其标签为",any"，
  Unmarshal会将该元素映射到该字段。
* 匿名字段被处理为其字段好像位于外层结构体中一样。
* 标签为"-"的结构体字段永不会被反序列化填写。
```

# 二. XML文件读取

* 给定XML文件内容如下

```xml
<?xml version="1.0" encoding="UTF-8" ?>
<people id="666">
    <name>xiaopang</name>
    <address>广东深圳</address>
</people>
```

* 新建结构体,装载XML数据
  * 结构体中属性首字母必须大写,否则无法装配

```go
type People struct {
	XMLName xml.Name `xml:"people"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Address string   `xml:"address"`
}

func main() {
	peo := new(People)
	b, err := ioutil.ReadFile("easy-xml.xml")
	fmt.Println(string(b))
	fmt.Println("111:", err)
	err = xml.Unmarshal(b, peo)
	fmt.Println("2222", err)
	fmt.Println(peo)
}
```

# 三.多层嵌套XML文件读取

* 给定XML中数据如下

```xml
<?xml version="1.0" encoding="UTF-8" ?>
<peoples version="0.9">
    <people id="666">
        <name>xiaopang</name>
        <address>广东深圳</address>
    </people>
    <people id="555">
        <name>小明</name>
        <address>广东东莞</address>
    </people>
</peoples>

```

* 编写读取XML数据代码

```go
type Peoples struct {
	XMLName xml.Name `xml:"peoples"`
	Version string   `xml:"version,attr"`
	Peos    []People `xml:"people"`
}

type People struct {
	XMLName xml.Name `xml:"people"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Address string   `xml:"address"`
}

func main() {
	peo := new(Peoples)
	b, err := ioutil.ReadFile("complex-xml.xml")
	fmt.Println(string(b))
	fmt.Println("111:", err)
	err = xml.Unmarshal(b, peo)
	fmt.Println("2222", err)
	fmt.Println(peo)
}
```

# 一.生成XML

* 生成XML只要在学习下encoding/xml包下的Marshal()函数,结合输入流就可以完成xml文件生成
* 在encoding/xml中有常量,常量中是xml文档头

```go
const (
	// Header is a generic XML header suitable for use with the output of Marshal.
	// This is not automatically added to any output of this package,
	// it is provided as a convenience.
	Header = `<?xml version="1.0" encoding="UTF-8"?>` + "\n"
)
```

# 二.代码示例

* 使用Marshal()函数生成的[]byte没有格式化
* 使用MarshalIndent()可以对内容进行格式化
  * 第一个参数:结构体对象
  * 第二个参数:每行的前缀
  * 第三个参数:层级缩进内容

```go
type PeopleGen struct {
	XMLName xml.Name `xml:"people"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Address string   `xml:"address"`
}

func main() {
	peogen := PeopleGen{Id: 123, Name: "xiaopang", Address: "广东深圳"}
	b, _ := xml.MarshalIndent(peogen, "", "	")
	b = append([]byte(xml.Header), b...)
	ioutil.WriteFile("D:/okk.xml", b, 0666)
	fmt.Println("程序结束")
}
```