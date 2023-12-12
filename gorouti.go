// how to implement a concurrency in  go
// we can obtain concurrency in go by using goroutine and channels.
//goroutine is a lightweight threads can be created by using "GO" keyword. It allows concurrent execution of functions.
// channels is a communication and synchronization bwt goroutines.this can be created by using make and can be used to send and receive values.
// to start a go routine , we use "go" keyword before the function call this will create a new goroutine that executes the function concurrently.
// channels can be used to send or receive data bwt goroutines.

package main

import (
	"fmt"
	"time" // we import time for time related functionality
)

func f(fro int) { // function f takes the int parameter "fro"
	for i := 0; i < 3; i++ { // here it contains a for loop that prints thw value of "fro" and "i" three times
		fmt.Println(fro, ":", i) // print fro and i
	}
}

func main() { //main function

	f(19) // we call function "f" with argument "19"

	go f(45) // we started new goroutine using "go" keyword

	go func(msg string) { // created goroutine for anonymous func
		fmt.Println(msg)

	}("going")

	time.Sleep(time.Second) // sleep function is used to pause the execution of the program for one second
	fmt.Println("done")
}
