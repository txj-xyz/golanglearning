package main

import "fmt"

func main() {
	i := 1       // inferred `i int = 1`
	for i <= 3 { // 3 iterations
		fmt.Println(i)
		i = i + 1 // increment `int i` by `+1`every loop iteration.
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j) // loops 4 times including 0 iteration
	}

	for {
		fmt.Println("loop") // iterates one time and exits
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 { // no idea what the percent is doing, perhaps `modulo` remainders
			continue
		}
		fmt.Println(n)
	}
}
