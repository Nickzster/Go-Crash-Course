package main

import (
	"fmt"
	"time"
)

func routine(msg string) {
	fmt.Println("Creating new routine:", msg)
}

func main() {
	fmt.Println("Hello, go routines!")

	go routine("Routine 1!")
	go routine("Routine 2!")
	go routine("Routine 3!")

	time.Sleep(1 * time.Second)
}