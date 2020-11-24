package main

import "fmt"

type live interface {
	run()
}
type people struct{}
type animate struct{}

func (p *people) run() {
	fmt.Println("人在跑")
}
func (a *animate) run() {
	fmt.Println("动物在跑")
}

func sport(live live) {
	fmt.Println(live.run)
}

func main() {
	peo := &people{}
	peo.run() //输出:人在跑
	ani := &animate{}
	ani.run() //输出:动物在跑
}
