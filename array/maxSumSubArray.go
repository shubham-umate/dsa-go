package main

import "fmt"

func main() {
	// Kadane's Algorithm O(N)
	// Find maximum sum of a subArray

	result := findMaxSumSubArray(10, 20, 30, 40, -99)

	fmt.Println("Maximum sum of a subarray is", result)

}

func findMaxSumSubArray(values ...int) int {
	maxSum := -9223372036854775808
	currentSum := 0

	for _, v := range values {
		currentSum += v
		if currentSum > maxSum {
			maxSum = currentSum
		}
		if currentSum < 0 {
			currentSum = 0
		}
	}

	return maxSum
}
