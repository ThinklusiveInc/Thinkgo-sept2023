//Write a Go program that creates a goroutine to print numbers from 1 to 5 concurrently.
//Make sure the main goroutine also prints "Main Goroutine" after the goroutine finishes.

package main

import (
	"fmt"
	"sync"
)

func printNumbers() {
	for i := 0; i <= 5; i++ {
		fmt.Println(i)

	}

}
func main() {

	go printNumbers()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		printNumbers()

	}()

	wg.Wait()

	fmt.Println("main goroutine")
}
