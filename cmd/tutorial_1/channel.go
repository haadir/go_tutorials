package main

import "fmt"

func main() {
	// the := syntax means declare AND assign, = is just for assigning, can't use to create variable
	messages := make(chan string)

	// go in front means run concurrently, the () means call this function immediately
	go func() {
		messages <- "1"
		messages <- "2"
		messages <- "3"
		close(messages) 
	}()

	for msg := range messages {
		fmt.Println(msg)
	}
}