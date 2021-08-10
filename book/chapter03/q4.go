package main

import (
	"fmt"
)

func rec(num int) (res []int){
	s := make([]int, num)
	for i:=0;i<num;i++{
		if i==0 || i==1{
			s[i] = 1
		} else {
			s[i] =s[i-1]+s[i-2]
		}
	}
	res = s
	return res
}
func main(){
	arr := rec(11)
	fmt.Println(arr)
}
