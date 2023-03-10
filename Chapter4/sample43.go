package main

import (
	"fmt"
)

func pow(x float64, n int) float64 {
	v := 1.0
	for ; n > 0; n-- {
		v *= x
	}
	return v
}

func main() {
	fmt.Println(pow(2, 9))
	fmt.Println(pow(2, 16))
	fmt.Println(pow(2, 32))
	fmt.Println(pow(2, 64))
}
