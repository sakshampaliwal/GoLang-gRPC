package main

import (
    "reflect"
    "testing"
)

func TestSortInts(t *testing.T) {
    // Define test cases using a slice of structs
    testCases := []struct {
        input    []int // input
        expected []int // expected output
    }{
        {[]int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}, []int{1, 1, 2, 3, 3, 4, 5, 5, 5, 6, 9}}, // Test case 1
        {[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},         // Test case 2
        {[]int{}, []int{}},                                                               // Test case 3: empty slice
        {[]int{1}, []int{1}},                                                             // Test case 4: single element
    }

    // Iterate over test cases
    for _, tc := range testCases {
        // Create a copy of input slice
        inputCopy := make([]int, len(tc.input))
        copy(inputCopy, tc.input)

        // Run test for each test case
        SortInts(tc.input)
        if !reflect.DeepEqual(tc.input, tc.expected) {
            t.Errorf("SortInts(%v) = %v, expected %v", inputCopy, tc.input, tc.expected)
        }
    }
}
