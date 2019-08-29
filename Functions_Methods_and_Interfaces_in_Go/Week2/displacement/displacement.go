package main

import (
	"fmt"
	"math"
)

// GenDisplaceFn : Function to return a displacement function
func GenDisplaceFn(a, u, s0 float64) func(t float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + u*t + s0
	}
}

func main() {
	a := 0.0
	u := 0.0
	s0 := 0.0
	t := 0.0

	fmt.Print("Acceleration: ")
	fmt.Scan(&a)

	fmt.Print("Initial Velocity: ")
	fmt.Scan(&u)

	fmt.Print("Initial Displacement: ")
	fmt.Scan(&s0)

	fmt.Print("Time: ")
	fmt.Scan(&t)

	displacementFunc := GenDisplaceFn(a, u, s0)

	fmt.Printf("\nDisplacement after %v seconds: %v\n", t, displacementFunc(t))
}
