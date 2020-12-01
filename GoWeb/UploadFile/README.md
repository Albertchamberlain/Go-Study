# 一.文件上传

* 文件上传:客户端把上传文件转换为二进制流后发送给服务器,服务器对二进制流进行解析
* HTML表单(form)enctype(Encode Type)属性控制表单在提交数据到服务器时数据的编码类型.
  * enctype=”application/x-www-form-urlencoded” 默认值,表单数据会被编码为名称/值形式
  * enctype=”multipart/form-data” 编码成消息,每个控件对应消息的一部分.请求方式必须是post
  * enctype=”text/plain” 纯文本形式进行编码的


* 服务端可以使用FormFIle("name")获取上传到的文件,官方定义如下

```go
// FormFile returns the first file for the provided form key.
// FormFile calls ParseMultipartForm and ParseForm if necessary.
func (r *Request) FormFile(key string) (multipart.File, *multipart.FileHeader, error) {
	if r.MultipartForm == multipartByReader {
		return nil, nil, errors.New("http: multipart handled by MultipartReader")
	}
	if r.MultipartForm == nil {
		err := r.ParseMultipartForm(defaultMaxMemory)
		if err != nil {
			return nil, nil, err
		}
	}
	if r.MultipartForm != nil && r.MultipartForm.File != nil {
		if fhs := r.MultipartForm.File[key]; len(fhs) > 0 {
			f, err := fhs[0].Open()
			return f, fhs[0], err
		}
	}
	return nil, nil, ErrMissingFile
}
```

* multipart.File是文件对象

```go
// File is an interface to access the file part of a multipart message.
// Its contents may be either stored in memory or on disk.
// If stored on disk, the File's underlying concrete type will be an *os.File.
type File interface {
	io.Reader
	io.ReaderAt
	io.Seeker
	io.Closer
}
```

* 封装了文件的基本信息

```go
// A FileHeader describes a file part of a multipart request.
type FileHeader struct {
	Filename string					//文件名
	Header   textproto.MIMEHeader	//MIME信息
	Size     int64					//文件大小,单位bit

	content []byte					//文件内容,类型[]byte
	tmpfile string					//临时文件
}
```

* 服务器端获取客户端传递后的文件流,把文件保存到服务器即可


