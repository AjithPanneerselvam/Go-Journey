package main 

import "fmt"

type sampleStruct struct {
	name string 
	marks *[]int 
}

func main() {
	sample := sampleStruct{"ajith", &([]int{1,2,3,4,5})}
	fmt.Printf("%T\n", (*sample.marks))
	
	for i := 0; i < len(*sample.marks); i++ {
		fmt.Println((*sample.marks)[i])
	}
}