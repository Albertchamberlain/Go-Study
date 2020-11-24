# 一. map

* map以`散列表`方式存储键值对集合

* map中每个元素都是`键值对`

```
map[key] Value    写类型
```

* key是操作map的`唯一标准`.可以通过key对map中元素进行`增加/删除/修改/查看`
* key是`唯一`的,添加重复的key会覆盖之前的元素.
* map是值类型,`只声明时为空指针(nil)`

```go
	var m map[string]int
	fmt.Println(m == nil) //输出:true
	fmt.Printf("%p", m)   //输出:0x0
```

* map读写数据时并不是`并发安全`的,可以结合RWMutex保证并发安全(RWMutex在后面讲解),java提供concurrenthashmap解决方案

# 二.实例化map的几种方式

* 使用`make函数`实例化一个没有初始值的map

```go
	m:= make(map[string]string)
	fmt.Println(m==nil)//输出:false
	fmt.Printf("%p", m)//输出:内存地址
```

* 可以在声明map时直接给map赋初始值.注意初始值在一行和在多行写时的语法区别
  * map中元素键值对语法满足: key:value
  * key和value的类型必须和map[key]value类型严格对应

```go
	m := map[string]string{"name": "smallming", "address": "北京海淀"}
	m1 := map[string]string{
		"name":     "smallming",
		"addresss": "北京海淀",
	}
	fmt.Println(m, m1)
```

# 三.操作map中的元素

* 使用key判断,如果key不存在向map中新增数据,如果key存在会覆盖map中元素

```go
	m := make(map[string]int)
	m["money"] = 5
	fmt.Println(m) //输出:map[money:5]
	m["money"] = 6
	fmt.Println(m) //map[money:6]
```

* Go语言标准库中提供了对map元素删除的函数,使用`顶层delete()`即可完成删除
  * 如果key`存在执行`删除元素
  * 如果`key不存在,map中内容不变`,也不会有错误

```go
	m := make(map[string]int)
	m["money"] = 5
	delete(m, "没有的key")
	fmt.Println(m) //输出:map[money:5]
	delete(m, "money")
	fmt.Println(m) //输出:map[]
```

* 获取map中指定key对应的值
  * 使用:map变量[key]获取key对应的值
  * 如果`key不存在返回map[key]Value中Value类型的默认值`.例如:Value是string类型就返回`""`
  * 返回值可以是一个,也可以是两个.
    * 一个表示key对应的值
    * 两个分别表示:key对应的值和这个key是否存在

```go
	m := map[string]string{"name": "smallming", "address": "北京海淀"}
	fmt.Println(m["name"]) //输出:smallming
	fmt.Println(m["age"])  //输出:空字符串
	value, ok := m["age"]
	fmt.Println(value, ok) //输出:空字符串 false
```

* 如果希望把map中所有元素都`遍历`,可以使用`for结合range`实现

```go
	m := map[string]string{"name": "smallming", "address": "北京海淀"}
	//range遍历map时返回值分别表示key和value
	for key, value := range m {
		fmt.Println(key, value)
	}
```