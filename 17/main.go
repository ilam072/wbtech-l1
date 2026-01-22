package main

import "fmt"

func main() {
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5}, 3))  // 2
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5}, -2)) // -1
	fmt.Println(binarySearch([]int{1, 2, 3, 4, 5}, 6))  // -1
}

func binarySearch(nums []int, item int) int {
	low, high := 0, len(nums)-1

	for low <= high {
		mid := (low + high) / 2
		guess := nums[mid]

		if guess == item {
			return mid
		}
		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}
