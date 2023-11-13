package main

import (
	"slices"
)

// Given an integer array nums and an integer val,
// remove all occurrences of val in nums in-place.
// The order of the elements may be changed.
// Then return the number of elements in nums which are not equal to val.
//
// Consider the number of elements in nums which are not equal to val be k,
// to get accepted, you need to do the following things:
//
// Change the array nums such that the first k elements of nums contain the elements which are not equal to val.
// The remaining elements of nums are not important as well as the size of nums.
// Return k.

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	slices.Sort(nums)

	start := slices.Index(nums, val)
	if start == -1 {
		return len(nums)
	}

	part2 := nums[start:]
	end := slices.IndexFunc(part2, func(i int) bool {
		return i != val
	})

	if end == -1 {
		nums = nums[:start]
	} else {
		nums = append(nums[:start], nums[end+start:]...)
	}

	return len(nums)
}

func removeElementHidden(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	l := len(nums)
	for i, n := range nums {
		if n == val {
			// 0 <= val <= 100
			// The remaining elements of nums are not important as well as the size of nums.
			nums[i] = 101
			l--
		}
	}

	slices.Sort(nums)

	return l
}
