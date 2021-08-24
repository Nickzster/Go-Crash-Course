package main

import "fmt"

func main(){
	// Define map
	emails := make(map[string]string)
	// Assign key/values
	emails["Bob"] = "bob@example.com"
	emails["Alice"] = "alice@example.com"

	fmt.Println(emails)
	KEY := "Bob"
	fmt.Printf("Name: %s, Email: %s\n", KEY, emails[KEY])
	fmt.Printf("Number of emails: %d\n", len(emails))

	// Delete from map
	delete(emails, KEY)
	fmt.Println("Number of emails:",len(emails))

	// Alternate way
	ages := map[string]int{"Bob": 42, "Alice": 33, "Eve": 29}
	fmt.Println(ages)

}