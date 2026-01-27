package main

import "fmt"

func main() {
	var intNum int16 = 32767
	// this will cause issues
	intNum = intNum + 1 
	fmt.Println(intNum)
}