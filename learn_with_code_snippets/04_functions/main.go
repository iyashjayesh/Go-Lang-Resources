package main

import (
	"fmt"
)

func Greetings(name string) string {
	return "Hello " + name
}

func Sum(a, b int, c float64) int {
	return a + b + int(c)
}

func main() {
	fmt.Println(Greetings("Yash"))
	fmt.Println(Sum(2, 3, 4.5))
}
