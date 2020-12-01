# 一. ioutil包

* ioutil包下提供了对`文件读写`的工具函数,通过这些函数快速实现文件的`读写操作`
* ioutil包下提供的函数比较少,但是都是很方便使用的函数

```
func NopCloser(r io.Reader) io.ReadCloser
func ReadAll(r io.Reader) ([]byte, error)
func ReadFile(filename string) ([]byte, error)
func WriteFile(filename string, data []byte, perm os.FileMode) error
func ReadDir(dirname string) ([]os.FileInfo, error)
func TempDir(dir, prefix string) (name string, err error)
func TempFile(dir, prefix string) (f *os.File, err error)
```

# 二.代码演示

* 打开完文件后可以使用ReadAll把文件中所有内容都读取到

```go
	f, err := os.Open("D:/go.txt")
	defer f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("文件中内容:\n", string(b))
```

* 也可以直接读取文件中内容

```go
	b, err := ioutil.ReadFile("D:/go.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(b))
```

* 写文件也很简单,直接使用WriteFile函数即可,但是源码中已经规定此文件只能是可写状态,且不是尾加数据

```go
	err := ioutil.WriteFile("D:/abc.txt", []byte("内容123123"), 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("数据写入成功")
```

* 还提供了`快速获取`某个文件夹中所有文件信息的函数

```go
	fs,_:=ioutil.ReadDir("D:/")
	for _,n := range fs {
		fmt.Println(n.Name())
	}
```

# 一.输入流

* 流(stream)是应用程序和外部资源进行数据交互的纽带

* 流分为输入流和输出流,输入和输出都是相对于`程序`,把外部数据传入到程序中叫做输入,反之叫做输出流

* 输入流(Input Stream),输入流(Output Stream) 平时所说的I/O流

* 在Go语言标准库中io包下是`Reader`接口表示输入流,只要实现这个接口就属于输入流


# 二.代码演示

* 可以使用strings包下的NewReader创建字符串流

```go
	r := strings.NewReader("hello 世界")
	b := make([]byte, r.Size())//创建字节切片,存放流中数据,根据流数据大小创建切片大小
	n, err := r.Read(b)//把流中数据读取到切片中
	if err != nil {
		fmt.Println("读取失败,", err)
		return
	}
	fmt.Println("读取数据长度,", n)

	fmt.Println("流中数据",string(b))//以字符串形式输入切片中数据
```

* 最常用的是文件流,把外部文件中数据读取到程序中

```go
	f, err := os.Open("D:/go.txt")//打开文件
	defer f.Close()
	if err != nil {
		fmt.Println("文件读取失败,", err)
		return
	}
	fileInfo, err := f.Stat()//获取文件信息
	if err != nil {
		fmt.Println("文件信息获取失败,", err)
		return
	}
	b := make([]byte, fileInfo.Size())//根据文件中数据大小创建切片
	_, err = f.Read(b)//读取数据到切片中
	if err != nil {
		fmt.Println("文件流读取失败:", err)
		return
	}
	fmt.Println("文件中内容为:", string(b))//以字符串形式输入切片中数据
```


# 一. 输入流

* 输入流就是把程序中数据写出到外部资源
* Go语言标准库中输出流是Writer接口

```go
type Writer interface {
	Write(p []byte) (n int, err error)
}
```


# 二.代码操作

* 注意:输入流时不要使用`os.Open()`因为这种方式获取的文件是`只读`的

```go
	fp := "D:/go.txt"
	/*
	第三个参数表示文件权限
	第 1 位在权限中总是为 0
	第 2 位为 0 表示文件不可以被读， 为 1 表示可以被读
	第 3 位为 0 表示文件不可以被写， 为 1 表示可以被写
	第 4 位为 0 表示文件不可以被执行， 为 1 表示可以被执行
	整理如下:
	   0(0000): 不可读写,不能被执行
	   1(0001): 不可读写,能被执行
	   2(0010): 可写不可读,不能被执行
	   3(0011): 可写不可读,能被执行
	   4(0100): 可读不可写,不能被执行
	   5(0101): 可读不可写,能被执行
	   6(0110): 可读写,不能执行
	   7(0111): 可读写,可执行

	0666:
	第一个 0 表示这个数是 八进制
	第一个 6 表示文件拥有者有读写权限，但没有执行权限
	第二个 6 表示文件拥有者同组用户有读写权限，但没有执行权限
	第三个 6 表示其它用户有读写权限，但没有执行权限

	 */

	//第二个参数表示文件内容追加
	//第三个参数表示创建文件时文件权限
	f, err := os.OpenFile(fp, os.O_APPEND, 0660)
	defer f.Close()
	if err != nil {
		fmt.Println("文件不存在,创建文件")
		f, _ = os.Create(fp)
	}

	/*
	内容中识别特殊字符
	\r\n 换行
	\t 缩进
	 */

	/*
	使用文件对象重写的Writer接口,参数是[]byte
	 */
	f.Write([]byte("使用Writer接口写数据\r\n"))

	/*
	使用stringWriter接口的方法,参数是字符串,使用更方便
	 */
	f.WriteString("写了\t一段\r\n内容123")
	fmt.Println("程序执行结束")
```
