package main

import "fmt"

func main() {
	x := 1
	{
		y := 2
		{
			z := 3
			{
				fmt.Println(x)
				fmt.Println(y)
				fmt.Println(z)
			}
		}
		fmt.Println(x)
		fmt.Println(y)
		//fmt.Println(z) //zは範囲外（コンパイルエラー）
	}
	fmt.Println(x)
	//fmt.Println(y) //yは範囲外（コンパイルエラー）
	//fmt.Pirntln(z) //zは範囲外（コンパイルエラー）
}
