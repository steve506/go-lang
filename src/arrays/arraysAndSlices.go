package main

import (
	"fmt"
)

// here, an array is called. If the [] do not have a specific value
// it is considered a Slice
func printer(w [4]string){
	for _, word := range w {
		fmt.Printf("%s", word)
	}
	fmt.Printf("\n")
	// an array when passed to a function, is passed as a copy of the entire array,
	//not the original array allocated in Memory, so any changes would be applied
	//to the copy only
	w[2]= "Blue"

}
// For slices, the size is not specified. They are faster than arrays
// given that smaller pieces of data are passed to the functions.
func printerSlice(w []string){
	for _, word := range w {
		fmt.Printf("%s", word)
	}
	fmt.Printf("\n")
	/*This time the changes are applied to the Slice itself */
	w[2]= "Blue"

}

func printerSliceOfSlice(w []string){
	for _, word := range w {
		fmt.Printf("%s", word)
	}
	fmt.Printf("\n")
}

func main(){
	// this is how to declare and assign values to an array
	// the size of the array is hardcoded within the []
	words := [4]string{"the", "quick", "brown", "fox"}
	fmt.Printf("%s\n", words[2])

	//slice
	wordsSlice := []string{"the", "quick", "brown", "fox", "jumps", "over", "the",
	"lazy", "dog"}

	printer(words)
	printer(words)
	printerSlice(wordsSlice)
	printerSlice(wordsSlice)
	//you can specify to pass a slice of a slice, not necessarily the whole thing.
	// This way you work all the time with the same piece of memory instead of copies.

	// NOTE: the ending reference of the range of the slice is not included
	// i.e. [0:2] passes positions 0 and 1 but not 2.
	printerSliceOfSlice(wordsSlice[0:2])
	// print from [5] to the end of the slice
	printerSliceOfSlice(wordsSlice[5:len(wordsSlice)])
	//this syntax does the same
	printerSliceOfSlice(wordsSlice[5:])	


	//another way to declare a slice with the make keyword
	wordsDeclare := make([]string, 4)
	wordsDeclare[0] = "The"
	wordsDeclare[1] = "Quick"
	wordsDeclare[2] = "Brown"
	wordsDeclare[3] = "Fox"

	printerSlice(wordsDeclare)

	//another way to declare a slice. This time specifying that it
	//starts with zero items and can hold up to 4.
	wordsAppend := make([]string, 0, 4)
	
	//Find out what's the current size of a slice and its capacity
	fmt.Printf("%s\n", "Length and capacity:")
	fmt.Printf("%d %d\n", len(wordsAppend), cap(wordsAppend))

	wordsAppend = append(wordsAppend, "THE")
	wordsAppend = append(wordsAppend, "QUICK")
	wordsAppend = append(wordsAppend, "BROWN")
	wordsAppend = append(wordsAppend, "FOX")

	printerSlice(wordsAppend)

	fmt.Printf("%s\n", "Length and capacity:")
	fmt.Printf("%d %d\n", len(wordsAppend), cap(wordsAppend))

	// appending allows to exceed the original capacity of a slice
	// by doubling the initial allocated space

	//This would be the fifth element for a slice of 4
	wordsAppend = append(wordsAppend, "JUMPS")
	
	// now the slice of 4 has 5 elements and a new capacity of 8
	fmt.Printf("%s\n", "Length and capacity:")
	fmt.Printf("%d %d\n", len(wordsAppend), cap(wordsAppend))

	printerSlice(wordsAppend)

	// slices can be copied also
	newWords := make([]string, 4)
	copy(newWords, wordsSlice)
	// the copy is an entirely new slice which can be modified without impacting
	// the original one
	newWords[2] = "Green"
	printerSlice(newWords)
	printerSlice(wordsSlice)

	//On the other hand, if a set of an array is re-assigned to a different variable,
	//both elements would be pointing to the same memory space
	partialSlice := wordsSlice[:4]
	partialSlice[2] = "Black"
	printerSlice(partialSlice)
	printerSlice(wordsSlice)
}
