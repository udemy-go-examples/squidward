/*
Hands-on exercise #73 - wrapper
*/
package main

import (
	"fmt"
	"time"
)

// Wrapper function for adding timing information
func TimedFunction(fn func()) {
	start := time.Now()
	fn()
	elapsed := time.Since(start)
	fmt.Println("Elapsed time:", elapsed)
}

// Function to be wrapped
func MyFunction() {
	time.Sleep(7 * time.Second) // Simulate some work
	fmt.Println("MyFunction completed")
}
func main() {
	// Call the wrapped function
	//TimedFunction(MyFunction)

	// wrap functions
	wrappedFunc := func() {
		TimedFunction(MyFunction)
	}
	Logger(wrappedFunc)

	// call function with each item in slice
	data := []int{1, 2, 3, 4, 5}
	callbackFunc := func(num int) {
		fmt.Println("Processing number:", num)
	}
	ProcessData(data, callbackFunc)
}

func Logger(f func()) {
	fmt.Println("Starting execution...")
	f()
	fmt.Println("Execution completed.")
}

func ProcessData(data []int, callback func(int)) {
	for _, item := range data {
		callback(item)
	}
}
