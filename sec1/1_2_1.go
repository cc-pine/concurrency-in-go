package main

import "fmt"
import "time"

func main() {
	var data int
	go func() {
		data++
		fmt.Println("goroutine", time.Now())
		} ()
	// fmt.Println(" ", time.Now())
	if data == 0 {
		fmt.Println("if block", time.Now())
		fmt.Printf("the value is %v.\n", data)
		fmt.Println("if block", time.Now())
	}
}

a = 0
defer fmt.Println(a)
a += 1
return