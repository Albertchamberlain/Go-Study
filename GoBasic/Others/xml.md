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