package main

import "fmt"

type People struct {
	name string
	age  int
}
type Eat interface {
	eat()
}
type Live interface {
	run(run int)
	Eat
}

func (p *People) run(run int) {
	fmt.Println(p.name, "正在跑步,跑了,", run, "米")
}
func (p *People) eat() {
	fmt.Println(p.name, "正在吃饭")
}

func test() {
	peo := People{"张三", 17}
	peo.run(100)
}
