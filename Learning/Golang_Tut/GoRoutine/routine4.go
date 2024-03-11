package main

import (
	"fmt"
	"time"
)

func reteight(result chan int) {
	time.Sleep(8 * time.Second)
	result <- 8
}
func retfive(result chan int) {
	time.Sleep(5 * time.Second)
	result <- 5
}

func rettwo(result chan int) {
	time.Sleep(2 * time.Second)
	result <- 2
}

func main() {
	result := make(chan int)

	go reteight(result)
	go retfive(result)
	go rettwo(result)

	sum1 := <-result
	sum2 := <-result
	sum3 := <-result

	fmt.Println("sum1: ",sum1)
	fmt.Println("sum2: ",sum2)
	fmt.Println("sum3: ",sum3)
	
	totalSum := sum1 + sum2

	fmt.Println("Total Sum:", totalSum)
}
