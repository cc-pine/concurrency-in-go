package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 4)
	go func() {
			for i := 0; i < 4; i++ {
			c <- i
		}
	} ()
	for i := 0; i < 2; i++ {
		var d int
		d = <- c
		fmt.Printf("%v", d)
	}
	close(c)
	d, e := <- c
	fmt.Printf("%v", d)
	fmt.Printf("%v", e)
	loop:
	for i := 0; i < count; i++ {
		
	}
}
// できない