package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
)

func main() {

	peogen := PeopleGen{Id: 123, Name: "xiaopang", Address: "广东深圳"}
	b, _ := xml.MarshalIndent(peogen, "", "	")
	b = append([]byte(xml.Header), b...)
	ioutil.WriteFile("D:\\GoPro\\GoBasic\\Others\\okk.xml", b, 0666)
	fmt.Println("程序结束")

	log.Println("打印日志信息")

	f13, _ := os.OpenFile("D:\\GoPro\\GoBasic\\Others\\golog.log", os.O_APPEND|os.O_CREATE, 07777)
	defer f13.Close()
	logger := log.New(f13, "[info]\t", log.Ltime)
	logger.Println("输出日志信息")

	log.Fatal("打印日志信息")

	t := reflect.TypeOf(People2{})
	name, _ := t.FieldByName("Name")
	fmt.Println(name.Tag)            //获取完整标记
	fmt.Println(name.Tag.Get("xml")) //获取标记中xml对应内容

	fmt.Println(f())  //输出:0
	fmt.Println(f2()) //输出:2

	fmt.Println("1")
	//panic("panic执行了,哈哈")
	fmt.Println("2")

	//defer func(){
	//	fmt.Println("defer执行")
	//}()

	fmt.Println("1")
	//panic("panic执行了,哈哈")
	fmt.Println("2")

	//defer func() {
	//	if error:=recover();error!=nil{
	//		fmt.Println("出现了panic,使用reover获取信息:",error)
	//	}
	//}()
	fmt.Println("11111111111")
	//panic("出现panic")
	fmt.Println("22222222222")

	fmt.Println("程序开始")
	//demo1()
	fmt.Println("程序结束")

	//u, _ := user.Lookup(``)
	//fmt.Println(u.Name)
	//fmt.Println(u.Gid)
	//fmt.Println(u.HomeDir)
	//fmt.Println(u.Uid)
	//fmt.Println(u.Username)

	error := os.MkdirAll("D:/okk/a/b", os.ModeDir)
	if error != nil {
		fmt.Println("文件夹创建失败", error)
		return
	}
	fmt.Println("文件夹创建成功")

	file, err := os.Create("D:/okk/test.txt")
	if err != nil {
		fmt.Println("文件创建失败,", err)
		return
	}
	fmt.Println("文件创建成功", file.Name())

	err = os.Rename("D:/okk/test.txt", "D:/okk/test1.txt")
	if err != nil {
		fmt.Println("重命名文件失败,", err)
		return
	}
	fmt.Println("文件重命名成功")

	err2 := os.Rename("D:/okk", "D:/ok")
	if err != nil {
		fmt.Println("重命名文件夹失败,", err2)
		return
	}
	fmt.Println("文件夹重命名成功")

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
	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.IsDir())
	fmt.Println(fileInfo.Mode())
	fmt.Println(fileInfo.ModTime())
	fmt.Println(fileInfo.Size())

	err3 := os.RemoveAll("D:/okk/a.txt")
	if err3 != nil {
		fmt.Println("删除失败", err3)
		return
	}
	fmt.Println("删除成功")

	a := 1.5
	fmt.Println(reflect.TypeOf(a))

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

	peo2 := new(People3)
	b12, err12 := ioutil.ReadFile("easy-xml.xml")
	fmt.Println(string(b12))
	fmt.Println("111:", err12)
	err = xml.Unmarshal(b12, peo2)
	fmt.Println("2222", err12)
	fmt.Println(peo2)

	peoples_com := new(PeoplesCom)
	b11, err11 := ioutil.ReadFile("complex-xml.xml")
	fmt.Println(string(b11))
	fmt.Println("111:", err11)
	err = xml.Unmarshal(b11, peoples_com)
	fmt.Println("2222", err11)
	fmt.Println(peoples_com)

}

type PeopleGen struct {
	XMLName xml.Name `xml:"people"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Address string   `xml:"address"`
}
type PeoplesCom struct {
	XMLName xml.Name    `xml:"peoples"`
	Version string      `xml:"version,attr"`
	Peos    []PeopleSon `xml:"people"`
}
type PeopleSon struct {
	XMLName xml.Name `xml:"people"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Address string   `xml:"address"`
}

type People struct {
	Id   int
	Name string
}

func f2() (i int) {
	defer func() {
		i = i + 2
	}()
	return
}

func f() int {
	i := 0
	defer func() {
		i = i + 2
	}()
	return i
}

func demo1() {
	fmt.Println("demo1上半部分")
	demo2()
	fmt.Println("demo1下半部分")
}
func demo2() {
	defer func() {
		recover() //此处进行恢复
	}()
	fmt.Println("demo2上半部分")
	demo3()
	fmt.Println("demo2下半部分")
}
func demo3() {
	fmt.Println("demo3上半部分")
	panic("在demo3出现了panic")
	fmt.Println("demo3下半部分")
}

type People2 struct {
	Name    string `xml:"name"`
	Address string `xml:"address"`
}

type People3 struct {
	XMLName xml.Name `xml:"people"`
	Id      int      `xml:"id,attr"`
	Name    string   `xml:"name"`
	Address string   `xml:"address"`
}
