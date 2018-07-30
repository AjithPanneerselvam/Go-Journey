package main 

import (
	"fmt"
	"time"
)


func working(message string){
	fmt.Println(message)
}


func main() {
	go working("Working")
	fmt.Println("I'm listening")
	time.Sleep(2 * time.Second)
	fmt.Println("I'm done")
}