package main

import (
	"fmt"
)
func max(x []int){
	sum := x[0]
	for i:=1;i<len(x);i++{
	if  x[i]>sum{
		sum =x[i]
	}
	}
	fmt.Println(sum)
}
func main(){
	x := [5]int{22,33,11,44,12}
	max(x[:])
}
