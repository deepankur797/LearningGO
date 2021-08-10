package main

import (
	"fmt"
)

func mapper(f func(int) int, list []int) []int{
	s := make([]int,len(list))
	for i:=0;i<len(s);i++{
		s[i] = f(list[i])
	}
return s
}

func square(num int) int{
	return num*num
}


func main(){
	arr := [5]int{1,2,3,4,5}
	x := arr[:]
	//f := func(i int) int{
        //return i*i}

	x = mapper(square, x)
	fmt.Println(x)
}
