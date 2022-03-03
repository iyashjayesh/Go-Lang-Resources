package main

import "fmt"

//Go Channels
//Channels are the pipes that connect concurrent goroutines. You can send values into channels from one goroutine and receive those values into another goroutine.
//Go channels are typed by the values they convey.
// Channels is used to share data between gorountines.
// <-

func foo(c chan int, someValue int) {
	c <- someValue * 5
}

func main() {
	fmt.Println("Go Channels")

	//making a channel
	fooVal := make(chan int)

	go foo(fooVal, 5)
	go foo(fooVal, 10)

	// v1 := <- fooVal
	// v2 := <- fooVal

	v1, v2 := <-fooVal, <-fooVal
	fmt.Println(v1, v2)

}
