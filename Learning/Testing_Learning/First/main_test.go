package main

import (
    "testing"
)

func TestAdd(t *testing.T) {
    result := Add(2, 3)
    expected := 5
    if result != expected {
        t.Errorf("Add(2, 3) = %d; want %d", result, expected)
    }
}

func TestSubtract(t *testing.T) {
    result := Subtract(5, 3)
    expected := 2
    if result != expected {
        t.Errorf("Subtract(5, 3) = %d; want %d", result, expected)
    }
}

func TestMultiply(t *testing.T) {
    result := Multiply(2, 3)
    expected := 6
    if result != expected {
        t.Errorf("Multiply(2, 3) = %d; want %d", result, expected)
    }
}

func TestDivide(t *testing.T) {
    result, err := Divide(6, 3)
    if err != nil {
        t.Errorf("Divide(6, 3) returned error: %v", err)
    }
    expected := 2
    if result != expected {
        t.Errorf("Divide(6, 3) = %d; want %d", result, expected)
    }

    // Test division by zero
    _, err = Divide(6, 0)
    if err == nil {
        t.Error("Expected error for division by zero")
    }
}
