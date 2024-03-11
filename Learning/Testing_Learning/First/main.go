package main

import "errors"

// Add takes two integers and returns their sum.
func Add(a, b int) int {
    return a + b
}

// Subtract takes two integers and returns the result of subtracting the second from the first.
func Subtract(a, b int) int {
    return a - b
}

// Multiply takes two integers and returns their product.
func Multiply(a, b int) int {
    return a * b
}

// Divide takes two integers and returns the result of dividing the first by the second.
// It returns an error if the second number is zero.
func Divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
