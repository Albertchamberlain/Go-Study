# 一.反射介绍

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