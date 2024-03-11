package stringutil

import (
	"testing"
	
)

func TestReverse(t *testing.T) {
    input := "hello"
    result := Reverse(input)
    expected := "olleh"
    if result != expected {
        t.Errorf("Reverse(%q) = %q; want %q", input, result, expected)
    }
}

func TestCountWords(t *testing.T) {
    input := "hello world"
    result := CountWords(input)
    expected := 2
    if result != expected {
        t.Errorf("CountWords(%q) = %d; want %d", input, result, expected)
    }
}

func TestCountWords_EmptyString(t *testing.T) {
    input := ""
    result := CountWords(input)
    expected := 0
    if result != expected {
        t.Errorf("CountWords(%q) = %d; want %d", input, result, expected)
    }
}

func TestCountWords_MultipleSpaces(t *testing.T) {
    input := "hello   world"
    result := CountWords(input)
    expected := 2
    if result != expected {
        t.Errorf("CountWords(%q) = %d; want %d", input, result, expected)
    }
}
