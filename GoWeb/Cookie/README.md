# 一.Cookie 简介

* Cookie就是客户端存储技术.以键值对的形式存在
* 在B/S架构中,服务器端产生Cookie响应给客户端,浏览器接收后把Cookie存在在特定的文件夹中,以后每次请求浏览器会把Cookie内容放入到请求中

# 二.Go语言对Cookie的支持

* 在net/http包下提供了Cookie结构体

  * Name设置Cookie的名称
  * Value 表示Cookie的值
  * Path 有效范围
  * Domain 可访问Cookie 的域
  * Expires 过期时间
  * MaxAge 最大存活时间,单位秒
  * HttpOnly 是否可以通过脚本访问

  ```go
  type Cookie struct {
  	Name  string
  	Value string
  
  	Path       string    // optional
  	Domain     string    // optional
  	Expires    time.Time // optional
  	RawExpires string    // for reading cookies only
  
  	// MaxAge=0 means no 'Max-Age' attribute specified.
  	// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
  	// MaxAge>0 means Max-Age attribute present and given in seconds
  	MaxAge   int
  	Secure   bool
  	HttpOnly bool
  	Raw      string
  	Unparsed []string // Raw text of unparsed attribute-value pairs
  }
  ```


  # 一.HttpOnly

* HttpOnly:控制Cookie的内容是否可以被JavaScript访问到。通过设置HttpOnly为true时防止XSS攻击防御手段之一
* 默认HttpOnly为false,表示客户端可以通过js获取
* 在项目中导入jquery.cookie.js库,使用jquery获取客户端Cookie内容


# 二. Path

* Path属性设置Cookie的访问范围
* 默认为”/”表示当前项目下所有都可以访问
* Path设置路径及子路径内容都可以访问


# 三.Expires

* Cookie默认存活时间是浏览器不关闭,当浏览器关闭后,Cookie失效
* 可以通过Expires设置具体什么时候过期,Cookie失效. 也可以通过MaxAge设置Cookie多长时间后实现
* IE6,7,8和很多浏览器不支持MaxAge,建议使用Expires
* Expires是time.Time类型,所以设置时需要明确设置过期时间


