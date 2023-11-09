package main

import "slices"

// Given an array nums. We define a running sum of an array as runningSum[i] = sum(nums[0]…nums[i]).
//
// Return the running sum of nums.
// Time complexity: O(n)
// Space complexity: O(1)
// /!\ This function mutates the input array
func runningSum(nums []int) []int {
	for i := 1; i < len(nums); i++ {
		nums[i] = nums[i] + nums[i-1]
	}

	return nums
}

func runningSum_clone(nums []int) []int {
	result := slices.Clone(nums)
	for i := 1; i < len(result); i++ {
		result[i] = result[i] + result[i-1]
	}

	return result
}
