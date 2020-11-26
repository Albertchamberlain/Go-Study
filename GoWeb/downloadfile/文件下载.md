# 一.文件下载简介

* 文件下载总体步骤
  * 客户端向服务端发起请求,请求参数包含要下载文件的名称
  * 服务器接收到客户端请求后把文件设置到响应对象中,响应给客户端浏览器
* 载时需要设置的响应头信息
  * Content-Type: 内容MIME类型
    * application/octet-stream 任意类型
  * Content-Disposition:客户端对内容的操作方式
    * inline 默认值,表示浏览器能解析就解析,不能解析下载
    * attachment;filename=下载时显示的文件名 ,客户端浏览器恒下载


