package main

import (
	"fmt"
	"time"
)
// function that will create a channel of channels.
// This is useful when it is required that the function
// itself generates the data for multiple receiving channels
func emit(chanChannel chan chan string, done chan bool){
	defer close(chanChannel)
	// This is the timer
	t := time.NewTimer(3 * time.Second)
	// This is the channel that will be passed
	// through the outer channel
	wordChannel := make(chan string)
	chanChannel<-wordChannel
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
		case <-t.C:
			fmt.Printf("STOP!\n")
			return
		}
	}
}

func main() {
	fmt.Printf("Execution starts\n\n")
	// Create a channel of channels
	channelCh := make(chan chan  string)
	doneCh := make(chan bool)

	go emit(channelCh, doneCh)
	// wordCh is receiving the strings coming from the channel of channels
	wordCh := <- channelCh
	// here word is reading each element provided from wordCh, which
	// is receiving from the channels of channels created by emit
	for word := range wordCh {
		fmt.Printf("%s ", word)
	}

}