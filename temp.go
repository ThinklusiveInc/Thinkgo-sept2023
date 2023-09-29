package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter the temperature in Celsius:")
	fmt.Scan(&n)
	f := (float32(n) * 1.8) + 32
	fmt.Println("Temperature in Fahrenheit is:", f)
}
