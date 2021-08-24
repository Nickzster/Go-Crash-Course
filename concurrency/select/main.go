package main

import (
	"fmt"
	"time"
)

func main(){
	c1 := make(chan string)
	c2 := make(chan string)
	
	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func(){
		for {
			c2 <- "Every two seconds"
			time.Sleep(time.Second * 2)
		}
	}()

	// blocking code:
	// for {
	// 	fmt.Println(<-c1)
	// 	fmt.Println(<-c2)
	// }

	// non-blocking:
	for {
		select {
		case msg1 := <- c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}
}