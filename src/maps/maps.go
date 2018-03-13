package main

import (
	"fmt"
)

func main() {
	// A map is an array in which the index can be arbitrary
	// it works like a dictionary
	dayMonths := make(map[string]int)
	dayMonths["Jan"] = 31
	dayMonths["Feb"] = 28
	dayMonths["Mar"] = 31
	dayMonths["Apr"] = 30
	dayMonths["May"] = 31
	dayMonths["Jun"] = 30
	dayMonths["Jul"] = 31
	dayMonths["Aug"] = 31
	dayMonths["Sep"] = 30
	dayMonths["Oct"] = 31
	dayMonths["Nov"] = 30
	dayMonths["Dec"] = 31

	fmt.Printf("Days in february: %d\n", dayMonths["Feb"])
	// everytime a non-existent index is called, a value of 0 is returned
	fmt.Printf("Days in february: %d\n", dayMonths["January"])

	/* Validation syntax.
	declare two variables, the 1st one will take the value of the assigned map index,
	while the second one (ok) will be true only if the index exists.
	This example makes ok go false  */
	days, ok := dayMonths["January"]
	if !ok {
		fmt.Printf("Can't get days for January\n")
	} else {
		fmt.Printf("%d\n", days)
	}

	/*
	this example keeps ok as true
	*/
	days, ok = dayMonths["Jan"]
	if !ok {
		fmt.Printf("Can't get days for January\n")
	} else {
		fmt.Printf("%d\n", days)
	}

	// a for loop can handle dual variables for the indexes and the values
	for myMonth, myDays := range dayMonths {
		fmt.Printf("%s has %d days\n", myMonth, myDays)
	}

	// You can just ignore the indexes if you don't need them, such as the name of the month in this example
	has31 := 0

	for _, myDays2 := range dayMonths {
		if myDays2 == 31{
			has31 ++
		}
	}
	fmt.Printf("%d months have 31 days\n", has31)

	// Elements from the map can be deleted
	delete(dayMonths, "Feb")

	for mon := range dayMonths {
		fmt.Printf("%s\n", mon)
	}

	// Another way to declare a map, very similar to a dictionary
	dayMonths2 := map[string]int{
		"Jan": 31,
		"Feb": 28,
		"Mar": 31,
		"Apr": 30,
		"May": 31,
		"Jun": 30,
		"Jul": 31,
		"Aug": 31,
		"Sep": 30,
		"Oct": 31,
		"Nov": 30,
		"Dec": 31,
	}

	for month2, days2 := range dayMonths2{
		fmt.Printf("Second map: %s has %d days\n", month2, days2)
	}

}