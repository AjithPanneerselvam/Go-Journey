package main

import (
	"fmt"
	"Interfaces/pkgExample"
)

type artithmetic interface{
	increment()
	decrement()
}

type customInt int;

func(i *customInt) increment(){
	*i = *i + 1
}

func(i *customInt) decrement(){
	*i = *i - 1
}

func main(){
	var i  artithmetic;
	var value customInt = 2
	i = &value 
	i.increment()
	i.decrement()
	i.increment()
	i.increment()
	i.increment()
	fmt.Println(i)

	pkgExample.Hello()	
}