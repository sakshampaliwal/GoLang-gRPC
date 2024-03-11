package main

import "fmt"

// Define a media interface
type Media interface {
    Play()
    Stop()
}

// Define a concrete type - Audio
type Audio struct {
    Title string
    Duration int
}

// Implement the Media interface for Audio
func (a Audio) Play() {
    fmt.Printf("Playing audio: %s\n", a.Title)
}

func (a Audio) Stop() {
    fmt.Printf("Stopping audio: %s\n", a.Title)
}

// Define another concrete type - Video
type Video struct {
    Title string
    Duration int
}

// Implement the Media interface for Video
func (v Video) Play() {
    fmt.Printf("Playing video: %s\n", v.Title)
}

func (v Video) Stop() {
    fmt.Printf("Stopping video: %s\n", v.Title)
}

// Function to play media
func PlayMedia(media Media) {
    media.Play()
}

func main() {
    // Create instances of Audio and Video
    audio := Audio{Title: "Song Title", Duration: 240}
    video := Video{Title: "Movie Title", Duration: 7200}

    // Play media
    PlayMedia(audio)
    PlayMedia(video)
}
