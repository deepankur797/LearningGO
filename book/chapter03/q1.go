package main

import (
	"fmt"
)

	//fmt.Println("Hello World")
func average (nums ...float64) (result float64){
		sum := 0.0
		//for k,v := range nums{
		//	sum =sum+nums[k]
		//}
		for i:=0; i<len(nums); i++{
			sum=sum + nums[i]
		}
		result= sum/float64(len(nums))
		return result
}

func main(){
		//a := []float64{0.0,1.0,2.0,3.0}
		x := average(0.0,1.0,2.0,3.0,4.0)
		fmt.Println(x)
}
