package main

import "fmt"

func sum(result *int, x int) {
	*result += x
}

func mul(result *int, x int) {
	*result *= x
}

// RACE CONDITION
// It happens when 2 or more processes executes concurrently, accessing
// to shared resources, this can cause inconsistencies in the result,
// that cannot be deterministic anymore.

// This happens in this example, where the result is not deterministic,
// being the sum and multiplication not commutative on each other, and
// being the order of the operations not deterministic itself.

func racecond() {
	res := 1
	for i := 1; i < 10; i++ {
		go mul(&res, i)
		go sum(&res, i)
	}
	fmt.Println("result: ", res)
}
