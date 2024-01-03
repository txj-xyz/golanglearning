package main

import "fmt"

func main() {
	if 7%2 == 0 { // 7 goes into 2 -> 3 times oddly so there is a `remainder`
		fmt.Println("7 is even")
	} else { // this statement will always fire because 7 into 2 has a remainder and is not `0`
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 { // 8 goes into 4 evenly 2 times so there is `0` remainder
		fmt.Println("8 is divisible by 4")
	}

	if 8%2 == 0 || 7%2 == 0 { // 7 or 8 could be even if its got no remainders
		fmt.Println("either 8 or 7 are even")
	}

	if num := 9; num < 0 { // if lower than 0 then show its negative
		fmt.Println(num, "is negative")
	} else if num < 10 { // lower than 10 (making it 1 digit numbers)
		fmt.Println(num, "has 1 digit")
	} else { // this means that all conditions are not true
		fmt.Println(num, "has multiple digits")
	}
}
