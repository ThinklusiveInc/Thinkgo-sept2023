package main

import "fmt"

func sum(a, b int) int {
	return a + b

}

func main() {
	var a = 5.5
	const b int = 6
	var c = "hi"
	d := 8
	e := "hello"
	var f int
	f = b + d
	fmt.Println("Results", a, b, c, d, e)
	fmt.Println("Sum of ", b, "and", d, "is", f)
	if f == b && b == f {
		fmt.Println("Correct")
	} else {
		fmt.Println("False")
	}
}

// package main

// // fmt package provides the function to print anything
// import "fmt"

// // function to add the two integer numbers
// func addNumber(number1, number2 int) int {
// 	number1 = 18
// 	number2 = 9

// 	return number1 + number2
// }
// func main() {

// 	// define the integer variables we want to add
// 	var number1, number2, number3 int
// 	var number4, number5, number6 int
// 	// initializing the variables
// 	number1 = 18
// 	number2 = 9

// 	number4 = 5
// 	number5 = 6

// 	// calling the function and storing the result
// 	number3 = addNumber(number1, number2)
// 	number6 = addNumber(number4, number5)

// 	// printing the results
// 	fmt.Println("The addition of ", number1, " and ", number2, " is \n", number3, "\n(adding two integers outside the function)")
// 	fmt.Println("The addition of ", number4, " and ", number5, " is \n", number6, "\n(adding two integers outside the function)")
// }
