package main

import (
	"fmt"
	"sort"
)

type ItemT struct {
	Grade int
	Age   int
	Name  string
}

type ItemsT []ItemT

//重写三个方法
func (p ItemsT) Len() int {
	return len(p)
}

func (p ItemsT) Less(i, j int) bool {
	return p[i].Grade < p[j].Grade
}

func (p ItemsT) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p *ItemsT) Sort() {
	sort.Sort(p)
}

// 主函数
func main() {
	tmp := ItemsT{
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
