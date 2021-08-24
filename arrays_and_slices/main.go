package main

import "fmt"

func main(){
	// Arrays
	var fruitArr [2] string 

	// Assign Values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"

	fmt.Println(fruitArr)

	vegitableArr  := [2] string { "Carrot", "Brussel-Sprouts"}

	fmt.Println(vegitableArr)

	// Slices
	hobbySlices := []string{"Biking", "Coding", "Legos"}
	hobbySlices = append(hobbySlices, "Foobar");

	fmt.Println(hobbySlices)

	fmt.Println("Number of hobbies: ",len(hobbySlices))
	fmt.Println("Middle hobbies:", hobbySlices[1:3])

	
}

