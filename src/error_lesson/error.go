package main

import (
	// There is a dedicated package for error handling
	"errors"
	"fmt"
	"os"
)

var (
	errorEmptyString = errors.New("Unwilling to print an empty string. Errors package")
)

// Go relies on error handling rather than exceptions
func printer(msg string) error {
	// Errors can be custom
	if msg == "" {
		// The Errorf function returns a given string in an error format
		return fmt.Errorf("Unwilling to print an empty string")

	}
	_, err := fmt.Printf("%s\n", msg)
	return err
}

func printer2(msg string) error {
	if msg == "" {
		// Use the error created using the errors package
		return errorEmptyString

	}
	_, err := fmt.Printf("%s\n", msg)
	return err
}
// In case there is a scenario where something goes horribly wrong,
// panic can be used as a Exception to stop execution of the program
func printerPanic(msg string) error {
	if msg == "" {
		panic(errorEmptyString)

	}
	_, err := fmt.Printf("%s\n", msg)
	return err
}


func main(){
	if err := printer("Hello, World!"); err != nil {
		os.Exit(1)
	}

	//trigger the custom error
	if err := printer(""); err != nil {
		fmt.Printf("Printer failed: %s\n", err)
	}

	//trigger the error using the errors package
	if err := printer2(""); err != nil {
		fmt.Printf("Printer failed: %s\n", err)
	}

	//error variables can be used for comparison and validation
	if err := printer2(""); err != nil {
		fmt.Printf("You tried to print an empty string!\n")
	} else {
		fmt.Printf("Printer failed: %s\n", err)
	}

	//When the program panics, all execution stops and a stack trace is generated
	//Panic should be used ONLY when it is strictly necessary
	if err := printerPanic(""); err != nil {
		fmt.Printf("You tried to print an empty string! Panic!\n")
	} else {
		fmt.Printf("Printer failed: %s\n", err)
	}

}