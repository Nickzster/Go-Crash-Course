package main

import "fmt"

func main(){
	// Long Method
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// Short Method
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// Fizzbuzz
	for i := 1; i <= 100; i++ {
		if i % 15 == 0 {
			fmt.Printf("FizzBuzz %d\n", i)
			continue
		}
		if i % 3 == 0 {
			fmt.Printf("Fizz %d\n", i)
			continue
		}
		if i % 5 == 0 {
			fmt.Printf("Buzz %d\n", i)
			continue
		}
	}
}