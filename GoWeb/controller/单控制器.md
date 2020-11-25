# 一. 单控制器

* 在Golang的net/http包下有ServeMux实现了Front设计模式的Front窗口,ServeMux负责接收请求并把请求分发给处理器(Handler)
* http.ServeMux实现了Handler接口

```go
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
type ServeMux struct {
	mu    sync.RWMutex
	m     map[string]muxEntry
	hosts bool // whether any patterns contain hostnames
}
func (mux *ServeMux) ServeHTTP(w ResponseWriter, r *Request) {
	if r.RequestURI == "*" {
		if r.ProtoAtLeast(1, 1) {
			w.Header().Set("Connection", "close")
		}
		w.WriteHeader(StatusBadRequest)
		return
	}
	h, _ := mux.Handler(r)
	h.ServeHTTP(w, r)
}
```

* 自定义结构体,实现Handler接口后,这个结构体就属于一个处理器,可以处理全部请求
  * 无论在浏览器中输入的资源地址是什么,都可以访问ServeHTTP

```go
package main

import "fmt"
import "net/http"

type MyHandler struct {
}

func (mh *MyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
   fmt.Fprintln(res,"输出内容")
}

func main() {
   myhandler := MyHandler{}
   server := http.Server{
      Addr:    "127.0.0.1:8090",
      Handler: &myhandler,
   }
   server.ListenAndServe()
}
```


# 二.多控制器

* 在实际开发中大部分情况是不应该只有一个控制器的,不同的请求应该交给不同的处理单元.在Golang中支持两种多处理方式
  * 多个处理器(Handler)
  * 多个处理函数(HandleFunc)
* 使用多处理器
  * 使用http.Handle把不同的URL绑定到不同的处理器
  * 在浏览器中输入http://localhost:8090/myhandler或http://localhost:8090/myother可以访问两个处理器方法.但是其他URl会出现404(资源未找到)页面

```go
package main

import "fmt"
import "net/http"

type MyHandler struct{}
type MyOtherHandler struct{}

func (mh *MyHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
   fmt.Fprintln(res, "第一个")
}
func (mh *MyOtherHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
   fmt.Fprintln(res, "第二个")
}

func main() {
   myhandler := MyHandler{}
   myother := MyOtherHandler{}
   server := http.Server{
      Addr: "localhost:8090",
   }
   http.Handle("/myhandler", &myhandler)
   http.Handle("/myother", &myother)
   server.ListenAndServe()
}
```

* 多函数方式要比多处理器方式简便.直接把资源路径与函数绑定

```go
package main

import "fmt"
import "net/http"

//不需要定义结构体
//函数的参数需要按照ServeHTTP函数参数列表进行定义
func first(res http.ResponseWriter, req *http.Request) {
   fmt.Fprintln(res, "第一个")
}
func second(res http.ResponseWriter, req *http.Request) {
   fmt.Fprintln(res, "第二个")
}

func main() {
   server := http.Server{
      Addr: "localhost:8090",
   }
   //注意此处使用HandleFunc函数
   http.HandleFunc("/first", first)
   http.HandleFunc("/second", second)
   server.ListenAndServe()
}
```