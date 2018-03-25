package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)
//Function to return the size of the body of
//a URL once downloaded
func getPage(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	return len(body), nil

}

// channels allow processes to run in parallel
func getter(url string, size chan int) {
	length, err := getPage(url)
	if err == nil {
		size <- length
	}
}
// An improved version of the same function.
// This one uses Sprintf which returns a string instead of printing it
// This way the channel even processes the output in parallel.
func getterImproved(url string, size chan string) {
	length, err := getPage(url)
	if err == nil {
		size <- fmt.Sprintf("Improved: The size of %s is %d bytes", url, length)
	}
}

// function that receives the URLs on one channel
// and return the response on another
// it also reveals the ID of the worker that processed the operation
func worker(urlCh chan string, sizeCh chan string, id int ) {
	for {
		url := <-urlCh
		length, err := getPage(url)
		if err == nil {
			sizeCh <- fmt.Sprintf("Dual Channel: The size of %s is %d bytes. ID: %d", url, length, id)
		} else {
			sizeCh <- fmt.Sprintf("Error getting %s: %s", url, err)
		}
	}
}

// Not only transmission can be executed concurrently. reception can run
// in parallel as well. This method will put each url from the array 
// into the channel for processing via a goroutine instead of the sequential form
func generator(url string, urlCh chan string){
	urlCh <- url
}

func main(){
	fmt.Printf("Start\n\n")

	// Test the url function
	url := "http://www.google.com/"
	// put the results of getPage into variables
	pageLength, err := getPage(url)
	if err != nil {
		os.Exit(1)
	}
	fmt.Printf("%s is length %d bytes\n", url, pageLength)

	// try getPage over a list of URLs
	urls := []string{"http://www.google.com", "http://www.yahoo.com",
		"http://www.bing.com", "http://bbc.co.uk"}

	// this way of processing the array is sequential, which
	// can be slow since every iteration is processed after the other
	for _, url := range urls {
		pageLength, err := getPage(url)
		if err != nil {
			os.Exit(1)
		}
		fmt.Printf("%s is length %d bytes\n", url, pageLength)
	}

	// channel that will receive the sizes of getter
	size := make(chan int)
	// create four separate goroutines connected to the same
	// channel called "size" and transmitting/receiving at the same time
	for _, url := range urls {
		go getter(url, size)
	}
	// this time the processing is faster because the goroutines run in parallel
	// so there is no need to 
	for i := 0; i < len(urls); i++ {
		fmt.Printf("Parallel: %s length is %d\n", urls[i], <-size)
	}
	// channel to receive the channel from getterImproved
	sizeImproved := make(chan string)
	// create a separate goroutine for each url in the array
	for _, url := range urls {
		go getterImproved(url, sizeImproved)
	}

	// consume the entire array urls and the transmission of the channel
	// which is a string with the results. All running concurrently
	for i := 0; i < len(urls); i++ {
		fmt.Printf("%s\n", <-sizeImproved)
	}
	
	// Channels that transmit to and receive from the worker
	// urlCh is provided with the urls
	//sizeCh provides the results from the function
	urlCh := make(chan string)
	sizeCh := make(chan string)

	// create the channels for worker. Ten channels are created
	// so i would represent the worker ID
	for i :=0; i < 10; i++ {
		go worker(urlCh, sizeCh, i)
	}

	// Transmit the URLs to urlCh
	for _, url := range urls {
		urlCh <-url		
	}

	// consume the results from sizeCh.
	// from the 10 channels created above,
	// any of them can be assigned the processing of each iteration
	for i := 0; i < len(urls); i++ {
		fmt.Printf("%s\n", <-sizeCh)
	}

	fmt.Printf("\nWorker and generator\n")

	// channels for the scenario where the urls are transmitted via a goroutine
	urlCh2 := make(chan string)
	sizeCh2 := make(chan string)

	// this time only two workers will be used, just to test load balancing
	// whenever there is more work than goroutines, the goroutines wait until they
	// are finished sending their work i.e. availability
	for i :=0; i < 2; i++ {
		go worker(urlCh2, sizeCh2, i)
	}
	// here, instead of "manually" putting each URL into the channel for processing
	// the generator function does it via a Channel i.e. concurrently
	// so go generator populates the url channel and go worker receives the channel
	// for processing
	for _, url := range urls {
		go generator(url, urlCh2)		
	}

	// consume the results from sizeCh.
	// from the 10 channels created above,
	// any of them can be assigned the processing of each iteration
	for i := 0; i < len(urls); i++ {
		fmt.Printf("%s\n", <-sizeCh2)
	}

}