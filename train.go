package main

import (
	"fmt"
	"strings"
)

func main() {
	var a string
	fmt.Scan(&a)
	s := strings.Split(a, "")
	m := strings.Join(s, "*")
	fmt.Println(m)
}
