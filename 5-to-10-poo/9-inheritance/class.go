package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
}

func GetMessage(p Person) {
	fmt.Printf("%s with age %d\n", p.name, p.age)
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Name"
	ftEmployee.age = 2
	ftEmployee.id = 5
	fmt.Printf("%v\n", ftEmployee)
	//GetMessage(ftEmployee) //-> cannot use ftEmployee (type FullTimeEmployee) as type Person in argument to GetMessage: FullTimeEmployee does not implement Person (missing method name)
	GetMessage(ftEmployee.Person) //-> works because we pass the Person part of FullTimeEmployee
}
