package main 

import ("fmt")

func work(){
	for i := 1; i < 10; i++ {
		fmt.Println("Hey there")
	}
}

func main(){
	go work() 

	a := 5
	if a == 5 {
		fmt.Println(a)
	}

	var val int
	fmt.Scanf("%d", &val)

	slices = make([]int, val)
	for i := 0; i < val; i++ {
		fmt.Println(slices[i])
	}	
}