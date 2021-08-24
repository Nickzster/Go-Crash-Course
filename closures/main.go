package main

import "fmt"

func adder() func(int)int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main(){
	// Closures ==> Functions are first class citizens ==> Functional Programming Paradigms apply!
	sum := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(sum(i))
	}
}