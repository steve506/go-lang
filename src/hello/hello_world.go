package main

import (
	"fmt"
	"os"
)

const(
	message1 = "The meaning of life is %d\n"
	answer = 42

	messagenumbers = "%d %d\n"
	answer1 = iota*2
	answer2

)

func main() {
	var message string
	message = "Hello World!\n"
	
	message2 := "Hello everybody!\n"

	fmt.Printf(message)
	fmt.Printf(message2)
	fmt.Printf(message1, answer)
	fmt.Printf(messagenumbers, answer1, answer2)

	if numberBytes, theError := fmt.Printf("Hello, World!\n"); theError != nil {
		os.Exit(1)
	} else {
	fmt.Printf("Printed %d bytes\n", numberBytes)
	}
}