package main

// You are given an m x n integer grid accounts where accounts[i][j] is the amount of money the ith customer has in the jth bank.
// Return the wealth that the richest customer has.
//
// A customer's wealth is the amount of money they have in all their bank accounts.
// The richest customer is the customer that has the maximum wealth.
//
// Time complexity: O(n*m)
// Space complexity: O(1)
func maximumWealth(accounts [][]int) int {
	max := 0
	for _, account := range accounts {
		sum := 0
		for _, money := range account {
			sum += money
		}
		if sum > max {
			max = sum
		}
	}

	return max
}
