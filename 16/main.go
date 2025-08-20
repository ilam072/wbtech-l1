package main

import (
	"fmt"
)

func main() {
	fmt.Println(quickSort([]int{19, 15, 23, 67, 1, 4, 3, 2, 33, 33, 19}))

}

func quickSort(nums []int) []int {
	var less, greater []int
	if len(nums) < 2 {
		return nums
	} else {
		pivot := nums[0]
		for i := 0; i < len(nums[1:]); i++ {
			if nums[1:][i] > pivot {
				greater = append(greater, nums[1:][i])
			} else {
				less = append(less, nums[1:][i])
			}
		}

		less = append(quickSort(less), pivot)
		greater = quickSort(greater)
		return append(less, greater...)
	}
}
