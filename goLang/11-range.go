package main

import "fmt"

func main() {

	nums := []int{2, 3, 4}

	sum := 0

	for _, num := range nums {
		sum += num
	}
	fmt.Println(sum)

	var kvs = make(map[string]string)

	kvs["a"] = "apple"
	kvs["b"] = "banana"

	for k, v := range kvs {
		fmt.Printf("%s-> %s\n", k, v)
	}

	for k := range kvs {
		fmt.Println(k)
	}

	for _, v := range kvs {
		fmt.Println(v)
	}

	for i, c := range "go is an amazing language" {
		fmt.Println(i, c)
	}
	z := "string"
	fmt.Println(z)
}
