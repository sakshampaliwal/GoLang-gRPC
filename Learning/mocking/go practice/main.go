package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
	fmt.Println(sayHello())
}

func sayHello() string {
	return "Hello guys "
}