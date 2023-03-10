package main

import "fmt"

func main() {
	var s string = "hello, world"
	var p *string = &s
	var a = []int{1, 2, 3, 4, 5, 6, 7, 8}
	var q *[]int = &a
	q1, q2 := &a[0], &a[1]
	fmt.Println(s)
	fmt.Println(p)
	fmt.Println(*p)
	*p = "oops!"
	fmt.Println(s)
	fmt.Println(*p)
	fmt.Println(q)
	fmt.Println(*q)
	fmt.Println(*q1)
	fmt.Println(*q2)
	*q = []int{10, 20, 30, 40}
	fmt.Println(q)
	fmt.Println(*q)
}
