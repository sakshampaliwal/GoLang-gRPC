package main

import (
    "fmt"
)

// Function to add two numbers
func Add(a, b int) int {
    return a + b
}

// Function to multiply two numbers
func Multiply(a, b int) int {
    return a * b
}

// Function to calculate a complex result using Add and Multiply
func Calculate(x, y int) int {
    // Some complex calculation
    sum := Add(x, y)
    product := Multiply(x, y)
    result := sum + product
    return result
}
