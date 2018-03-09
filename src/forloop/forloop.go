package main

import(
	"fmt"
)

func main(){
	var counter int
	counter = 0

	for counter < 10{
		fmt.Printf("Hello, World!\n")
		counter++
	}
// variable declaration; condition; action for every iteration
	for counter2 :=0; counter2 <10; counter2++{
		fmt.Printf("Hello, World!\n")
	}

	for i :=0; i <10; i++{
		fmt.Printf("Hello, Everybody!\n")
	}
// using multiple variables
	for i, j :=0, 1; i <10; i, j = i+1, j*2{
		fmt.Printf("%d Hello, Everybody!\n", j)
	}
	
// using booleans
	var stop bool
	i := 0

	for !stop{
		fmt.Printf("Hello, Loop!\n")
		i++
		if i == 10{
			stop = true
		}
	}

}

