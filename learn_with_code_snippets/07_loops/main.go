package main

import "fmt"

func main(){
	fmt.Println("Loops")

	// for loop -  long method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		// i++
		i = i+1
	}

	// for loop - short method
	for j := 0; j < 10; j++ {
		fmt.Printf("Numbers :%d\n", j)
	}

	// for loop - infinite loop
	// for {
	// 	fmt.Println("Infinite loop")	
	// }

	// for loop - break statement
	for k := 0; k < 10; k++ {
		if k == 5 {
			break
		}
		fmt.Printf("Numbers :%d\n", k)
	}

	// for loop - continue statement
	for l := 0; l < 10; l++ {
		if l == 5 {
			continue
		}
		fmt.Printf("Numbers :%d\n", l)
	}

	// for loop - range loop
	numbers := []int{1,2,3,4,5,6,7,8,9,10}
	for i, v := range numbers {
		fmt.Printf("Numbers :%d and index :%d\n", v, i)
	}

	// for loop - range loop with strings
	names := []string{"Yash", "Jayesh", "Kumar", "Singh"}
	for i, v := range names {
		fmt.Printf("Names :%s and index :%d\n", v, i)
	}

	// for loop - range loop with maps
	m := map[string]string{"a": "apple", "b": "ball", "c": "cat"}
	for k, v := range m {
		fmt.Printf("Key :%s and Value :%s\n", k, v)
	}

	// for loop - range loop with arrays
	arr := [5]int{1,2,3,4,5}
	for i, v := range arr {
		fmt.Printf("Numbers :%d and index :%d\n", v, i)
	}

	// for loop - range loop with slices
	s := []int{1,2,3,4,5}
	for i, v := range s {
		fmt.Printf("Numbers :%d and index :%d\n", v, i)
	}

	
}