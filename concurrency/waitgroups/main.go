package main

import (
	"fmt"
	"sync"
)

func myFunc(waitgroup *sync.WaitGroup) {
	fmt.Println("Inside goroutine")
	waitgroup.Done()
}

func main(){

	fmt.Println("Beginning of main")
	var waitgroup sync.WaitGroup
	waitgroup.Add(1)
	go myFunc(&waitgroup)
	waitgroup.Wait()
	fmt.Println("End of main")

}