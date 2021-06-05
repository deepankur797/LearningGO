package main
import "fmt"

func main() {

	s:= make([]string,3)
	fmt.Println("emp:",s)
	s[0]="A"
	s[1]="B"
	s[2]="C"
	fmt.Println("set:",s)
	fmt.Println("get:",s[2])
	fmt.Println("len:",len(s))
	fmt.Println("cap:",cap(s))

	s=append(s,"D")
	s=append(s,"E","F")
	fmt.Println("apd:",s)

	var c = make([]string, len(s))
	copy(c,s)
	fmt.Println("cpy:",c)
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
        	innerLen := i + 1
        	twoD[i] = make([]int, innerLen)
        	for j := 0; j < innerLen; j++ {
            		twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ", twoD)
}
