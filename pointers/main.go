package main

import "fmt"

func main(){
	x := 5
	y := &x

	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	// Use * to read value from address

	fmt.Println(*y)

	// Change value

	// Can change x, expect to see same change in *y
	x = 10
	fmt.Println(*y)

	// Can change *y, expect to see same change in x
	*y = 15
	fmt.Println(x)

	


}