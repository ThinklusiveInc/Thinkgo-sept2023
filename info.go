package main

import (
	"fmt"
)

func main() {
	var name, email string
	var age int

	fmt.Print("Enter your name: ")
	fmt.Scan(&name)

	fmt.Print("Enter your age: ")
	fmt.Scan(&age)

	fmt.Print("Enter your email: ")
	fmt.Scan(&email)

	fmt.Printf("User Information:\nName: %s\nAge: %d\nEmail: %s\n", name, age, email)
}
