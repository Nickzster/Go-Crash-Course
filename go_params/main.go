package main

import "fmt"

func myFancyFunc(message string) (modifier string, newFancyMessage string) {
	myNewFancyMessage := message + " (Super fancy!)"
	return "Message:", myNewFancyMessage
}


func main(){
	fmt.Println(myFancyFunc("A message!"))
}