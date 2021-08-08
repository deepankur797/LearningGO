package main

import (
	"fmt"
)

func bubbleSort(arr []int) (res []int){
	swapped  := true
	for swapped {
		swapped =false
		for i:=1; i<len(arr);i++{
			if arr[i-1] >arr[i]{
				arr[i-1]=arr[i]+arr[i-1]
				arr[i]=arr[i-1]-arr[i]
				arr[i-1]=arr[i-1]-arr[i]
				swapped =true
			}
		}
	}

return arr
}

func main(){
	fmt.Printf("Hello Bubble Sort\n")
	arr := make([]int,0)
	for i:=0;i<10;i++{
		arr=append(arr,10-i)
	}
	fmt.Printf("Before Sorting \n")
	fmt.Println(arr)
	arr = bubbleSort(arr)
	fmt.Printf("After Sorting \n")
	fmt.Println(arr)
}

