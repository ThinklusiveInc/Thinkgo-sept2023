package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64
	fmt.Print("Enter the radius of the circle: ")
	fmt.Scan(&radius)

	area := math.Pi * radius * radius
	fmt.Printf("The area of the circle is %.2f\n", area)
}
