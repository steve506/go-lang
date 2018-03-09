package main

import(
	"fmt"
	"os"
)

func main() {

	n, err := fmt.Printf("Hello, World!")
// line added to cheat and force the script to match
// the second case and demonstrate the fallthrough keyword
//	n = 0

// the switch statement is a form of "if" logic
	switch {
	case err != nil:
		os.Exit(1)
	case n == 0:
		fmt.Printf("No bytes output")
// the fallthrough keyword allows to continue going through the next case
//		fallthrough
	case n != 13:
		fmt.Printf("wrong number of bytes: %d", n)
	default:
		fmt.Printf("OK!")
	}

	fmt.Printf("\n")

	atoz := "The quick brown fox jumps over the lazy dog"
	vowels := 0
	consonants := 0
	zeds := 0

	for _, r := range atoz {
		switch r {
		case 'a', 'e', 'i', 'o', 'u':
			vowels ++
		case 'z':
			zeds ++
			fallthrough
		default:
			consonants ++
		}


	}

	fmt.Printf("Vowels: %d; Consonants: %d (Zeds: %d)\n", vowels, consonants, zeds)




}
