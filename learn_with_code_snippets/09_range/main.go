package main

import "fmt"

func main(){
	fmt.Println("Range")
	// range
	// range is used to iterate over collections such as arrays and slices.
	// range returns two values, one is the index and the other is the value of the element.
	// range returns the index and value of the element in the array.
	// range returns the index and value of the element in the slice.
	// range returns the index and value of the element in the map.
	// range returns the index and value of the element in the string.
	// range returns the index and value of the element in the channel.
	// range returns the index and value of the element in the function.
	// range returns the index and value of the element in the pointer.
	// range returns the index and value of the element in the struct.
	// range returns the index and value of the element in the interface.
	// range returns the index and value of the element in the map.
	// range returns the index and value of the element in the array.
	// range returns the index and value of the element in the slice.

	// range over arrays
	var arr [5]int
	for i, v := range arr {
		fmt.Printf("Index :%d and Value :%d\n", i, v)
	}
	// while in case of not using index we can keep it as _ 
	// for _, v := range arr {
	// 	fmt.Printf("Value :%d\n", v)
	// }

	// range over slices
	var sli []int
	sli = append(sli, 1,2,3,4,5,6,7,8,9,10)
	for i, v := range sli {
		fmt.Printf("Index :%d and Value :%d\n", i, v)
	}

	// range over maps
	m := map[string]string{"a": "apple", "b": "ball", "c": "cat"}
	for k, v := range m {
		fmt.Printf("Key :%s and Value :%s\n", k, v)
	}

	// range over strings
	str := "Yash"
	for i, v := range str {
		fmt.Printf("Index :%d and Value :%c\n", i, v)
	}

	// range over functions

	// range over pointers

	// range over structs

	// range over interfaces

	// range over channels


}