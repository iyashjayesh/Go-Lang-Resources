package main

import "fmt"

func main() {
	fmt.Println("Conditionals")

	// if else statement
	x := 10
	y := 20
	if x < y {
		fmt.Println("x is less than y")
	} else {
		fmt.Println("x is greater than y")
	}

	// if else if else statement
	name := "Yash"
	if name == "Yash" {
		fmt.Println("Hello Yash")
	} else if name == "Jhon" {
		fmt.Println("Hello Jhon")
	} else {
		fmt.Println("Hello Stranger")
	}

	// switch statement
	switch name {
	case "Yash":
		fmt.Println("Hello Yash")
	case "Jhon":
		fmt.Println("Hello Jhon")
	default:
		fmt.Println("Hello Stranger")
	}
}
