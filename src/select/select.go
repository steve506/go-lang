// The select function permits to have multiple channels
// at the same time and selecting from them within a function as desired
package main

import (
	"fmt"
	"time"
)

// transmits strings indefinitely with an option
// to close
func emit(wordChannel chan string, done chan bool){
	words := []string{"The", "quick", "brown", "fox"}
	i := 0
	for{
		select {
		case wordChannel <- words[i]:
			i++
			if i== len(words) {
				i=0
			}
		case <- done:
			fmt.Printf("%s", "Got done!\n")
			close(done)
			return
		}
	}
}

// channels are bidirectional
// emit2 is similar to emit, with the difference that
// it returns a true value back after the function has received the true
// value that closes it
func emit2(wordChannel chan string, done chan bool){
	words := []string{"The", "quick", "brown", "fox"}
	i := 0
	for{
		select {
		case wordChannel <- words[i]:
			i++
			if i== len(words) {
				i=0
			}
		case <- done:
			done <- true
			return
		}
	}
}

// timers can be used also. This function has its own timer
// which provides a third branch to select.
// This time the program will close after the timeout gets to 0
func emitTimer(wordChannel chan string, done chan bool){
	words := []string{"The", "quick", "brown", "fox"}
	i := 0
	// This is the timer
	t := time.NewTimer(3 * time.Second)
	for{
		select {
		case wordChannel <- words[i]:
			i++
			if i== len(words) {
				i=0
			}
		case <- done:
			fmt.Printf("%s", "Got done!\n")
			close(done)
			return
		// here the timer is considered as a selection,
		// so the function is assigning a value to one of the select
		// branches on its own
		case <-t.C:
			fmt.Printf("STOP!\n")
			return

		}
	}
}

func main(){

	//receive the channel that transmits infinite strings
	wordCh := make(chan string)
	doneCh := make(chan bool)

	go emit(wordCh, doneCh)

	// receives strings from the channel 100 times
	for i := 0; i < 100; i++ {
		fmt.Printf("%s ", <-wordCh)
	} 
	// after the 100th string, sends a true statement to the
	// bool channel to close it
	doneCh <- true

	// here emit2 is called, a true statement is sent back to close the channel
	// and the program keeps waiting for the channel to send the other true
	// value back
	wordCh2 := make(chan string)
	doneCh2 := make(chan bool)

	go emit2(wordCh2, doneCh2)

	for i := 0; i < 100; i++ {
		fmt.Printf("%s ", <-wordCh2)
	} 
	doneCh2 <- true
	<-doneCh2

	// emit with a timer
	fmt.Printf("\n\nRun with timer\n\n")

	wordCh3 := make(chan string)
	doneCh3 := make(chan bool)

	go emitTimer(wordCh3, doneCh3)
	// This time there is an infinite for loop
	// which would never close (because the channel generates infinite strings)
	// and there is no control as with the 100 strings of the prior examples. However,
	// The timer has the program terminate as soon as it expires.
	for word := range wordCh3{
		fmt.Printf("%s ", word)
	} 
	doneCh3 <- true
	<-doneCh3
}