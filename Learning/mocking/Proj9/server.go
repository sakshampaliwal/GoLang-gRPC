package main

// import "fmt"

// Add adds two numbers.
func Add(a, b int) int {
	return a + b
}

// Subtract subtracts the second number from the first.
func Subtract(a, b int) int {
	return a - b
}

// CalculateResult performs addition and subtraction and returns the result.
func CalculateResult(a, b, c int) int {
	sum := Add(a, b)
	difference := Subtract(sum, c)
	return difference
}
