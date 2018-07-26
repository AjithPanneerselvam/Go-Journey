package main 

import(
	"fmt"
)

type employee struct{
	id int
	name string 
	role string
}

func (e *employee) displayInfo(){
	fmt.Println("ID: ", e.id, "\nName: ", e.name, "\nRole: ", e.role)
}


func main(){
	emp := &employee{id: 1, name: "Ajith", role: "SDE Intern"}
	emp.displayInfo()
}