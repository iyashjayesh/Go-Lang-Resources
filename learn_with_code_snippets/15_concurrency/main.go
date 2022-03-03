package main

import (
	"fmt"
	"time"
)

func main() {
	go count("sheep")
	count("fish")
}

//Concurrency is not Parallelism
//Concurrency is the ability to execute multiple tasks at the same time. It is the ability to execute multiple tasks concurrently.
//The ability to execute multiple tasks concurrently means that the tasks can run concurrently and multiple tasks can be running at the same time.

//Go routines are lightweight because they are lightweight.
//Maaged by the Go runtime.
//Flexible Stack Size 2KB

//Very Important for Go Routines
// Do not communicate by sharing memory, share memory by communicating.

//Paralleslism is not concurrency

//Threads
//Go is not a threading language.
//Go is a concurrency-enabled language.
// Threads are managed by OS.
// Fixed Stack Size By default 1MB

func count(thing string) {
	for i := 0; i < 10; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 100)
	}
}
