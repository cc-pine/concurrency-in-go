package main

import (
	"fmt"
)

const (
	a = "hello"
)

type myString string

func main() {
	var s0 string
	var s1 myString
	s0 = a
	s1 = a
	fmt.Println(s0)
	fmt.Println(s1)
}
