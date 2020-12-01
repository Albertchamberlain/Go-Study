# 一. 封装

* 封装主要体现在两个方面:封装`数据`、封装`业务`
* Go语言中通过首字母大小控制访问权限.属性首字母`小写`对外提供访问方法是封装数据最常见的实现方式
* 可以通过方法封装业务
  * 提出方法是封装
  * 控制结构体属性访问,对外提供访问方法也是封装
* 在面向对象中封装的好处:
  * 安全性.结构体属性访问受到限制,必须按照特定访问渠道
  * 可复用性,封装的方法实现可复用性
  * 可读写,多段增加代码可读性

# 二.代码实现

* Go语言同包任意位置可以访问全局内容,封装控制可以控制包外访问结构体中数据

```go
默认是函数,隶属于包,所以需要添加标识.告诉编译器这个方法`属性哪个结构体`
type People struct {
	name string //姓名
	age  int    //体重.单位斤
}

func (p *People) SetName(name string) {
	p.name = name
}
func (p *People) GetName() string {
	return p.name
}

func (p *People) SetAge(age int) {
	p.age = age
}

func (p *People) GetAge() int {
	return p.age
}
```

* 封装业务就是根据自己的需求提取代码,使用Go语言标准库中的函数过程就属性封装业务(代码)

# 一.多态

* 多态:同一件事情由于条件不同产生的结果不同
* 由于Go语言中结构体不能相互转换😂,所以没有结构体(父子结构体)的多态,只有基于接口的多态.这也符合Go语言对面向对象的诠释
* 多态在代码层面最常见的一种方式是`接口当作方法参数`


# 二.代码示例

* 结构体实现了接口的全部方法,就认为`结构体属于接口类型`,这是可以把结构体变量赋值给接口变量
* 重写接口时接收者为`Type`和`*Type`的区别
  * `*Type`可以调用`*Type`和`Type`作为接收者的方法.所以只要接口中多个方法中至少出现一个使用`*Type`作为接收者进行重写的方法,就必须把结构体指针赋值给接口变量,否则编译报错
  * `Type`只能调用`Type`作为接收者的方法

```go
type Live interface {
	run()
	eat()
}
type People struct {
	name string
}

func (p *People) run() {
	fmt.Println(p.name, "正在跑步")
}
func (p People) eat() {
	fmt.Println(p.name, "在吃饭")
}

func main() {
	//重写接口时
	var run Live = &People{"张三"}
	run.run()
	run.eat()
}
```

* 既然接口可以接收实现接口所有方法的结构体变量,接口也就可以作为方法(函数)参数

```go
type Live interface {
	run()
}
type People struct{}
type Animate struct{}

func (p *People) run() {
	fmt.Println("人在跑")
}
func (a *Animate) run() {
	fmt.Println("动物在跑")
}

func sport(live Live) {
	fmt.Println(live.run)
}

func main() {
	peo := &People{}
	peo.run() //输出:人在跑
	ani := &Animate{}
	ani.run() //输出:动物在跑
}
```


# 一. 接口

* 接口解释:接口是一组行为规范的定义.
* 接口中只能有方法声明,方法只能有名称、参数、返回值,不能有方法体
* 每个接口中可以有多个方法声明,结构体把接口中 **所有** 方法都重写后,结构体就属于接口类型
* Go语言中接口和结构体之间的关系是传统面向对象中is-like-a的关系
* 定义接口类型关键字是interface

```go
type 接口名 interface{
  方法名(参数列表) 返回值列表
}
```

* 接口可以`继承`接口,且Go语言推荐把接口中方法`拆分`成多个接口

# 二.代码示例

* 接口中声明完方法,结构体重写接口中方法后,编译器认为结构体实现了接口
  * 重写的方法要求必须和接口中方法名称、方法参数(参数名称可以不同)、返回值列表`完全相同`

