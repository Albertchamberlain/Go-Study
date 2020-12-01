# 一.sort包

* Go语言标准库中sort提供了排序API
* sort包提供了多种排序算法,这些算法是内部实现的,每次使用sort包排序时,会`自动选择`最优算法实现
  * `插入排序`
  * `快速排序`
  * `堆排`
* sort包中最上层是一个名称为`Interface`的接口,只要满足sort.Interface类型都可以排序

```go
// A type, typically a collection, that satisfies sort.Interface can be
// sorted by the routines in this package. The methods require that the
// elements of the collection be enumerated by an integer index.
type Interface interface {
	// Len is the number of elements in the collection.
	Len() int
	// Less reports whether the element with
	// index i should sort before the element with index j.
	Less(i, j int) bool
	// Swap swaps the elements with indexes i and j.
	Swap(i, j int)
}
```

* Go语言标准库默认提供了对int、float64、string进行排序的API
* 很多函数的参数都是sort包下类型,需要进行转换.

# 二.排序实现

* 对int类型切片排序

```go
	num := [] int{1, 7, 5, 2, 6}
	sort.Ints(num) //升序
	fmt.Println(num)
	sort.Sort(sort.Reverse(sort.IntSlice(num))) //降序
	fmt.Println(num)
```

* 对float64类型切片排序

```go
	f := [] float64{1.5, 7.2, 5.8, 2.3, 6.9}
	sort.Float64s(f) //升序
	fmt.Println(f)
	sort.Sort(sort.Reverse(sort.Float64Slice(f))) //降序
	fmt.Println(f)
```

* 对string类型切片排序
  * 按照编码表数值进行排序
  * 多字符串中按照第一个字符进行比较
  * 如果第一个字符相同,比较第二个字符

```go
	s := []string{"我", "我是中国人", "a", "d", "国家", "你", "我a"}
	sort.Sort(sort.StringSlice(s)) //升序
	fmt.Println(s)
	//查找内容的索引,如果不存在,返回内容应该在升序排序切片的哪个位置插入
	fmt.Println(sort.SearchStrings(s, "你是"))
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	fmt.Println(s)
```





* math包提供了基本数学常数和数学函数以及数学常数


```go
// Mathematical constants.
const (
	E   = 2.71828182845904523536028747135266249775724709369995957496696763 
	Pi  = 3.14159265358979323846264338327950288419716939937510582097494459 
	Phi = 1.61803398874989484820458683436563811772030917980576286213544862 
	Sqrt2   = 1.41421356237309504880168872420969807856967187537694807317667974
	SqrtE   = 1.64872127070012814684865078781416357165377610071014801157507931 
	SqrtPi  = 1.77245385090551602729816748334114518279754945612238712821380779 
	SqrtPhi = 1.27201964951406896425242246173749149171560804184009624861664038 
	Ln2    = 0.693147180559945309417232121458176568075500134360255254120680009 
	Log2E  = 1 / Ln2
	Ln10   = 2.30258509299404568401799145468436420760110148862877297603332790 
	Log10E = 1 / Ln10
)

// Floating-point limit values.
// Max is the largest finite value representable by the type.
// SmallestNonzero is the smallest positive, non-zero value representable by the type.
const (
	MaxFloat32= 3.40282346638528859811704183484516925440e+38  // 2**127 * (2**24 - 1) / 2**23
	SmallestNonzeroFloat32 = 1.401298464324817070923729583289916131280e-45 // 1 / 2**(127 - 1 + 23)

	MaxFloat64= 1.797693134862315708145274237317043567981e+308 // 2**1023 * (2**53 - 1) / 2**52
	SmallestNonzeroFloat64 = 4.940656458412465441765687928682213723651e-324 // 1 / 2**(1023 - 1 + 52)
)

// Integer limit values.
const (
	MaxInt8   = 1<<7 - 1
	MinInt8   = -1 << 7
	MaxInt16  = 1<<15 - 1
	MinInt16  = -1 << 15
	MaxInt32  = 1<<31 - 1
	MinInt32  = -1 << 31
	MaxInt64  = 1<<63 - 1
	MinInt64  = -1 << 63
	MaxUint8  = 1<<8 - 1
	MaxUint16 = 1<<16 - 1
	MaxUint32 = 1<<32 - 1
	MaxUint64 = 1<<64 - 1
)
```

