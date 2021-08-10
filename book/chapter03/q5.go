package main

import (
	"fmt"
)
func pri(nums...int){
	for i:=0;i<len(nums);i++{
	fmt.Println(nums[i])
	}
}

func main(){
	pri(0,1,2,3,4,5,4,3,2,1,0)
}
