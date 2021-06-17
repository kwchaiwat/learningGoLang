package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 3, 4, 5, 5, 7, 7}

	for i := 0; i < len(nums); i++ {
		index := removeDuplicates(nums)
		if index > len(nums)-1 {
			break
		}
		nums = removeIndex(nums, index)
	}
	fmt.Println(nums)
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var i int = 0
	for j := 1; j < len(nums); j++ {
		if nums[j] != nums[i] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}
