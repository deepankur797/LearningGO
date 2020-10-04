// In go variables are explicitly declared and used by the complier. 

package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	var  b, c int = 1,2
	fmt.Println(b,c)

	var d =true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple" //short hand for var f = "apple"

	fmt.Println(f)

	var g string
	fmt.Println(g)
}
