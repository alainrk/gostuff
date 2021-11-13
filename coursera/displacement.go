package main

import (
	"fmt"
	"math"
)

func promptForFloat(prompt string) float64 {
	var input float64
	fmt.Print(prompt)
	fmt.Scanf("%f", &input)
	return input
}

func GenDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v0*t + s0
	}
}

func dispacement() {
	a := promptForFloat("Enter the value of a: ")
	v0 := promptForFloat("Enter the value of v0: ")
	s0 := promptForFloat("Enter the value of s0: ")

	displaceFn := GenDisplaceFn(a, v0, s0)

	t := promptForFloat("Enter the value of t: ")
	fmt.Printf("Dispacement at t = %.2f is %.2f\n", t, displaceFn(t))
}
