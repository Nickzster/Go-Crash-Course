package main

import (
	"fmt"
	"strconv"
)

// Goroutines are concurrently running jobs.
// Can be "producers" or "consumers" of some kind.
func age(msg chan string) {
	for {
		fmt.Println("In age!")
		msg <- strconv.Itoa(25)
	}

}

func lastName(msg chan string) {
	for {
		fmt.Println("In lastname!")
		msg <- "Zimmermann "
	}

}

func firstName(msg chan string) {
	for {
		fmt.Println("In firstname!")
		msg <- "Nick "
	}

}

func factoryFunc(msg chan string, fn chan string, ln chan string, age chan string) {
	for {
		fullName := ""
		fullName += <-fn
		fullName += <-ln
		fullName += <-age
		msg <- fullName
	}
}



func main() {

	firstNameChannel := make(chan string)
	lastNameChannel := make(chan string)
	ageChannel := make(chan string)
	factory := make(chan string)

	defer close(firstNameChannel) //defer will defer this call until the END of execution of main.
	defer close(lastNameChannel) //defer will defer this call until the END of execution of main.
	defer close(ageChannel) //defer will defer this call until the END of execution of main.
	defer close(factory)

	go firstName(firstNameChannel)
	go lastName(lastNameChannel)
	go age(ageChannel)
	go factoryFunc(factory, firstNameChannel, lastNameChannel, ageChannel)

	fmt.Println(<-factory)
	fmt.Println("This will execute last since channels are blocking!")
	
}