package main

import "fmt"

func mapcar(f func(int) int, ary []int) []int {
	buff := make([]int, len(ary))
	for i, v := range ary {
		buff[i] = f(v)
	}
	return buff
}

func removeIf(f func(int) bool, ary []int) []int {
	buff := make([]int, 0)
	for _, v := range ary {
		if !f(v) {
			buff = append(buff, v)
		}
	}
	return buff
}

func filter(f func(int) bool, ary []int) []int {
	buff := make([]int, 0)
	for _, v := range ary {
		if f(v) {
			buff = append(buff, v)
		}
	}
	return buff
}

func foldl(f func(int, int) int, a int, ary []int) int {
	for _, x := range ary {
		a = f(a, x)
	}
	return a
}

func foldr(f func(int, int) int, a int, ary []int) int {
	for i := len(ary) - 1; i >= 0; i-- {
		a = f(ary[i], a)
	}
	return a
}

func sumOfInteger(n, m int) int {
	a := 0
	for ; n <= m; n++ {
		a += n
	}
	return a
}

func square(x int) int {
	return x * x
}

func cube(x int) int {
	return x * x * x
}

func sumOfSquare(n, m int) int {
	a := 0
	for ; n <= m; n++ {
		a += square(n)
	}
	return a
}

func sumOfCube(n, m int) int {
	a := 0
	for ; n <= m; n++ {
		a += cube(n)
	}
	return a
}

func sumOf(f func(int) int, n, m int) int {
	a := 0
	for ; n <= m; n++ {
		a += f(n)
	}
	return a
}

func isEven(x int) bool {
	return x%2 == 0
}

func isOdd(x int) bool {
	return x%2 != 0
}

func add(x, y int) int {
	return x + y
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	b := mapcar(square, a)
	c := mapcar(cube, a)
	fmt.Println(b)
	fmt.Println(c)
	d := removeIf(isEven, a)
	e := removeIf(isOdd, a)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(foldl(add, 0, a))
	fmt.Println(foldr(add, 0, a))
}
