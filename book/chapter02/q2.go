package main

import (
	"fmt"
)

func main() {
	var arr [10]int
	sum := 0
	var avg float64
	for i:=0; i<10;i++{
		arr[i]=i
		sum=sum+arr[i]
	}
	avg = float64(sum)/float64(len(arr))
	fmt.Println(avg)
}
