package main

import (
	"fmt"
)

// Swap : Function to swap numbers at location 'index' and 'index + 1'
func Swap(nums []int, index int) {
	temp := nums[index]
	nums[index] = nums[index+1]
	nums[index+1] = temp
}

// BubbleSort : Function to sort the input slice of numbers using BubbleSort algorithm
func BubbleSort(nums []int) {
	for i := len(nums); i > 0; i-- {
		for j := 1; j < i; j++ {
			if nums[j-1] > nums[j] {
				Swap(nums, j-1)
			}
		}
	}
}

func main() {
	numbers := make([]int, 0, 10)
	var num int

	for idx := 0; idx < cap(numbers); idx++ {
		fmt.Scan(&num)
		numbers = append(numbers, num)
	}

	BubbleSort(numbers)

	for _, num := range numbers {
		fmt.Print(num, " ")
	}
	fmt.Println()
}
