# 一. 项目结构

* 在Go语言中web项目标准结构如下

```
--项目名
	--src
	--static
		--css
		--images
		--js
	--view
		--index.html
	--main.go
```

* Go语言标准库中html/template包提供了html模版支持,把HTML当作模版可以在访问控制器时显示HTML模版信息

  * 这也符合标准的MVC思想

# 二.HTML模版显示

* 使用template.ParseFiles()可以解析多个模版文件

```go

func ParseFiles(filenames ...string) (*Template, error) {
	return parseFiles(nil, filenames...)
}
```

* 把`模版信息`响应写入到`输出流`中

```go

func (t *Template) Execute(wr io.Writer, data interface{}) error {
	if err := t.escape(); err != nil {
		return err
	}
	return t.text.Execute(wr, data)
}
```

* 代码演示,显示index.html信息
  * 因为配置的pattern为"/"所以资源路径任意,都可以访问到这个HTML

```go
package main

import (
	"net/http"
	"html/template"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view/index.html")
	t.Execute(w, nil) //第二个参数表示向模版传递的数据
}

func main() {
	server := http.Server{Addr: ":8090"}
	http.HandleFunc("/", welcome)
	server.ListenAndServe()
}
```

# 三.引用静态文件

* 把静态文件放入到特定的文件夹中,使用Go语言的文件服务就可以进行加载
* 项目结构

```
--项目
	--static
		--js
			--index.js
	--view
		--index.html
	--main.go
```

* index.html代码如下

```html
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
        "http://www.w3.org/TR/html4/loose.dtd">
<html>
<head>
    <title>Title</title>
    <!--路径以斜杠开头,表示项目根目录下-->
    <script type="text/javascript" src="/static/js/index.js"></script>
</head>
<body>
    这是要显示的html页面信息<br/>
    <button onclick="myclick()">按钮</button>
</body>
</html>
```

* index.js代码如下

```javascript
function myclick(){
    alert("您点击了按钮")
}
```

* 代码示例

```go
package main

import (
	"net/http"
	"html/template"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view/index.html")
	t.Execute(w, nil) //第二个参数表示向模版传递的数据
}

func main() {
	server := http.Server{Addr: ":8090"}
	/*
	访问url以/static/开头,就会把访问信息映射到指定的目录中
	 */
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", welcome)
	server.ListenAndServe()
}
```

# 向模版传递数据

* 可以在HTML中使用{{}}获取template.Execute()第二个参数传递的值
* 最常用的`{{.}}`中的"."是指针,指向当前变量,称为"dot"
* 在{{}}可以有的Argument,官方给定如下

```
- go语法的布尔值、字符串、字符、整数、浮点数、虚数、复数，视为无类型字面常数，字符串不能跨行
- 关键字nil，代表一个go的无类型的nil值
- 字符'.'（句点，用时不加单引号），代表dot的值
- 变量名，以美元符号起始加上（可为空的）字母和数字构成的字符串，如：$piOver2和$；
  执行结果为变量的值，变量参见下面的介绍
- 结构体数据的字段名，以句点起始，如：.Field；
  执行结果为字段的值，支持链式调用：.Field1.Field2；
  字段也可以在变量上使用（包括链式调用）：$x.Field1.Field2；
- 字典类型数据的键名；以句点起始，如：.Key；
  执行结果是该键在字典中对应的成员元素的值；
  键也可以和字段配合做链式调用，深度不限：.Field1.Key1.Field2.Key2；
  虽然键也必须是字母和数字构成的标识字符串，但不需要以大写字母起始；
  键也可以用于变量（包括链式调用）：$x.key1.key2；
- 数据的无参数方法名，以句点为起始，如：.Method；
  执行结果为dot调用该方法的返回值，dot.Method()；
  该方法必须有1到2个返回值，如果有2个则后一个必须是error接口类型；
  如果有2个返回值的方法返回的error非nil，模板执行会中断并返回给调用模板执行者该错误；
  方法可和字段、键配合做链式调用，深度不限：.Field1.Key1.Method1.Field2.Key2.Method2；
  方法也可以在变量上使用（包括链式调用）：$x.Method1.Field；
- 无参数的函数名，如：fun；
  执行结果是调用该函数的返回值fun()；对返回值的要求和方法一样；函数和函数名细节参见后面。
- 上面某一条的实例加上括弧（用于分组）
  执行结果可以访问其字段或者键对应的值：
      print (.F1 arg1) (.F2 arg2)
      (.StructValuedMethod "arg").Field
```

* 向HTML传递字符串数据.在HTML中使用{{.}}获取传递数据即可.所有基本类型都是使用此方式进行传递

