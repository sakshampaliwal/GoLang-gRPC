// func TestIsEven(t *testing.T) {
// 	// Test case: 4 is an even number
// 	if !IsEven(4) {
// 		t.Errorf("IsEven(4) = false, expected true")
// 	}

// 	// Test case: 7 is an odd number
// 	if IsEven(7) {
// 		t.Errorf("IsEven(7) = true, expected false")
// 	}

// 	// Test case: -2 is an even number
// 	if !IsEven(-2) {
// 		t.Errorf("IsEven(-2) = false, expected true")
// 	}

// 	// Test case: 0 is an even number
// 	if !IsEven(0) {
// 		t.Errorf("IsEven(0) = false, expected true")
// 	}

// 	// Test case: 1000000 is an even number
// 	if !IsEven(1000000) {
// 		t.Errorf("IsEven(1000000) = false, expected true")
// 	}
// }

// or

package main

import "testing"

func TestIsEven(t *testing.T) {
    // Define test cases using a slice of structs
    testCases := []struct {
        number   int  // input
        expected bool // expected output
    }{
        {4, true},  // Test case 1: 4 is an even number
        {7, false}, // Test case 2: 7 is an odd number
        {-2, true}, // Test case 3: -2 is an even number
        {0, true},  // Test case 4: 0 is an even number
        {1000000, true}, // Test case 5: 1000000 is an even number
    }

    // Iterate over test cases
    for _, tc := range testCases {
        // Run test for each test case
        if result := IsEven(tc.number); result != tc.expected {
            t.Errorf("IsEven(%d) = %t, expected %t", tc.number, result, tc.expected)
        }
    }
}
