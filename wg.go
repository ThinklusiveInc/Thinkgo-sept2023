package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

var mutex sync.Mutex

func printNumber(i int) {
	defer wg.Done()
	mutex.Lock()
	fmt.Println(i)
	mutex.Unlock()
}

func main() {

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go printNumber(i)
	}
	wg.Wait()
}
