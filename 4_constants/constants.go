package main

import (
	"fmt"
	"math"
)

const s string = "constant" // locked string type for the variable

func main() {
	fmt.Println(s)      // "constant" locked to a `string` type
	const n = 500000000 // untyped `int` idk what this does.

	const d = 3e20 / n // 3 + 20zeros divided by "n" variable
	fmt.Println(d)     // show the result of the division product

	fmt.Println(int64(d))    // show the number now as an `int64()`
	fmt.Println(math.Sin(n)) // Sine of `int n`

}
