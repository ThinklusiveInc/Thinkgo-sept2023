package main

import "fmt"

var (
	b, c int = 8, 10
)

func main() {

	var a string = "Madhu"
	fmt.Println(a)

	result := NewNum(b, c)
	fmt.Println(result)
}

func NewNum(x, y int) int {
	return x + y
}
