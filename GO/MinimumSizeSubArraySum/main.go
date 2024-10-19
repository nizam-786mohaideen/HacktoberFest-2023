package main

import (
	"fmt"
	"math"
	"strings"
)

/*
*
Problem Statement -
Given an array of positive integers nums and a positive integer target, return the minimal length of a subarray
whose sum is greater than or equal to target.
If there is no such subarray, return 0 instead.
*/

func minSubArrayLen(target int, nums []int) int {
	// Initializing minLength to a large value
	minLength := math.MaxInt64

	// Initialize start pointer to 0 and totalSum to 0
	start, totalSum := 0, 0

	for end := 0; end < len(nums); end++ {
		totalSum += nums[end] // Add current number to the total sum

		// Check if we can shrink the window from the left
		for totalSum >= target {
			currSubArrSize := (end + 1) - start // Current subarray size
			minLength = min(minLength, currSubArrSize) // Update min length
			totalSum -= nums[start] // Remove the leftmost element
			start++ // Move the start pointer to the right
		}
	}

	if minLength != math.MaxInt64 {
		return minLength // Return the minimal length found
	}
	return 0 // Return 0 if no valid subarray exists
}

// Helper function to find the minimum of two integers
func min(v1, v2 int) int {
	if v1 < v2 {
		return v1
	}
	return v2
}

func main() {
	targets := []int{
		7,
		4,
		11,
		10,
		5,
		15,
	}
	numsList := [][]int{
		{2, 3, 1, 2, 4, 3},
		{1, 4, 4},
		{1, 1, 1, 1, 1, 1, 1, 1},
		{1, 2, 3, 4},
		{1, 2, 1, 3},
		{5, 4, 9, 8, 11, 3, 7, 12, 15, 44},
	}
	for i, target := range targets {
		result := minSubArrayLen(target, numsList[i])
		fmt.Printf("%d.\tminSubArrayLen(%d, %s): %d\n", i+1, target,
			strings.Replace(fmt.Sprint(numsList[i]), " ", ", ", -1), result)
		fmt.Printf("%s\n", strings.Repeat("-", 100))
	}
}
