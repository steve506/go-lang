package main

import (
	"fmt"
)

func main() {
	var pi float64
	pi = 3.14
	fmt.Printf("Value: %.2f\n", pi)

	pi2 := float64(3.14)
	pi = 3.14
	fmt.Printf("Value: %f\n", pi2)

	nine := uint8(9)
	fmt.Printf("Value: %d\n", nine)

}