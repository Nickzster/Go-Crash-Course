package main

import "fmt"

func main(){
	x := 15
	y := 10

	// Common Convention
	if x <= y {
		fmt.Printf("%d is less than or equal to %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}
	// Also Acceptable
	if (x < y) {
		fmt.Printf("%d is less than %d\n", x, y)
	}

	color := "red"

	switch color {
	case "red": 
		fmt.Println("Red!")
	case "blue":
		fmt.Println("Blue!")
	default:
		fmt.Println("No Color!")
	}


}