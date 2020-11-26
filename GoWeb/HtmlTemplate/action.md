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