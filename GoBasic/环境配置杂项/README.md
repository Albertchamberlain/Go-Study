## 环境配置简要步骤


下载好压缩包后解压，配置到系统环境变量里即可(简略步骤如下⚙)


* "我的电脑" --> 右键"属性"--> "高级" --> "环境变量" --> "系统变量"--> "新建"
* 在"系统变量"中PATH属性里面添加` %GOROOT%\bin; `
* DOS界面输入"go version" 出现版本信息即说明配置正确


## 无需重启使环境变量生效

以修改环境变量"PATH"为例,修改完成后,进入DOS命令提示符,输入：set
PATH=C:,关闭DOS窗口. 再次打开DOS窗口,输入：echo %PATH%
,可以发现"我的电脑"->"属性"->"高级"->"环境变量"中设置的 PATH 值已经生效.


## 目录结构解释

* **api** : 每个版本更新说明
* **bin** : go自带工具
* **blog**:博客
* **doc**:文档
* **misc**: 代码配置
* **lib**:额外引用
* **src**:标准库源码,以后第三方库放入到这个文件夹中.
* **test**:测试


# 一.Go语言安装包中自带工具

* 在%GOROOT%/bin中有三个工具
  * go.exe 编译、运行、构建
  * godoc.exe 查看包或函数的源码
  * gofmt.exe 格式化文件

```
--bin
	--go.exe
	--godoc.exe 
	--gofmt.exe
```

`godoc`
* 可以使用`godoc [包] [函数名]`查看包或函数的详细源码
* 源码在学习中非常重要,经常查看源码方便理解GO的原理

`gofmt工具`
* 规范的代码方便自己的阅读也方便别人的阅读.编写规范代码是每个程序的必修课
* gofmt工具可以帮助程序员把代码进行`格式化`,按照规范进行`格式化`
* 使用gofmt前提是`文件编译通过`

# 二.go.exe参数列表

* 在命令行中通过`go help`查看go参数如下

```
Usage:

        go command [arguments]

The commands are:

        build       compile packages and dependencies
        clean       remove object files and cached files
        doc         show documentation for package or symbol
        env         print Go environment information
        bug         start a bug report
        fix         update packages to use new APIs
        fmt         gofmt (reformat) package sources
        generate    generate Go files by processing source
        get         download and install packages and dependencies
        install     compile and install packages and dependencies
        list        list packages
        run         compile and run Go program
        test        test packages
        tool        run specified go tool
        version     print Go version
		vet         report likely mistakes in packages
```


# 三.常用参数解释

* `go version`查看Go语言版本
* `go env`查看Go语言详细环境
* `go list`查看Go语言文件目录
* `go build`把源码文件构建成系统可执行文件
* `go clean`清空生成的可执行文件
* `go vet`静态解析文件,检查是否有语法错误等
* `go get`从远程下载第三方Go语言库
* `go bug`提交bug
* `go test`测试(在后面章节中讲解)
* `go run`运行文件

