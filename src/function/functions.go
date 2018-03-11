package main

import (
    "fmt"
    "os"
)


func printer(msg, msg2 string, repeat int) {
    for repeat > 0{
        fmt.Printf("%s", msg)
        fmt.Printf("%s\n", msg2)
        repeat --
    }
}

// func name(parameters) return type {}
func printer2(msg string) error {
    _, err := fmt.Printf("%s\n", msg)
    return err
}

// Go can return multiple values simultaneously
func printer3(msg string) (string, error) {
    msg += "\n"
    _, err := fmt.Printf("%s\n", msg)
    return msg, err
}

func printer4(msg string) error {
    /*defer keyword queues an action to be
    executed as soon as the program finishes.
    Defers are executed from bottom to top
    */
    defer fmt.Printf("\n\n\n")
    defer fmt.Printf("More\n")

    _, err := fmt.Printf("%s", msg)

    return err
}

// if the output type is called by name, it doesn't
// need to be specified as return 
func printer5(msg string) (e error) {
    f, e := os.Create("helloworld.txt")
    if e != nil{
        return e
    }
    defer f.Close()
    
    f.WriteString(msg)
    return
}

// functions in Go can leverage an undefined amount of parameters,
// similar to *args in Python
func printer6(msgs ...string) {
    for _, message := range msgs {
        fmt.Printf("%s\n", message)
    }
}

func main() {

    printer("Hello,", "world!", 5)
    printer2("Hello, world!2")
    appendedMessage, myError := printer3("hello, world!3\n")

    if myError != nil{
        fmt.Printf("% x\n", appendedMessage)
    }

    printer4("Hello, world!4")
    printer5("Hello 5!")
    printer6("Hello", "Hola", "How are you?", "Sabor")

}
