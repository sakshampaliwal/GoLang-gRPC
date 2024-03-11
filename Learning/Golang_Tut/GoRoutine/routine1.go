package main

import (
	"fmt"
	"time"
)

func main() {
	// Start a new Go routine
	go printNumbers()
	
	// Main function continues execution
	for i := 0; i < 5; i++ {
		fmt.Println("Main function")
		time.Sleep(time.Second)
	}
}

func printNumbers() {
	// Function running concurrently
	for i := 0; i < 5; i++ {
		fmt.Println("Go routine")
		time.Sleep(time.Second)
	}
}


// Output:
// Main function
// Go routine
// Main function
// Go routine
// Main function
// Go routine
// Main function
// Go routine
// Main function
// Go routine
// Go routine
