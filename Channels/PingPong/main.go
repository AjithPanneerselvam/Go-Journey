package main 

import (
	"fmt"
)

func Ping(ping chan<- string, message string){
	ping<- message
}

func PingPong(ping <-chan string, pong chan<- string){
	message :=  <-ping 
	pong<- message
}

func main(){
	ping := make(chan string, 1)
	pong := make(chan string, 1)
	Ping(ping, "Ping Pong")
	PingPong(ping, pong)
	fmt.Println(<-pong)
}