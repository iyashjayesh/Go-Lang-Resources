package main

import "fmt"

func main() {
	fmt.Println("Arrays and Slices")

	// Arrays
	var a [5]int //It shoul be always fixed size in array declaration
	a[0] = 1
	a[1] = 2
	a[2] = 3
	a[3] = 4
	a[4] = 5

	fmt.Println("Array a:", a)

	// Decalre and Assign at same time
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array b:", b)

	//Slices - Dynamic size
	c := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice c:", c)

	fmt.Println("Length of c:", len(c))
	fmt.Println("Capacity of c:", cap(c))
	fmt.Println("Value of c[2]:", c[2])
	fmt.Println("Value of c[2:4]:", c[2:4])
}
