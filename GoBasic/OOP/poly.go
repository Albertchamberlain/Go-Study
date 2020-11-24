package main

import "fmt"

type Livee interface {
	run()
	eat()
}
type Peoplee struct {
	name string
}

func (p *Peoplee) run() {
	fmt.Println(p.name, "正在跑步")
}
func (p Peoplee) eat() {
	fmt.Println(p.name, "在吃饭")
}

func main() {
	//重写接口时
	var run Livee = &Peoplee{"张三"}
	run.run()
	run.eat()
}
