package main

import(
	"fmt"
)

func main(){
	s := make([]int,0)
	for i:=0;i<100;i++{
		if i%3==0 && i%5==0{
			s=append(s,i)
			fmt.Println("FizzBuzz")
		}else if i%3==0{
			s=append(s,i)
			fmt.Println("Fizz")
		}else if i%5==0{
			s=append(s,i)
			fmt.Println("Buzz")
		}else{
			continue
		}
	}
	fmt.Println(s)
}
