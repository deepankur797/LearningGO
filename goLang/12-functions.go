package main

import "fmt"

func plus(a, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}

func main() {
	res := plus(1, 2)
	fmt.Printf("%d \n", res)

	res = plusPlus(1, 2, 4)
	fmt.Println(res)
}
