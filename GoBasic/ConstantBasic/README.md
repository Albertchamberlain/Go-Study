* 常量关键字const
* 常量定义完可以不使用 `与变量不同`

# 一. 常量定义

* 定义常量时如果不是必须指定特定类型,可以省略类型,使用默认类型.且数值类型常量(不定义类型)可以直接进行运算
* 常量的值可以是表达式,但是不允许出现变量

```go
func main() {
	const a string = "smallming"
	const b = 123
	const c = 3*2 + 5//不要指定类型
	const d = 1.5//不要指定类型

	fmt.Printf("%T %T",c,d)//int float

	fmt.Println(c+d)//12.5

	//下面这种方式是错误的
	i := 3
	const e = i*2 + 5 //const initializer i * 2 + 5 is not a constant
}
```

* 当定义多个常量时官方推荐的方式

```go
	const (
		a = 1
		b = 2
		c = true
	)
```

* 定义多常量时后一个常量如果没有赋值,与前一个常量值相同.
* 第一个常量必须赋值

```go
func main() {
	const (
		a = 1
		b
		c
	)
	fmt.Println(a,b,c)//输出:1 1 1
}
```

# 二.常量生成器

* 当一组常量都是数值类型,可以使用常量生成器iota指定这组常量按照特定规则变化
* iota起始值为0,每次增加1

```go
func main() {
	const (
		a = iota
		b 
		c 
	)
	fmt.Println(a, b, c) //输出: 0 1 2

	const (
		d = iota << 1
		e 
		f 
	)
	fmt.Println(d, e, f) //输出:0 2 4
}
```

* 无论是否使用iota,一组常量中每个的iota值是固定的,iota按照顺序自增1
* 每组iota之间无影响

```go
func main() {
	const (
		a = 5    //iota=0
		b = 3    //iota=1
		c = iota //iota=2
		d        //iota=3
	)
	fmt.Println(a, b, c, d) //输出5 3 2 3

	const (
		e = iota //iota=0
		f        //iota=1
		g = 10   //iota=2
		h        //iota=3
		i = iota //iota=4
		j        //iota=5
	)
	fmt.Println(e, f, g, h, i, j) // 0 1 10 10 4 5
}
```

