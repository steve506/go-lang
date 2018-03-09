package main

import(
	"fmt"
)

func main() {
	atoz  := `the quick brown fox jumps over the lazy dog\n`

	// atoz[0:9] prints the string from the start to the ninth character
	fmt.Printf("%s\n", atoz[0:9])

	fmt.Printf("%s\n", atoz[15:19])
// take the string from char 15 to the end
	fmt.Printf("%s\n", atoz[15:])

	for i, r := range atoz {
		fmt.Printf("%d %c\n", i, r)
	}

	for _, r := range atoz {
		fmt.Printf("%c\n", r)
	}

	fmt.Printf("%s\n", atoz)
}