```html
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
        "http://www.w3.org/TR/html4/loose.dtd">
<html>
<head>
    <title>Title</title>
</head>
<body>
<pre>
尊敬的{{.}}先生/女士
    您已经被我公司录取,收到此消息后请您仔细阅读附件中"注意事项"
    再次祝您:{{.}}好运
</pre>
</body>
</html>
```

```go
package main

import (
	"net/http"
	"html/template"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view/index.html")
	t.Execute(w, "smallming") //此处传递数据
}

func main() {
	server := http.Server{Addr: ":8090"}
	http.HandleFunc("/", welcome)
	server.ListenAndServe()
}

```

# 传递结构体类型数据

* 结构体的属性首字母必须大写才能被模版访问
* 在模版中直接使用`{{.属性名}}`获取结构体的属性
* HTML代码如下

```html
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
        "http://www.w3.org/TR/html4/loose.dtd">
<html>
<head>
    <title>Title</title>
</head>
<body>
<pre>
当前登录用户信息:<br/>
    姓名:{{.Name}}<br/>
    年龄:{{.Age}}
</pre>
</body>
</html>
```

* go文件代码如下

```go
package main

import (
	"net/http"
	"html/template"
)
//注意:只有首字母大写的属性才能在模版中访问到
type User struct {
	Name string
	Age  int
}

func welcome(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view/index.html")
	t.Execute(w, User{"张三", 20}) //此处传递数据
}

func main() {
	server := http.Server{Addr: ":8090"}
	http.HandleFunc("/", welcome)
	server.ListenAndServe()
}

```

# 向模版传递map类型数据

* 直接使用`{{.key}}`获取map中数据
* 模版中支持连缀写法(不仅仅是map)
* go文件代码如下

```go
package main

import (
	"net/http"
	"html/template"
)

//注意:只有首字母大写的属性才能在模版中访问到
type User struct {
	Name string
	Age  int
}

func welcome(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("view/index.html")
	m := make(map[string]interface{})
	m["user"] = User{"张三", 20}
	m["money"] = 998
	t.Execute(w, m) //此处传递数据
}

func main() {
	server := http.Server{Addr: ":8090"}
	http.HandleFunc("/", welcome)
	server.ListenAndServe()
}

```

* HTML代码如下,里面使用了连缀写法

```html
<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN"
        "http://www.w3.org/TR/html4/loose.dtd">
<html>
<head>
    <title>Title</title>
</head>
<body>
<pre>
当前登录用户信息:<br/>
    姓名:{{.user.Name}}<br/>
    年龄:{{.user.Age}}<br/>
    购物金额:{{.money}}
</pre>
</body>
</html>
```



# 一. Action

* Go语言官方文档给出action（动作）的列表。"Arguments"和"pipelines"代表数据的执行结果

```
{{/* a comment */}}
    注释，执行时会忽略。可以多行。注释不能嵌套，并且必须紧贴分界符始止，就像这里表示的一样。
{{pipeline}}
    pipeline的值的默认文本表示会被拷贝到输出里。
{{if pipeline}} T1 {{end}}
    如果pipeline的值为empty，不产生输出，否则输出T1执行结果。不改变dot的值。
    Empty值包括false、0、任意nil指针或者nil接口，任意长度为0的数组、切片、字典。
{{if pipeline}} T1 {{else}} T0 {{end}}
    如果pipeline的值为empty，输出T0执行结果，否则输出T1执行结果。不改变dot的值。
{{if pipeline}} T1 {{else if pipeline}} T0 {{end}}
    用于简化if-else链条，else action可以直接包含另一个if；等价于：
        {{if pipeline}} T1 {{else}}{{if pipeline}} T0 {{end}}{{end}}
{{range pipeline}} T1 {{end}}
    pipeline的值必须是数组、切片、字典或者通道。
    如果pipeline的值其长度为0，不会有任何输出；
    否则dot依次设为数组、切片、字典或者通道的每一个成员元素并执行T1；
    如果pipeline的值为字典，且键可排序的基本类型，元素也会按键的顺序排序。
{{range pipeline}} T1 {{else}} T0 {{end}}
    pipeline的值必须是数组、切片、字典或者通道。
    如果pipeline的值其长度为0，不改变dot的值并执行T0；否则会修改dot并执行T1。
{{template "name"}}
    执行名为name的模板，提供给模板的参数为nil，如模板不存在输出为""
{{template "name" pipeline}}
    执行名为name的模板，提供给模板的参数为pipeline的值。
{{with pipeline}} T1 {{end}}
    如果pipeline为empty不产生输出，否则将dot设为pipeline的值并执行T1。不修改外面的dot。
{{with pipeline}} T1 {{else}} T0 {{end}}
    如果pipeline为empty，不改变dot并执行T0，否则dot设为pipeline的值并执行T1。
```

