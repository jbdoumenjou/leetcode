package main

import "strconv"

// Given an integer n, return a string array answer (1-indexed) where:
//
// answer[i] == "FizzBuzz" if i is divisible by 3 and 5.
// answer[i] == "Fizz" if i is divisible by 3.
// answer[i] == "Buzz" if i is divisible by 5.
// answer[i] == i (as a string) if none of the above conditions are true.
//
// Time Complexity O(n)
// Space Complexity O(n)
func fizzBuzz(n int) []string {
	var result []string

	for i := 1; i <= n; i++ {
		result = append(result, value(i))
	}

	return result
}

func value(i int) string {
	if i%15 == 0 {
		return "FizzBuzz"
	}

	if i%3 == 0 {
		return "Fizz"
	}

	if i%5 == 0 {
		return "Buzz"
	}

	return strconv.Itoa(i)
}
