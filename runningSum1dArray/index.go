package main

import "fmt"

func main() {

	// 	Input: nums = [1,2,3,4]
	// Output: [1,3,6,10]
	// Explanation: Running sum is obtained as follows: [1, 1+2, 1+2+3, 1+2+3+4].

	nums := []int{1, 1, 1, 1, 1}
	fmt.Println(runningSum(nums))
}

func runningSum(nums []int) []int {
	result := nums
	for i := 1; i < len(nums); i++ {
		result[i] = result[i-1] + nums[i]
	}
	return result
}
