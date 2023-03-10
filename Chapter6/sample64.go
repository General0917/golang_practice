package main

import "fmt"

func swap(x, y *int) {
	tmp := *x
	*x = *y
	*y = tmp
}

func timesArray(n int, ary *[8]int) {
	for i := 0; i < len(*ary); i++ {
		ary[i] *= n
	}
}

func main() {
	var a int = 10
	var b int = 20
	var c [8]int = [8]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	swap(&a, &b)
	timesArray(10, &c)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
