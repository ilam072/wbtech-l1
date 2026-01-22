package main

import "fmt"

func main() {
	fmt.Println(quickSort([]int{19, 15, 23, 67, 1, 4, 3, 2, 33, 33, 19}))
}

func quickSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	var less, greater []int
	pivot := nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] > pivot {
			greater = append(greater, nums[i])
		} else {
			less = append(less, nums[i])
		}
	}

	less = append(quickSort(less), pivot)
	greater = quickSort(greater)

	return append(less, greater...)
}
