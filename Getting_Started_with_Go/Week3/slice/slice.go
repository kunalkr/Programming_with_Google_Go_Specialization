package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var x string
	slice := make([]int, 0, 3)

	for {
		fmt.Printf("Enter integer: ")
		fmt.Scan(&x)

		if num, err := strconv.Atoi(x); err == nil {
			slice = append(slice, num)
			sort.Ints(slice)
			fmt.Printf("%v\n", slice)
		} else {
			break
		}
	}

}
