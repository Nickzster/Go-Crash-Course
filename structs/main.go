package main

import (
	"fmt"
	"strconv"
)

// Define person struct

type Person struct {
	firstName string
	lastName string
	city string
	gender string
	age int
}

// Greeting method (value receiver)

func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// Has Birthday method (pointer receiver)

func (p *Person) hasBirthday() {
	p.age++
}


// getMarried (pointer receiver)

func (p *Person) getsMarried(newLastName string){
	if p.gender == "f" || p.gender == "Female" {
		p.lastName = newLastName
	}
}



func main(){
	//Init person using struct
	person1 := Person {firstName: "Nick", lastName: "Zimmermann", city: "Pleasanton", gender: "Male", age: 25}
	person2 := Person {"Erin", "O'Bra", "Fairfield", "Female", 22}

	fmt.Println(person1)
	fmt.Println(person2)


	person1.age++ //This is NOT immutable!
	fmt.Println(person1)

	// Method
	person1.hasBirthday() //This is also NOT immutable!
	fmt.Println(person1.greet())

	person2.getsMarried(person1.lastName)
	fmt.Println(person2.greet())


	
}