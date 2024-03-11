package mathutil

import "testing"

func TestFactorial(t *testing.T) {
    testCases := []struct {
        n        int
        expected int
    }{
        {0, 1},   // 0! = 1
        {1, 1},   // 1! = 1
        {2, 2},   // 2! = 2 * 1 = 2
        {3, 6},   // 3! = 3 * 2 * 1 = 6
        {4, 24},  // 4! = 4 * 3 * 2 * 1 = 24
        {5, 120}, // 5! = 5 * 4 * 3 * 2 * 1 = 120
    }

    for _, tc := range testCases {
        result := Factorial(tc.n)
        if result != tc.expected {
            t.Errorf("Factorial(%d) = %d; want %d", tc.n, result, tc.expected)
        }
    }
}
