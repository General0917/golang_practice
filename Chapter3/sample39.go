package main

import "fmt"

func insertSort(ary []int) {
	for i := 1; i < len(ary); i++ {
		tmp := ary[i]
		j := i - 1
		for ; j >= 0 && tmp < ary[j]; j-- {
			ary[j+1] = ary[j]
		}
		ary[j+1] = tmp
	}
}

func main() {
	a := []int{5, 6, 4, 7, 3, 8, 2, 9, 1, 0}
	b := []int{3, 1, 7, 6, 9, 2, 8, 0, 4, 5}
	c := []int{2, 5, 4, 1, 8, 6, 9, 3, 0, 7}
	insertSort(a)
	insertSort(b)
	insertSort(c)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