```go
type People struct {
	name string
	age  int
}

type Live interface {
	run(run int)
}

func (p *People) run(run int) {
	fmt.Println(p.name, "正在跑步,跑了,", run, "米")
}

func main() {
	peo := People{"张三", 17}
	peo.run(100)
}
```

* 如果接口中有多个方法声明,接口体必须重写接口中`全部方法`才任务结构体实现了接口

```go
type People struct {
	name string
	age  int
}

type Live interface {
	run(run int)
	eat()
}

func (p *People) run(run int) {
	fmt.Println(p.name, "正在跑步,跑了,", run, "米")
}
func (p *People) eat() {
	fmt.Println(p.name, "正在吃饭")
}

func main() {
	peo := People{"张三", 17}
	peo.run(100)
}
```

* 接口可以`继承`接口(组合),上面代码可以改写成下面代码

```go
type People struct {
	name string
	age  int
}

type Live interface {
	run(run int)
	Eat
}

type Eat interface {
	eat()
}

func (p *People) run(run int) {
	fmt.Println(p.name, "正在跑步,跑了,", run, "米")
}
func (p *People) eat() {
	fmt.Println(p.name, "正在吃饭")
}

func main() {
	peo := People{"张三", 17}
	peo.run(100)
}
```

# 一. 方法

* `方法`和`函数`语法比较像,区别是`函数属于包`,通过包调用函数,而`方法属于结构体`,类似于Java的方法属于类,通过`结构体变量`调用
* 默认是函数,隶属于包,所以需要添加标识.告诉编译器这个方法`属性哪个结构体`
* 调用方法时就把调用者赋值给接收者(下面的变量名就是接受者)

```go
func (变量名 结构体类型) 方法名(参数列表) 返回值列表{
  //方法体
}
```

* Go语言中已经有函数了,又添加了对方法的支持主要是保证Go语言是面向对象的.


# 二.代码示例

* 定义一个People类型结构体,在对People结构体定义个run()方法

```go
type People struct {
	Name string//姓名
	Weight	float64//体重.单位斤
}

func (p People) run(){
	fmt.Println(p.Name,"正在跑步")
}

func main() {
	peo:=People{"张三",17}
	peo.run()
}
```

* 如果设定需求,在每次跑步后体重都减少0.1斤.上面代码就需要修改了.因为结构体是值类型,修改方法中结构体变量p的值,主函数中peo的值不会改变,因为传递的是值副本.所以修改方法中结构体类型为结构体指针类型就可以完成设定需求

```go
type People struct {
	Name string//姓名
	Weight	float64//体重.单位斤
}

func (p *People) run(){
	fmt.Println(p.Name,"正在跑步,体重为:",p.Weight)//输出:张三 正在跑步,体重为: 17
	p.Weight-=0.1
}

func main() {
	peo:=&People{"张三",17}
	peo.run()
	fmt.Println(peo.Name,"跑完步后的体重是",peo.Weight)//输出:张三 跑完步后的体重是 16.9
}
```




## 结构体
* Go语言中的结构体和C++结构体有点类似,而Java或C#中类本质就是结构体
* 结构体是值类型
* 结构体定义语法
  * 结构体就是一种自定义类型
    ```
     type 结构体名称 struct{
      名称 类型//成员或属性
    }
    ```


* 结构体可以定义在函数内部或函数外部(与普通变量一样),定义位置影响到结构体的访问范围
* 如果结构体定义在函数外面,结构体名称首字母是否大写影响到结构体是否能跨包访问
* 如果结构体能`跨包`访问,属性`首字母是否大写`影响到属性是否`跨包访问`

```
type People struct {
	Name string
	Age  int
}
```
声明结构体变量

* 由于结构体是`值类型`,所以`声明后就会开辟内存空间`
* 所有成员为类型对应的初始值
```
var peo People
	fmt.Print(peo)//输出:{0 }
	fmt.Printf("%p",&peo)//会打印内存地址值
```

