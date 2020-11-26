package main

import (
	"fmt"
	"regexp"
)

func main() {
	r := regexp.MustCompile(`\d[a-zA-Z]`)
	fmt.Println(r.MatchString("5A1"))
	fmt.Println(r.FindAllString("56A6B7C", -1))
	fmt.Println(r.Split("12345qwert", -1))
	fmt.Println(r.ReplaceAllString("12345qwert", "替换了"))
}
