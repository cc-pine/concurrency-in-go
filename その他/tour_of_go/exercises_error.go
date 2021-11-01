package main

import (
	"fmt"
)

type MyError struct {
	What float64
}

func (e MyError) Error() string {
    return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e.What))
}

func Sqrt(x float64) (float64, error) {
	if x >= 0 {
		zold := 1.
		znew := 0.

		for i:=0;; i++ {
			znew = zold - (zold*zold - x) / (2*zold)
			if (znew - zold) < 0.0000000001 && (zold - znew) < 0.0000000001 {
				return znew, nil
			} else {
				zold = znew
			}
		}
	} else {
		return 0, MyError{x}
	}
}


func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

