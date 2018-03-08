package main

import (
	"fmt"
)

func main() {
	isTrue := !true
	var mybool bool
	
	mybool = false

	fmt.Printf("Value: %t\n", isTrue)
	fmt.Printf("Value: %t\n", mybool)
}