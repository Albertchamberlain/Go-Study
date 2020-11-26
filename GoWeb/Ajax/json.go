package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func main() {
	user := User{"xiaopang", 20}
	b, _ := json.Marshal(user)
	fmt.Println(string(b))
	u2 := new(User)
	json.Unmarshal(b, u2)
	fmt.Println(u2)
}
