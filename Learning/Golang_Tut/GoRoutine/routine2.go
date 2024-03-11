package main

import (
    "fmt"
)

func sender(c chan string) {
    for i := 0; i < 5; i++ {
        // Send a message into the channel
        c <- fmt.Sprintf("Message %d", i)
    }
    // Close the channel after sending all messages
    close(c)
}

func receiver(c chan string) {
    for msg := range c {
        // Receive a message from the channel
        fmt.Println("Received:", msg)
    }
}

func main() {
    // Create a channel of type string
    c := make(chan string)

    // Start sender and receiver goroutines
    go sender(c)
    go receiver(c)

    // Keep the main function running until both sender and receiver goroutines complete
    var input string
    fmt.Println("Press Enter to exit")
    fmt.Scanln(&input)
    fmt.Println("Exiting...")
}


// Output:
// Press Enter to exit
// Received: Message 0
// Received: Message 1
// Received: Message 2
// Received: Message 3
// Received: Message 4