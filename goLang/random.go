package main

import "fmt"

func main(){
	fmt.Println("Hello Deepankur!!")
	fmt.Printf("Let's start with some Practice")
	fmt.Println()
	var(
		a int
		b bool
		c string
	)
	a = 17
	b = true
	c = "This is A string"
	fmt.Printf("A is equal to %X, b is %v, c is %s\n", a,b,c)
	// Runes for traversing characters in the string

	list := []rune(c) // THis will generate an array of Ascii values for the character
	fmt.Printf("%v\n",list)
	fmt.Printf(string(list[8]))
	fmt.Println()
	/* The three types of for loops*/
	sum := 0
	for i:=0;i<10;i++ {
		sum=sum+i
	}
	fmt.Printf("%v\n",sum)
	sum =0

	for sum<20{
		sum=sum+5
	}
	fmt.Printf("%v\n",sum)

	for {
		if sum>20{
			break
		}
		sum=sum+sum
	}
	fmt.Printf("%v\n",sum)
	/*
		Range operator is used for loops. It can loop over slices,arrays,strings,maps and channels

		It returns a key-value from the value on which it runs
	*/
	list2 :=[]string{"a","b","c","d","e","f"}
	for k,v:= range list2{
		fmt.Printf("character %d starts at byte position %s\n",k,v)
	}
	/*
		Arrays and slices
	*/
	var array [5]int
	for i:=0;i<len(array);i++{
		array[i]=i
	}
	list3 := [...]int{1,2,3,4,5,6}
	fmt.Println(array)
	fmt.Println(list3)
	// slices update the underlying array but are of variable lenght
	//s := make([]int,5)
	s := array[0:]
	s[0] =1
	fmt.Println(array)
	fmt.Println(cap(s))
	s = append(s,2)
	s = append(s,33)
	s = append(s,34)
	s = append(s,35)
	s = append(s,36)
	s = append(s,37)
	fmt.Println(cap(s))
	// maps
	
	monthDays := map[string]int{
		"a": 0,
		"b": 1,
		"c": 2,
		"d": 3,
		"e": 4,
		"f": 5,
	}
	for k,v := range monthDays{
		fmt.Printf("Key is %s,Value is %d\n",k,v)
	}
	fmt.Println(monthDays["a"])
}
