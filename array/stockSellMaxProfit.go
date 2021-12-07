package main

import "fmt"

func main() {

	/* 1st Approach: Brute Force: Use two for loops to find min and maximum number and profit will be difference between them
	Time Complexity: O(N*N)
	Space Complexity: O(1)
	*/

	/*
		Optimized Solution: Time Complexity: O(N)
		Space Complexity: O(1)
	*/
	// Maximum profit on buying and selling of stock O(N)
	// Find maximum sum of a subArray

	result := findMaxProfit(50, 20, 30, 40, -99)

	fmt.Println("Maximum profit from buying and selling a stock is", result)
}

func findMaxProfit(values ...int) int {
	minSellPrice := 9223372036854775807
	maxProfit := 0

	for _, v := range values {
		if v < minSellPrice {
			minSellPrice = v
		}
		profit := v - minSellPrice

		if profit > maxProfit {
			maxProfit = profit
		}

	}
	return maxProfit
}
