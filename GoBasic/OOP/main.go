package main

import "fmt"

type People2 struct {
	Name   string  //姓名
	Weight float64 //体重.单位斤
}

type Teacher struct {
	People2
	classroom string
}

func (p2 People2) run() {
	fmt.Println(p2.Name, "正在跑步")
}

func main() {
	type People struct {
		Name string
		Age  int
	}

	var pe People
	fmt.Print(pe)
	fmt.Println("%p", &pe)

	var peo People
	peo = People{"hou", 20}
	fmt.Println(peo)

	peo = People{Age: 18, Name: "xiaopang"}
	fmt.Println(peo)

	var pp People
	pp = People{Age: 12, Name: "can"}
	fmt.Println(pp)

	var pep People //我更倾向于这种方式，比较有面向对象的感觉
	pep.Name = "xiaopang"
	pep.Age = 17
	fmt.Println(pep)
	fmt.Println(pep.Name)
	fmt.Println(pep.Age)

	p1 := People{"xiaopang", 17}
	p2 := People{"xiaopang", 17}
	fmt.Printf("%p %p\n", &p1, &p2) //输出地址不同 0xc0420484e0 0xc042048500
	fmt.Println(p1 == p2)           //输出:true

	peop := new(People)
	//因为结构体本质是值类型,所以创建结构体指针时已经开辟了内存空间
	fmt.Println(peop == nil) //输出:false
	//由于结构体中属性并不是指针类型,所以可以直接调用
	peop.Name = "xiaopang"
	fmt.Println(peop)
	peo1 := peop
	peo1.Name = "xiaopang"
	fmt.Println(peo1, peop)

	var peo2 *People

	//给结构体指针赋值
	peo2 = &People{"xiaopang", 17}
	/*
		上面代码使用短变量方式如下
		peo2:= &People{"xiaopang", 17}
	*/
	fmt.Println(peo2)

	peo4 := People2{"xiaopang", 23}
	fmt.Println(peo4)

	peo4.run()

	peo5 := &People2{"张三", 17}
	peo5.run()
	fmt.Println(peo5.Name, "跑完步后的体重是", peo5.Weight)

	teacher := Teacher{People2{Name: "xiaopang", Weight: 13}, "311"}
	fmt.Println(teacher.classroom, teacher.Weight, teacher.Name)

}

func (p *People2) run2() {
	fmt.Println(p.Name, "正在跑步,体重为:", p.Weight) //输出:张三 正在跑步,体重为: 17
	p.Weight -= 0.1
}
