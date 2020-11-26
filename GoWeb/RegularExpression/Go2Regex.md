# Go语言对正则的支持

* 在`regexp`包中提供了对正则表达式的支持,并提供了RegExp结构体
  * 可以看出里面有`互斥锁`,所以在并发下是安全的

```go
// Regexp is the representation of a compiled regular expression.
// A Regexp is safe for concurrent use by multiple goroutines,
// except for configuration methods, such as Longest.
type Regexp struct {
	// read-only after Compile
	regexpRO

	// cache of machines for running regexp
	mu      sync.Mutex
	machine []*machine
}
```

* 判断字符串是否与正则匹配最简单的办法是


```go
	result,_:=regexp.MatchString(`^\d\w$`,"5A")
	fmt.Println(result)
```

* 如果需要更多的功能,可以使用Regexp的方式实现,下面列举除了一些常用方法

```go
package main

import (
	"regexp"
	"fmt"
)

func main() {
	//创建结构体变量
	r := regexp.MustCompile(`\d[a-zA-Z]`)
	//判断是否匹配
	fmt.Println(r.MatchString("5A1"))
	/*
	字符串中满足要求的片段,返回[]string
	第二个参数是[]string的长度,-1表示不限制长度
	 */
	fmt.Println(r.FindAllString("56A6B7C", -1))
	/*
	把正则表达式匹配的结果当作拆分符,拆分字符串
	n > 0 : 返回最多n个子字符串，最后一个子字符串是剩余未进行分割的部分。
	n == 0: 返回nil (zero substrings)
	n < 0 : 返回所有子字符串
	 */
	fmt.Println(r.Split("12345qwert", -1))
	//把满足正则要求内容替换成指定字符串
	fmt.Println(r.ReplaceAllString("12345qwert", "替换了"))
}

```

# 三.服务器端数据校验

* 数据校验可以有客户端数据校验和服务器端数据校验.双重保证是保证程序安全性的有效措施
* 客户端向服务端发送请求参数,服务器端接收到请求参数后使用正则验证,验证结果通过才能正确执行,例如注册时验证数据格式