* action主要完成流程控制、循环、模版等操作.通过使用action可以在模版中完成简单逻辑处理(复杂逻辑处理应该在go中实现,传递给模版的数据应该是已经加工完的数据)


# 二. if 使用

* if写在模版中和写在go文件中功能是相同的,区别是语法
* 布尔函数会将任何类型的零值视为假，其余视为真。
* if后面的表达式中如果包含逻辑控制符在模版中实际上是全局函数

```
and
    函数返回它的第一个empty参数或者最后一个参数；
    就是说"and x y"等价于"if x then y else x"；所有参数都会执行；
or
    返回第一个非empty参数或者最后一个参数；
    亦即"or x y"等价于"if x then x else y"；所有参数都会执行；
not
    返回它的单个参数的布尔值的否定
len
    返回它的参数的整数类型长度
index
    执行结果为第一个参数以剩下的参数为索引/键指向的值；
    如"index x 1 2 3"返回x[1][2][3]的值；每个被索引的主体必须是数组、切片或者字典。
print
    即fmt.Sprint
printf
    即fmt.Sprintf
println
    即fmt.Sprintln
html
    返回其参数文本表示的HTML逸码等价表示。
urlquery
    返回其参数文本表示的可嵌入URL查询的逸码等价表示。
js
    返回其参数文本表示的JavaScript逸码等价表示。
call
    执行结果是调用第一个参数的返回值，该参数必须是函数类型，其余参数作为调用该函数的参数；
    如"call .X.Y 1 2"等价于go语言里的dot.X.Y(1, 2)；
    其中Y是函数类型的字段或者字典的值，或者其他类似情况；
    call的第一个参数的执行结果必须是函数类型的值（和预定义函数如print明显不同）；
    该函数类型值必须有1到2个返回值，如果有2个则后一个必须是error接口类型；
    如果有2个返回值的方法返回的error非nil，模板执行会中断并返回给调用模板执行者该错误；
```

* 二元比较运算的集合：(也是函数,函数具有两个参数,满足参数语法)

```
eq      如果arg1 == arg2则返回真
ne      如果arg1 != arg2则返回真
lt      如果arg1 < arg2则返回真
le      如果arg1 <= arg2则返回真
gt      如果arg1 > arg2则返回真
ge      如果arg1 >= arg2则返回真
```

* 简单if示例-go文件

```go
package main

import (
   "net/http"
   "html/template"
)

func test(rw http.ResponseWriter, r *http.Request) {
   t, _ := template.ParseFiles("template/html/if.html")
   //第二个参数传递类型默认值:nil,"",0,false等都会导致if不成立
   t.Execute(rw, "")
}

func main() {
   //创建server服务
   server := http.Server{Addr: ":8090"}

   //设置处理器函数
   http.HandleFunc("/test", test)

   //监听和开始服务
   server.ListenAndServe()
}
```

* 简单if示例-html文件

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>if测试</title>
</head>
<body>
测试if是否执行<br/>
{{if . }}
if成立这个位置输出
{{end}}
</body>
</html>
```

* 直接在HTMl中定义变量演示if..else用法(go文件不变)

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>测试</title>
</head>
<body>
{{$n:=123}}
{{if ne $n 123}}
if成立这个位置输出
{{else}}{{/* 比if结构多了else */}}
这是else的功能
{{end}}
</body>
</html>
```

* go文件不变,演示if...else if...else用法

```html
<body>
{{$n:=124}}
{{if eq $n 123}}
123
{{else if eq $n 124}}
124
{{else if eq $n 125}}
125
{{else}}
else
{{end}}
</body>
```

* 在模版中也可以相互嵌套

```html
<body>
{{$n:=124}}
{{if gt $n 100}}
    {{if gt $n 200}}
        gt 200
    {{else}}
        lt 200
    {{end}}
{{else}}
    小于100
{{end}}
</body>
```

# 三.range使用

* range遍历数组或切片或map或channel时,在range内容中{{.}}表示获取迭代变量

```html
<body>
{{range .}}
    {{.}}{{/* 此处dot为迭代变量 */}}
{{end}}

{{.}}{{/* 此处获取还是传递给模版的切片 */}}
</body>
```



# 一.模版嵌套

* 在实际项目中经常出现页面复用的情况,例如:整个网站的头部信息和底部信息复用
* 可以使用动作{{template "模版名称"}}引用模版
* 引用的模版必须在HTML中定义这个模版

```html
{{define "名称"}}
html
{{end}}
```

* 执行主模版时也要给主模版一个名称,执行时调用的是ExecuteTemplate()方法




# 二. 调用模版时同时传递参数

* 如果直接引用html可以直接使用html标签的`<iframe>`,但是要动态效果时,可以在调用模版给模版传递参数



* 在子模版中依然是使用{{.}}获取传递过来的参数



