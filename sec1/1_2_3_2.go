package main

import ("fmt"
		"sync"
		"time"
)

func main() {
	var data int
	var memoryAccess sync.Mutex
	go func() {
		memoryAccess.Lock()
		data++
		memoryAccess.Unlock()
		} ()
	// time.Sleep(1*time.Second)
	memoryAccess.Lock()
	if data == 0 {
		fmt.Println("the value is 0.")
	} else {
		fmt.Printf("the value is %v.\n", data)
	}
	memoryAccess.Unlock()
}