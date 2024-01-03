package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a) // "initial"

	var b, c int = 1, 2 // type inferrence multiple vars "c int = 1"
	fmt.Println(b, c)   // "1 -> 2" ??? (b being explicitly set as an int, c being inferred here)

	var d = true
	fmt.Println(d) // true

	var e int      // constructing an `int` with nothing inside default `int = 0`
	fmt.Println(e) // 0

	f := "apple"   // inferred string type from `:=`
	fmt.Println(f) // "apple"
}
