package main

import "fmt"

type Employee struct {
	name     string
	position string
	salary   int
	currency string
}

// func (s Struct) MethodName(parameters type) type {}

func (e Employee) DisplayInfo() {
	fmt.Println("Name:", e.name)
	fmt.Println("Position:", e.position)
	fmt.Printf("Salary: %d %s\n", e.salary, e.currency)
}

func main() {
	employee := Employee{
		name:     "Bob",
		position: "Senior Golang developer",
		salary:   3000,
		currency: "USD",
	}

	// Call the method
	employee.DisplayInfo()
}
