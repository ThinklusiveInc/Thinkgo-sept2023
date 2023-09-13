package main

import (
	"fmt"
)

func main() {
	// Declare and initialize variables
	var integerVar int = 10
	var floatVar float64 = 3.14
	var boolVar bool = true
	var stringVar string = "Hello, Go!"

	// Declare constants
	const pi float64 = 3.14159265359
	const gravity float64 = 9.81

	// Perform operations
	sum := integerVar + int(floatVar)
	difference := integerVar - int(floatVar)
	product := integerVar * int(floatVar)
	quotient := float64(integerVar) / floatVar
	modulus := integerVar % int(floatVar)

	// Display results
	fmt.Println("Integer Variable:", integerVar)
	fmt.Println("Float Variable:", floatVar)
	fmt.Println("Boolean Variable:", boolVar)
	fmt.Println("String Variable:", stringVar)

	fmt.Println("Constant PI:", pi)
	fmt.Println("Constant Gravity:", gravity)

	fmt.Println("Sum:", sum)
	fmt.Println("Difference:", difference)
	fmt.Println("Product:", product)
	fmt.Println("Quotient:", quotient)
	fmt.Println("Modulus:", modulus)

}
