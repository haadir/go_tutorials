package main

import (
	"fmt"
	"time"
)

func printMessage(msg string) {
	fmt.Println(msg)
}

func main() {
	go printMessage("Hello world!")
	time.Sleep(time.Second)
}