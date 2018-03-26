package main

import (
	"fmt"
	"time"
)

// Closing channels is a good way to synchronize events
// This function will print a message as soon as goCh
// receives something. This is using the close function
// to start something
func printer(msg string, goCh chan bool) {
	<-goCh
	fmt.Printf("%s\n", msg)
}

// the close function can be used to stop something
// as well. This function will print non-stop
// until the close function is sent through the channel.
// The exact opposite of the previous func.
func printerStop(msg string, stopCh chan bool) {
	// this is an infinite loop that will
	// only stop when the channel receives a value
	for {
		select {
		case <- stopCh:
			return
		default:
			fmt.Printf("%s\n", msg)
		}
	}
	
}
func main() {
	fmt.Printf("\nSTART\n")
	// this is the receiving size of the channel
	goCh := make(chan bool)
	// Here 10 goroutines are created, each of them
	// is calling the printer function.
	// The result will be the program staying on hold
	// until goCh receives something
	for i := 0; i < 10; i++ {
		go printer(fmt.Sprintf("printer: %d", i), goCh)
	}
	// This is a hold timer of 5 seconds
	// after which something will be sent to the channel
	time.Sleep(5 * time.Second)
	// although goCh is a bool channel
	// all channels can receive the close instruction
	close(goCh)
	// After the close instruction is sent,
	// That will trigger the function to proceed
	// with the 10 printer instances that were triggered above.
	// So the close sent through the channel is a starter
	// the timer below will simply wait 5 seconds
	// before moving on with the program execution
	time.Sleep(5 * time.Second)

	fmt.Printf("\nCLOSE CHANNEL AS STOPPER\n")

	// this is the receiving channel for the stop
	// example
	stopCh := make(chan bool)
	// Here 10 goroutines are created and they
	// will keep on printing forever due to the
	// infinite loop within the function
	for i := 0; i < 10; i++ {
		go printerStop(fmt.Sprintf("printer: %d", i), stopCh)
	}
	// This is a hold timer of 5 seconds
	// after which something will be sent
	// to the stop channel
	time.Sleep(5 * time.Second)
	// although goCh is a bool channel
	// all channels can receive the close instruction
	close(stopCh)
	fmt.Printf("STOP!\n")
	// This time the close function did not start
	// the print but stopped it.
	// The timer below waits 5 seconds before the program terminates
	time.Sleep(5 * time.Second)


}