- 可以直接给结构体多个属性赋值,按照结构体中属性的顺序进行赋值,可以省略属性名称

- 明确指定给哪些属性赋值.可以都赋值,也可以只给其中一部分赋值


- 也可以通过结构体变量名称获取到属性进行赋值或查看

- 双等(==)判断结构体中内容是否相等

## 结构题指针

* 由于结构体是`值类型`,在方法传递时希望传递`结构体地址`,可以使用时`结构体指针`完成
* 可以结合`new(T)`函数创建结构体指针



* 如果不想使用new(T)函数,可以直接声明结构体指针并赋值



# 二.判断

* 结构体指针比较的是地址
* (*结构体指针)取出地址中对应的值

```go
	p1 := People{"smallming", 17}
	p2 := People{"smallming", 17}
	fmt.Printf("%p %p\n", &p1, &p2) //输出地址不同
	fmt.Println(p1 == p2)           //输出:true

	p3 := new(People)
	p3 = &People{"smallming", 17}
	//结构体变量不能和指针比较,使用*指针取出地址中值
	fmt.Println(p1 == *p3) //输出:true

	p4 := &People{"smallming", 17}
	//指针比较的是地址
	fmt.Println(p3 == p4) //输出:false
```


# 一. 继承

* 按照传统面向对象思想,继承就是把同一类事物提出共同点为父类,让子类可以复用父类的可访问性内容.
* 继承有多种实现方式
  * 通过关键字继承,`强耦合`实现方式
  * 组合式继承,`松耦合`继承方式
* Go语言中的继承是通过组合实现

# 二.匿名属性

* 在Go语言中支持匿名属性(结构体中属性名字),但是每个最多只能存在匿名属性.编译器认为`类型`就是`属性名`,我们在使用时就把`类型当作属性名`进行使用

```go
type People struct {
	string
	int
}

func main() {
	p:=People{"xiaopang",17}
	fmt.Println(p.string,p.int)
}
```

# 三.结构体之间的关系

* 传统面向对象中类与类之间的关系
  * 继承:is-a,强耦合性,一般认为类与类之间具有强关系
  * 实现:like-a,接口和实现类之间的关系
  * 依赖:use-a,具有偶然性的、临时性的、非常弱的，但是B类的变化会影响到A,一般作为方法参数
  * 关联:has-a一种强依赖关系，比如我和我的朋友；这种关系比依赖更强、不存在依赖关系的偶然性、关系也不是临时性的，一般是长期性的，而且双方的关系一般是平等的、关联可以是单向、双向的
  * 聚合:has-a,整体与部分、拥有的关系
  * 组合:contains-a,他体现的是一种contains-a的关系，这种关系比聚合更强，也称为强聚合；他同样体现整体与部分间的关系，但此时整体与部分是不可分的，整体的生命周期结束也就意味着部分的生命周期结束
  * 组合>聚合>关联>依赖
* Go语言中标准的组合关系

```go
type People struct {
	name string
	age  int
}

type Teacher struct {
	peo       People
	classroom string //班级
}

func main() {
	teacher := Teacher{People{"smallming", 17}, "302教室"}
	//必须通过包含的变量名调用另一个结构体中内容
	fmt.Println(teacher.classroom, teacher.peo.age, teacher.peo.name)
}
```


# 四. 使用匿名属性完成Go语言中的继承

* Go语言中的继承很好实现,把另一个结构体类型当作另一个结构体的属性,可以直接调用另一个结构体中的内容
* 因为Go语言中结构体不能相互转换,所以不能把子结构体变量赋值给父结构体变量

```go
type People struct {
	name string
	age  int
}

type Teacher struct {
	People
	classroom string //班级
}

func main() {
	teacher := Teacher{People{"xiaopang", 17}, "302教室"}
	fmt.Println(teacher.classroom, teacher.age, teacher.name)
}
```


