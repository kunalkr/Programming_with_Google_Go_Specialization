package main

import (
	"fmt"
)

func main() {
	var input float64

	fmt.Printf("Enter float value: ")
	fmt.Scan(&input)
	fmt.Printf("Truncated value: %v\n", int64(input))
}
