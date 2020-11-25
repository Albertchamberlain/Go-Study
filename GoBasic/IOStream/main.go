package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	r := strings.NewReader("hello world")
	b := make([]byte, r.Size())
	n, err := r.Read(b)
	if err != nil {
		fmt.Println("读取失败,", err)
		return
	}
	fmt.Println("读取数据长度,", n)

	fmt.Println("流中数据", string(b))

	f, err2 := os.Open("D:/okk.txt")
	defer f.Close()
	if err2 != nil {
		fmt.Println("文件读取失败,", err2)
		return
	}
	fileInfo, err2 := f.Stat()
	if err2 != nil {
		fmt.Println("文件信息获取失败,", err2)
		return
	}
	b2 := make([]byte, fileInfo.Size())
	_, err2 = f.Read(b2)
	if err2 != nil {
		fmt.Println("文件流读取失败:", err2)
		return
	}
	fmt.Println("文件中内容为:", string(b2))

	fp := "D:/okk.txt"

	ff, err := os.OpenFile(fp, os.O_APPEND, 0660)
	defer ff.Close()
	if err != nil {
		fmt.Println("文件不存在,创建文件")
		ff, _ = os.Create(fp)
	}

	ff.Write([]byte("使用Writer接口写数据\r\n")) //使用文件对象重写的Writer接口,参数是[]byte

	ff.WriteString("写了\t一段\r\n内容123") //用stringWriter接口的方法,参数是字符串,使用更方便
	fmt.Println("程序执行结束")

	f2, errre := os.Open("D:/okk.txt")
	defer f2.Close()
	if errre != nil {
		fmt.Println(errre)
		return
	}
	b3, errre := ioutil.ReadAll(f2)
	if err != nil {
		fmt.Println(errre)
		return
	}
	fmt.Println("文件中内容:\n", string(b3))

	b4, err5 := ioutil.ReadFile("D:/okk.txt")
	if err5 != nil {
		fmt.Println(err5)
		return
	}
	fmt.Println(string(b4))

	err6 := ioutil.WriteFile("D:/okk.txt", []byte("内容123123"), 0666)
	if err6 != nil {
		fmt.Println(err6)
		return
	}
	fmt.Println("数据写入成功")

	fss, _ := ioutil.ReadDir("D:/")
	for _, nn := range fss {
		fmt.Println(nn.Name())
	}
}
