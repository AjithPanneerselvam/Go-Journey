package main 

import (
	"fmt"
	"time"
)


func boring(message string, ch chan<- int){
	for i := 0; ; i++{
		ch <- fmt.Sprintf("%s %d", message, i)
		time.Sleep(time.Duration(rand.Intn(1e3)))
	}
}

func main(){
	var ch chan int 
	ch = make(chan int)
	go boring("boring")
	for i := 1; i <= 5; i++{
		fmt.Println("You say: ", <-ch)
	}
	fmt.Println("I'm leaving! You are so boring")
}