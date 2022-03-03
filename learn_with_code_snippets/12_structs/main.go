package main

import (
	"fmt"
	"strconv"
)

//Defining a struct
type Person struct {
	first string
	last  string
	age   int
}

// Greeting a person using a struct field name
// Greeting Method (Value Receiver)

func (p Person) greet() {
	fmt.Println("Hello", p.first, p.last, "you are", strconv.Itoa(p.age), "years old")
}

func main() {
	fmt.Println("Structs")

	//Creating a struct
	p1 := Person{
		first: "James",
		last:  "Bond",
		age:   32,
	}

	//Calling a greeting func on a struct
	p1.greet()

	//Accessing struct fields
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	fmt.Println(p1.age)

	//Accessing struct fields using a dot
	fmt.Println(p1.first, p1.last, p1.age)

	//Accessing struct fields using a dot
	fmt.Println(p1.first, p1.last, p1.age)
}
