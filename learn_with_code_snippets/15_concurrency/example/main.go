package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

// var wg sync.WaitGroup // A sysnchronization mechanism is a way to ensure that a set of goroutines
// var mut sync.Mutex    // A mutext is a mutual exclusion lock.

func sendRequest(s string) {
	// fmt.Println("Sending request to:", s)
	// defer wg.Done()
	res, err := http.Get(s)
	if err != nil {
		panic(err)
	}

	// mut.Lock()
	// defer mut.Unlock()

	fmt.Printf("Response status: [%d] %s\n", res.StatusCode, s)
}

func main() {
	fmt.Println("Go Concurrency Example")

	if len(os.Args) < 2 {
		log.Fatalln("Usage: go run main.go <url1> <url2> ... <urlN>")
	}

	for _, url := range os.Args[1:] {
		 sendRequest("https://" + url)
		// wg.Add(1)
	}

	// wg.Wait()
}

// Running in git lab
// time go run main.go google.com youtube.com github.com facebook.com microsoft.com apple.com amazon.com twitter.com

// Time Benchmarking
// 1st run: without go routines
// real    0m11.627s
// user    0m0.000s
// sys     0m0.046s

// 2nd run: with go routines
// real    0m4.149s
// user    0m0.015s
// sys     0m0.061s
