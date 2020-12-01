package main

import (
	"fmt"

	"sort"
)

type Obj struct {
	Var      string `json:"var"`
	Mtime    int    `json:"mtime"`
	Orderval int    `json:"orderval"`
}

type List []Obj

func (p List) Len() int {

	return len(p)

}

func (p List) Less(i, j int) bool {

	if p[i].Mtime > p[j].Mtime {

		return true

	}

	if p[i].Mtime < p[j].Mtime {

		return false

	}

	return p[i].Orderval > p[j].Orderval

}

func (p List) Swap(i, j int) {

	p[i], p[j] = p[j], p[i]

}

func (p *List) Sort() {

	sort.Sort(p)

}

func main() {

	s := List{{"f", 1595144638, 4}, {"d", 1595144646, 2}, {"a", 1595144648, 8}, {"t", 1595144648, 5}, {"e", 1595144650, 3}}

	fmt.Println(s)

	s.Sort()

	fmt.Println(s)

}
