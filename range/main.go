package main

import "fmt"

func main(){
	ids := []int{42, 56, 72, 89}

	// Loop thru ID's
	// for in / for of in javascript
	for i, id := range ids {
		fmt.Printf("ID %d: %d\n", i, id )
	}

	for _, id := range ids {
		fmt.Printf("ID %d\n", id )
	}

	// Sum ID's

	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum:", sum)

	// Map Example

	ages := map[string]int{"Bob": 42, "Alice": 33, "Eve": 29}

	for k,v := range ages {
		fmt.Printf("%s is %d years old\n", k, v)
	}

	for k := range ages {
		fmt.Println("Name: " + k)
	}

	
}