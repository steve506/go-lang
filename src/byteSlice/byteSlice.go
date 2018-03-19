package main

import(
	"fmt"
	"os"
)

func main(){
	// OpenFile requires the filename, the flag e.g. os.O_RDONLY, os.O_RDWR and a perm FileMode which represents the mod permissions over the file
	// OpenFile allows to read and edit files. For reading or creating files, methods such as os.Open or os.Create are used instead.
	f, err := os.OpenFile("text.txt", os.O_RDWR, 0660)
	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
	defer f.Close()

	b := make([]byte, 100)

	n, err := f.Read(b)

	// print the slice that read 100 bytes i.e. 100 characters from the file
	// byte slice is printed in hex (% x)
	fmt.Printf("%d: % x\n", n, b)

	// convert the byte slice to string
	stringVersion := string(b)
	fmt.Printf("%d %s\n", n, stringVersion)

	someString := "hi!\n"

	// Read and Write require byte slices to work
	_, error := f.Write([]byte(someString))
	if error != nil{
		fmt.Printf("%s", error)
	}


}