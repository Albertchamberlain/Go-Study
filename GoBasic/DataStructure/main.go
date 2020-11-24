package main

import (
	"container/list"
	"container/ring"
	"fmt"
)

func main() {
	var slice []string
	// var array [5] string

	fmt.Println(slice == nil)
	fmt.Println("%p", slice)

	names := []string{"algorithm", "算法"}
	fmt.Println(names)

	namess := []string{"algorithm", "算法"}
	names1 := namess
	names1[0] = "法"
	fmt.Println(namess, names1)
	fmt.Printf("%p %p", namess, names1) //地址相同

	s := make([]int, 0)
	fmt.Println(s == nil)

	fmt.Println("%f\n", s)
	fmt.Println("%p\n", s)

	fmt.Println(len(s), cap(s))

	s1 := make([]int, 0, 3)
	fmt.Println(len(s1), cap(s1))

	s4 := make([]string, 0)
	fmt.Println(len(s4), cap(s4))
	s4 = append(s4, "algorithm", "math")
	fmt.Print(len(s4), cap(s4))

	s6 := make([]string, 0, 3)
	s7 := make([]string, 0, 3)

	s6 = append(s6, s7...)

	fmt.Println(s6, len(s6), cap(s6)) //值不够，没有触发扩容
	fmt.Println(s7)

	num := []int{0, 1, 2, 3, 4, 5, 6}
	//要删除脚标为n的元素
	n := 2
	num1 := num[0:n] //左闭右开
	num1 = append(num1, num[n+1:]...)
	fmt.Println(num1)

	s5 := []int{1, 2}
	s2 := []int{3, 4, 5, 6}
	fmt.Println(cap(s2))
	copy(s2, s5)
	fmt.Println(s5)
	fmt.Println(cap(s2))
	fmt.Print(s2)

	s9 := []int{1, 2}
	s8 := []int{3, 4, 5, 6}
	copy(s9, s8)
	fmt.Println(s9) //输出:[3 4] ,对相应的角标进行覆盖，其他的舍弃
	fmt.Println(s8) //输出:[3 4 5 6]

	s11 := []int{1, 2}
	s21 := []int{3, 4, 5, 6}
	copy(s11, s21[1:])
	fmt.Println(s11) //输出:[4 5]
	fmt.Println(s21) //输出:[3 4 5 6]

	ss := []int{1, 2, 3, 4, 5, 6, 7}
	nn := 2 //要删除元素的索引
	newSlice := make([]int, nn)
	copy(newSlice, ss[0:n])
	newSlice = append(newSlice, ss[n+1:]...)
	fmt.Println(ss)       //原切片不变
	fmt.Println(newSlice) //删除指定元素后的切片

	var m map[string]int
	fmt.Println(m == nil) //输出:true
	fmt.Printf("%p", m)   //输出:0x0

	m2 := make(map[string]string)
	fmt.Println(m2 == nil) //输出:false
	fmt.Printf("%p", m2)   //输出:内存地址

	m3 := map[string]string{"name": "xiaopang", "address": "火星"}
	m4 := map[string]string{
		"name":     "xiaopang",
		"addresss": "火星",
	}
	fmt.Println(m3, m4)

	m12 := make(map[string]int)
	m12["ok"] = 5
	fmt.Println(m12)
	m12["ok"] = 6
	fmt.Print(m12)

	m13 := make(map[string]int)
	m13["okk"] = 5
	delete(m13, "没有的key")
	fmt.Println(m13)
	delete(m13, "okk")
	fmt.Println(m13)

	m15 := map[string]string{"name": "xiaozhu", "address": "火星"}
	fmt.Println(m15["name"]) //输出:xiaozhu
	fmt.Println(m15["age"])  //输出:空字符串
	value, ok := m15["age"]
	fmt.Println(value, ok) //输出:空字符串 false

	m15["age"] = "14"
	keyy, valuee := m15["age"]

	fmt.Println(keyy, valuee, m15)

	m17 := map[string]string{"name": "xiaoda", "address": "水星"}
	//range遍历map时返回值分别表示key和value
	for key, value := range m17 {
		fmt.Println(key, value)
	}

	mylist := list.New()
	fmt.Println(mylist)
	fmt.Println(mylist.Len())
	fmt.Printf("%p", mylist)

	//添加到最后,List["a"]
	mylist.PushBack("a")
	//添加到最前面,List["b","a"]
	mylist.PushFront("b")
	//向第一个元素后面添加元素,List["b","c","a"]
	mylist.InsertAfter("c", mylist.Front())
	//向最后一个元素前面添加元素,List["b","c","d","a"]
	mylist.InsertBefore("d", mylist.Back())

	fmt.Println(mylist.Back().Value)  //最后一个元素的值
	fmt.Println(mylist.Front().Value) //第一个元素的值

	//只能从头向后找,或从后往前找,获取元素内容
	n5 := 5
	var curr *list.Element
	if n5 > 0 && n5 <= mylist.Len() {
		if n5 == 1 {
			curr = mylist.Front()
		} else if n == mylist.Len() {
			curr = mylist.Back()
		} else {
			curr = mylist.Front()
			for i := 1; i < n5; i++ {
				curr = curr.Next()
			}
		}
	} else {
		fmt.Println("n的数值不对")
	}
	//遍历所有值 很快啊
	for e := mylist.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	mylist.MoveToBack(mylist.Front())                //把第一个移动到最后
	mylist.MoveToFront(mylist.Back())                //把最后一个移动最前面
	mylist.MoveAfter(mylist.Front(), mylist.Back())  //把第一个参数元素,移动到第二个参数元素后面
	mylist.MoveBefore(mylist.Front(), mylist.Back()) //把第一个参数元素,移动到第二个参数元素前面

	mylist.Remove(mylist.Front()) //  移除第一个节点

	//r代表第一个元素
	r := ring.New(3)
	for i := 0; i < r.Len(); i++ {
		r.Move(i).Value = i
	}
	r.Do(func(i interface{}) {
		fmt.Println(i)
	})

	fmt.Println(r.Next().Value)               //输出:1
	fmt.Println(r.Next().Next().Value)        //输出:2
	fmt.Println(r.Next().Next().Next().Value) //输出:0
	fmt.Println(r.Move(-1).Value)             //输出:2
	fmt.Println(r.Prev().Value)               //输出:2

	s99 := ring.New(1)
	s99.Value = 13
	//r是哪个元素,就把新的链表添加到哪个元素后面
	r.Link(s99) //把元素link到指定位置
	r.Do(func(i interface{}) {
		fmt.Print(i, " ")
	})
	fmt.Println("")
	//从r元素向后,n/r.Len()个元素被删除,当前元素和前面的保留
	r.Unlink(1) //从链表上删除指定元素
	r.Do(func(i interface{}) {
		fmt.Print(i, " ")
	})

}
