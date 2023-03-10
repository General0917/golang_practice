package main

import "fmt"

func gcdi(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	fmt.Println(gcdi(42, 30))
	fmt.Println(gcdi(15, 70))
}
