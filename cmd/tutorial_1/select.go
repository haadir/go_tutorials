package main

import (
	"fmt"
	"time"
)

func main () {
	// independent event sources
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func (){
		time.Sleep(2 * time.Second)
		ch1 <- "Message from channel 1"

		time.Sleep(2 * time.Second) 
		ch2 <- "Message from channel 2"
	}()
		
	// event listener
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <- ch1:
			fmt.Println(msg1)
		case msg2 := <- ch2:
			fmt.Println(msg2)
		}
	}
}