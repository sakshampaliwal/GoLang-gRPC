package main

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMainFunctionOutput(t *testing.T) {
	// Redirect stdout to capture output
	// old := captureStdout()
	// defer restoreStdout(old)

	// Call main function
	data := sayHello()

	// Get captured output
	// captured := restoreStdout(old)

	// Assert the output
	// assert.Contains(t, captured, "Hello world!", "Expected output contains 'Hello world!'")
	assert.Contains(t, captured, "Hello guys", "Expected output contains 'Hello guys'")
}

// // Function to capture stdout
// func captureStdout() *bytes.Buffer {
// 	var buf bytes.Buffer
// 	stdout = &buf
// 	return &buf
// }

// // Function to restore stdout
// func restoreStdout(old *bytes.Buffer) string {
// 	ret := old.String()
// 	stdout = old
// 	return ret
// }

// var stdout *bytes.Buffer // Global variable to store original stdout
