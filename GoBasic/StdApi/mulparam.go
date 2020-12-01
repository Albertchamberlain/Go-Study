package main

import (
	"fmt"
	"sort"
)

type ItemT2 struct {
	Grade int
	Age   int
	Name  string
}

type ItemsT2 []ItemT2

//重写三个方法
func (p ItemsT2) Len() int {
	return len(p)
}

func (p ItemsT2) Less(i, j int) bool {
	if p[i].Grade < p[j].Grade {
		return true
	}
	if p[i].Grade > p[j].Grade {
		return false
	}

	return p[i].Age > p[j].Age
}

func (p ItemsT2) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *ItemsT2) Sort() {
	sort.Sort(p)
}

// 主函数
func main() {
	tmp := ItemsT2{
		{
			Grade: 10,
			Name:  "A",
			Age:   28,
		},
		{
			Grade: 5,
			Name:  "B",
			Age:   57,
		},
		{
			Grade: 5,
			Name:  "C",
			Age:   43,
		},
		{
			Grade: 20,
			Name:  "D",
			Age:   16,
		},
		{
			Grade: 1,
			Name:  "E",
			Age:   31,
		},
	}
	tmp.Sort()
	fmt.Printf("%+v  ", tmp)
}
