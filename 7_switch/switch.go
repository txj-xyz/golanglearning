package main

import (
	"fmt"
	"time"
)

func main() {

	// Basic switch usage
	// checking each iteration of the loop
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	}

	// Catch if the time of week is either `saturday` or `sunday`
	// or just a generic `weekday` that !== `saturday` | `sunday`
	// ----
	// This demonstrates using commands as an "or" in this case is "true or true" from the weekday
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	// Check if its before or after noon from an inferred type from `Time.time`
	// ----
	// This demonstrates a basic `if{}else{}` using a switch instead from the `t` variable
	t := time.Now() // Time.time
	switch {
	case t.Hour() < 12: // before noon
		fmt.Println("It's before noon")
	default: // after noon
		fmt.Println("It's after noon")
	}

	// Decalre a function that compares the type given in the `i` param type given
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I am a bool")
		case int:
			fmt.Println("I am an int")
		default:
			fmt.Printf("Don't know the type %T\n", t)
		}
	}
	whatAmI("hello") // Don't know the type "string"
	whatAmI(true)    // I am a bool
	whatAmI(1)       // I am an int
}
