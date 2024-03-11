package stringutil

import "strings"

// Reverse reverses a given string.
func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

// CountWords counts the number of words in a given string.
func CountWords(s string) int {
    words := strings.Fields(s)
    return len(words)
}
