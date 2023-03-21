package main

import (
	"fmt"
)

func worker(input <-chan int, output chan<- int) {
	for data := range input {
		// Process the data
		result := data * 2

		// Send the result back to the main goroutine
		output <- result
	}
}

func main() {
	// Create the channels
	input := make(chan int)
	output := make(chan int)

	// Start the worker goroutine
	go worker(input, output)

	// Send some data to the worker
	input <- 10

	// Wait for the worker to send back the result
	result := <-output
	fmt.Println(result) // Output: 20
}
