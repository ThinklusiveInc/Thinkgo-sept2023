package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}

	fmt.Println(sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index", i)
		}
	}
	// iterating over data structures like arrays and slices
	number := []int{45, 67, 89, 45, 66}
	for index, value := range number {
		fmt.Println("index: , value:\n", index, value)
	}
	// Iterating over maps
	gradesstu := map[string]int{"madhu": 25, "kai": 25, "avi": 26}
	for index, value := range gradesstu {
		fmt.Println("map:", index, value)
	}
	// iterating over string
	communi := "I am jennifer wife of madhu"
	for index, runevalue := range communi {
		fmt.Println("string:", index, runevalue)

	}
	// iterating over channels
	datachannel := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			datachannel <- i
		}
		close(datachannel)
	}()
	for value := range datachannel {
		fmt.Println("received :", value)
	}
}
