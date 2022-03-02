package main

import "fmt"

func main() {
	fmt.Println("Maps")

	// map declaration
	var map_name map[string]string

	// map initialization
	map_name = make(map[string]string)

	// map assignment
	map_name["1"] = "test1"
	map_name["2"] = "test2"

	// map access
	fmt.Println(map_name["1"])

	// map iteration
	for key, value := range map_name {
		fmt.Println(key, value)
	}

	// map length
	fmt.Println(len(map_name))

	//test
	emails := make(map[string]string)

	emails["Bob"] = "bob@gmail.com"
	emails["Sharon"] = "sharon@gmail.com"

	fmt.Println(emails)

	fmt.Println(emails["Bob"])

	delete(emails, "Bob")

	fmt.Println(emails)

	// test
	emailss := map[string]string{"yash":"yash@gmail.com", "meet":"meet@gmail.com"}

	fmt.Println(emailss)
}
