package main

import (
	"fmt"
)
func plusX(x int) func(int) int{
	return func(y int) int {return y+x}
}

func main(){
	p :=plusX(3)
	fmt.Printf("%v\n",p(2))
}
