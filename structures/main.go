package main

import (
	"fmt"
	"time"
)

type University struct {
	age       int
	yearBased int
	infoShort string
	infoLong  string
}

type Student struct {
	firstName  string
	lastName   string
	university University
}

type Professor struct {
	firstName string
	lastName  string
	age       int
	greatWork string
	//papers     map[string]string // Adding this field will make the struct incomparable
	University // Adding all University struct fields into Professor struct
}

func main() {
	student := Student{
		firstName: "John",
		lastName:  "Johnson",
		university: University{
			yearBased: 1890,
			infoShort: "Do not attend this university",
			infoLong:  "No useful information will be given to you",
		},
	}

	fmt.Println("infoShort University:", student.university.infoShort)
	fmt.Println("Year based University:", student.university.yearBased)

	currentTime := time.Now()

	professor := Professor{
		firstName: "Hershel",
		lastName:  "Jackson",
		age:       125,
		greatWork: "Ultimate Assembler programming",
		University: University{
			age:       currentTime.Year() - 1734,
			yearBased: 1734,
			infoShort: "Very bad university",
		},
	}

	fmt.Println("First Name:", professor.firstName)
	fmt.Println("Year based:", professor.yearBased)
	fmt.Println("Info short:", professor.infoShort)
	fmt.Println("Professor age:", professor.age)
	fmt.Println("University age:", professor.University.age)

	// Comparing objects

	firstProfessor := Professor{
		firstName: "Ivan",
	}
	secondProfessor := Professor{}

	fmt.Println("Professors are identical:", firstProfessor == secondProfessor)
}