* 列举出常用的数学函数

```go
	var i, j float64 = 12.3, 9.6
	//向下取整,
	fmt.Println(math.Floor(i)) //输出:12
	//向上取整
	fmt.Println(math.Ceil(i)) //输出:13
	//绝对值
	fmt.Println(math.Abs(i)) //输出:12.3
	//返回值分别整数位和小数位,小数位可能出现误差
	num, decimal := math.Modf(i)
	fmt.Println(num, decimal)
	//返回两个变量中大的值
	fmt.Println(math.Max(i, j)) //输出:12.3
	//返回两个变量中小的值
	fmt.Println(math.Min(i, j)) //输出:9.6
	//x的y次方
	fmt.Println(math.Pow(3, 2)) //输出:输出9
	//四舍五入
	fmt.Println(math.Round(i)) //输出:12
```


* math/rand实现了伪随机数生成器
* 在Go语言中随机数需要设置种子,如果不设置种子随机数的结果每次运行都相同。
* 默认种子是1,且相同种子产生的随机数是相同的.
* 可以使用当前时间的纳秒差计算随机数,在一定程度上保证了种子的唯一性

```go
	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Int63n(10))
```

* 在Go语言中时间类型不是关键字而是使用time包下Time结构体
* 时间类型默认显示为`UTC`,所以经常需要把`时间类型`转换为`字符串`,显示成我们所熟悉的格式


## Time

- 声明Time时,默认时间是无意义的

```
func main() {
	var t time.Time
	fmt.Println(t)//输出:0001-01-01 00:00:00 +0000 UTC
}
```

- 可以通过time包下的Now()函数获取操作系统当前时间

```
t := time.Now()
	fmt.Println(t) //输出:年-月-日 小时:分钟:秒.纳秒 +0800 CST m=+0.003012301
```


- 通过时间戳创建时间类型变量(距离1970年1月1日的纳秒差)

```
/*
	1秒(s)=1000毫秒(ms)
	1秒(s)=1000000微秒(μs)
	1秒(s)=1000000000纳秒(ns)
	 */
	t := time.Now()
	t1 := time.Unix(0, t.UnixNano()) //根据时间戳创建时间.第二个值[0, 999999999]外合法
	fmt.Println(t.String())
	fmt.Println(t1)
```

- 根据自己要求创建时间

```
//time.Local取到本地时间位置对象,东八区
	t := time.Date(2020, 11, 23, 7, 8, 9, 0, time.Local)
	fmt.Println(t) //输出:2020-05-06 07:08:09 +0800 CST
```

### time包下提供了大量的函数或方法获取时间的某一项

    t := time.Now()
    fmt.Println(t)
    fmt.Println(t.Year())       //年
    fmt.Println(int(t.Month())) //月
    fmt.Println(t.Day())        //日
    fmt.Println(t.Date())       //三个参数,分别是:年,月,日
    fmt.Println(t.Hour())       //小时
    fmt.Println(t.Minute())     //分钟
    fmt.Println(t.Second())     //秒
    fmt.Println(t.Clock())      //三个参数,分别是:小时,分钟,秒
    fmt.Println(t.Nanosecond()) //纳秒
    fmt.Println(t.Unix())       //秒差
    fmt.Println(t.UnixNano())   //纳秒差


## 时间和string相互转换

- 时间转换为string

```
	t := time.Now()
	//参数必须是这个时间,格式任意
	s := t.Format("2020-11-23 7:04:05", )
	fmt.Println(s)
```

- string转换为Time

```
s:="2022-02-04 22:02:04"
	t,err:=time.Parse("2006-01-02 15:04:05",s)
	fmt.Println(t,err)
```


# Go自定义结构体多参数排序

- 单一参数排序 siparam.go
- 多参数排序 (修改比较方法的逻辑即可)  mulparam.go


# 结构体数组排序
- 结构体数组根据字段排序  sortstructslice.go


# 利用slice（）相关函数排序
1. Slice() 不稳定排序
2. SliceStable() 稳定排序
3. SliceIsSorted() 判断是否已排序



## 排序总结，根据需求修改Less函数即可






