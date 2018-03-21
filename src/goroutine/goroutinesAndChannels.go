package main

import(
	"fmt"
	"math/rand"
)

// Create a function that is a channel of strings to be passed
// this is the transmitting side of the goroutine
func emit(c chan string){
	// this is an array of strings
	words := []string{"The","quick","brown","fox"}
	// Go through the array and pass it via the channel
	for _, word := range words{
		c <- word
	} 
	// Close the channel as soon as it passed everything
	close(c)
}
func emitNotClose(c chan string){
	// this is an array of strings
	words := []string{"The","quick","brown","fox"}
	// Go through the array and pass it via the channel
	for _, word := range words{
		c <- word
	} 
}
// function to generate random numbers infinitely
func emitRandomIntegers(c chan int){
	// an infinite loop
	for{
		// a random number between 1 and 1000
		c <- rand.Intn(1000)
	}
}
// Here's a nice way of generating IDs infinitely
// which can be controlled from the receiving side of the channel
func emitID(idChan chan int){
	// declare a variable
	var id int
	// variable id will start with 0
	id = 0
	for{
		idChan <- id
		id++
	}
}
 


func main(){
	// This is the receiving part of the channel
	wordChannel := make(chan string)
	// Call the channel function. This is how GO runs the channel concurrently to
	// the rest of the program and populates wordChannel
	//This is the receiving part of the goroutine
	go emit(wordChannel)

	// Print the contents of the array already populated by the channel
	for word := range wordChannel{
		fmt.Printf("%s ", word)
	}

	fmt.Printf("\n")

	// It is not mandatory to use a range to receive from the channel
	// You can explicitly  receive to a variable
	// Here's another channel. You can open as many instances of the same
	// channel as desired
	wordChannel2 := make(chan string)
	go emit(wordChannel2)
	// Here the first item of wordChannel is received
	word2 := <- wordChannel2
	fmt.Printf("%s\n", word2)
	// Here the remaining items of wordChannel are sequentially received
	word2 = <- wordChannel2
	fmt.Printf("%s\n", word2)
	word2 = <- wordChannel2
	fmt.Printf("%s\n", word2)
	word2 = <- wordChannel2
	fmt.Printf("%s\n", word2)
	// 4 items have been received so far, so the channel is closed now and the call
	// below will receive nothing
	word2 = <- wordChannel2
	fmt.Printf("%s\n", word2)
	// the ok value can be used to validate whether there is still something to
	// i.e. the channel is still open. The code below will print the ok value as false
	word2, ok := <- wordChannel2
	fmt.Printf("%s %t\n", word2, ok)

	//multiple instances of exactly the same channel can be executed and 
	// they will run in parallel and will be interleaved as a single bigger channel
/*-	wordChannel3 := make(chan string)
	go emitNotClose(wordChannel3)
	go emitNotClose(wordChannel3)
	go emitNotClose(wordChannel3)
	for wordIter := range wordChannel3{
		_, ok := <- wordChannel3
		if ok == true{
			fmt.Printf("%s %t\n", wordIter, ok)
		}
	}
	*/
	fmt.Printf("\n")

	// This is the receiving part for the function that generates random numbers
	randoms := make(chan int)
	go emitRandomIntegers(randoms)
	// read from the channel
	// Given that the transmitting function of the channel has an infinite for
	// This execution will never end
	for number := range randoms{
		fmt.Printf("%d \n", number)
		break
	}



	// receiving part of the ID generator
	idChan := make(chan int)
	go emitID(idChan)
	// consume an ID from the channel
	fmt.Printf("%d\n", <-idChan)
	// Consume more IDs sequentially generated
	fmt.Printf("%d\n", <-idChan)
	fmt.Printf("%d\n", <-idChan)
	fmt.Printf("%d\n", <-idChan)
	fmt.Printf("%d\n", <-idChan)
	fmt.Printf("%d\n", <-idChan)

}