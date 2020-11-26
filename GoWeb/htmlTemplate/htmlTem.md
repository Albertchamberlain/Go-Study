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