package main 

import ("fmt")

func work(ch chan int){
	for i := 1; i < 10; i++ {
		fmt.Println("Hey there")
	}
	ch <- 2
}

func main(){
	var ch chan int 
	ch <- 1
	go work(ch) 
	fmt.Println(ch)
	close(ch)
}