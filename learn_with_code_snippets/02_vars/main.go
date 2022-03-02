package main

import "fmt"

var i string = "oyutside the main function"

// j := "outside the main function"  not allowed
const j string = "J outside the main function"

func main() {

	fmt.Println("Variables")

	var a int = 5
	var b string = "hi"
	var c bool = true
	var d float32 = 1.5

	//short hand
	e := 5
	f := "Hello"
	g := true
	h := 1.5

	// name := "Yash"
	// email := "iyashjayesh@gmail.com"

	name, email := "Yash", "iyashjayesh@gmail.com"

	fmt.Println(name, email)

	fmt.Println(a, b, c, d, e, f, g, h)

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", e)
	fmt.Printf("%T\n", f)
	fmt.Printf("%T\n", g)
	fmt.Printf("%T\n", h)
	fmt.Printf("%T\n", i)
	fmt.Printf("%T\n", j)

}
