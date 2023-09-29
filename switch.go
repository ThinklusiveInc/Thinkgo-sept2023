package main

import (
	"fmt"
	"time"
)

func main() {
	i := 2
	fmt.Println("write", i, "as")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() {
	case time.Friday, time.Monday:
		fmt.Println("It is weekend")
	default:
		fmt.Println("It is weekday")
	}

	t := time.Now()
	switch {
	case t.Hour() < 11:
		fmt.Println("it is before noon")
	default:
		fmt.Println("it is after noon")
	}
	whatamI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("it is bool")
		case int:
			fmt.Println("it is integer")
		default:
			fmt.Printf("i don't know type %T \n", t)
		}
	}
	whatamI(3)
	whatamI("hey")
	whatamI(true)

}
