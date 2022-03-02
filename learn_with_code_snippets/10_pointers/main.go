package pointers

import "fmt"

func main(){
	fmt.Println("Pointers")

	// Declaring a variable
	var x int = 10
	fmt.Println(x)

	// Declaring a pointer
	var p *int = &x
	fmt.Println(p)
	fmt.Println(*p)
	
	// Assigning a value to a pointer
	*p = 20
	fmt.Println(x)

	// Declaring a pointer to a pointer
	var pp **int = &p
	fmt.Println(pp)
	fmt.Println(*pp)
	fmt.Println(**pp)

	// Declaring a pointer to a struct
	type person struct {
		name string
		age int
	}

	var p1 *person = &person{"Yash", 20}
	fmt.Println(p1)
	fmt.Println(p1.name)
	fmt.Println(p1.age)

	// Declaring a pointer to a struct
	var p2 **person = &p1
	fmt.Println(p2)
	fmt.Println(*p2)
	fmt.Println((*p2).name)
	fmt.Println((*p2).age)
	
	// Declaring a pointer to a struct
	var p3 ***person = &p2
	fmt.Println(p3)
	fmt.Println(*p3)
	fmt.Println(**p3)
	fmt.Println((**p3).name)
	fmt.Println((**p3).age)

	// Declaring a pointer to a struct
	
}