package main

import "fmt"

func divMod(x, y int) (int, int) {
	return x / y, x % y
}
func main() {
	p, q := divMod(10, 3)
	fmt.Println(p)
	fmt.Println(q)
}
