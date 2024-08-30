# Array Algorithms Documentation

## Overview: What is an Array?
An **array** is a linear data structure where elements are stored in contiguous memory locations. All elements in an array are arranged sequentially, and each element can be accessed using its index. Arrays are widely used due to their simplicity and efficiency in storing and accessing data.

---

### Q: Majority Element
**Problem**: Given an array, find the majority element, which is the element that appears more than half the time.

**Solution**: Moore's Voting Algorithm.

```go
// Example Array
arr := []int{1, 2, 1, 1, 2, 1, 3}

// Steps in Moore's Voting Algorithm
// 1. Initialize a frequency counter and a candidate element.
// 2. Traverse the array and update the candidate and frequency.
// 3. The candidate at the end of the loop is the majority element if it exists.
```

---

### Q: Subarray Generation using Recursion
**Problem**: Generate all possible subarrays of a given array using recursion.

**Solution**:

```go
package main

import "fmt"

// Recursive function to generate subarrays
func subarr(arr, sub []int, idx int) {
	if idx == len(arr) {
		fmt.Println(sub)
		return
	}
	// Include the current element
	sub = append(sub, arr[idx])
	subarr(arr, sub, idx+1)

	// Exclude the current element and move to the next
	subarr(arr, sub[:len(sub)-1], idx+1)
}

func main() {
	// Example Array
	arr := []int{1, 2, 3}
	subarr(arr, []int{}, 0) // Start with an empty subarray and index 0
}
```

---

### Q: Maximum Sum Subarray
**Problem**: Find the maximum sum subarray within a given array.

**Solution**: Kadane's Algorithm.

```go
// Kadane's Algorithm to find the maximum sum subarray
func maxSumSub(arr []int) int {
    maxSum, currSum := 0, 0
    for _, value := range arr {
        currSum += value
        if currSum > maxSum {
            maxSum = currSum
        }
        if currSum < 0 {
            currSum = 0
        }
    }
    return maxSum
}
```

---

### Q: Best Time to Buy and Sell Stock - 1
**Problem**: Given an array where each element represents the price of a stock on a given day, find the maximum profit you can achieve by buying and selling one share of the stock.

**Solution**: 

1. **Brute Force**:
    - Compare each price with every subsequent price.
    - Time Complexity: O(N²), Space Complexity: O(1).

2. **Optimized - Using Right Maxima**:
    - Compute the right maximum for each element.
    - Compare the current price with the right maxima to calculate the profit.
    - Time Complexity: O(N), Space Complexity: O(N).

3. **Optimized - Using Two Variables**:
    - Track the lowest price so far and the maximum profit so far.
    - Time Complexity: O(N), Space Complexity: O(1).

```go
// Example Array
arr := []int{3, 1, 4, 8, 7, 2, 5}

// Optimized Approach
lowSoFar, maxProfit := arr[0], 0
for _, price := range arr {
    if price < lowSoFar {
        lowSoFar = price
    }
    profit := price - lowSoFar
    if profit > maxProfit {
        maxProfit = profit
    }
}
```

---

### Q: Buy and Sell Stock - 2
**Problem**: Maximize your profit with multiple buy-sell operations, but you must sell the stock before buying again.

**Solution**: Identify local minima and maxima.

```go
arr := []int{5, 2, 6, 1, 4, 7, 3, 6}

// Strategy: Add to profit whenever the price rises from one day to the next.
profit := 0
for i := 1; i < len(arr); i++ {
    if arr[i] > arr[i-1] {
        profit += arr[i] - arr[i-1]
    }
}
```

---

### Q: Rainwater Trapping Problem
**Problem**: Given the heights of buildings, determine how much rainwater can be trapped between the buildings.

**Solution**:

1. **Find the left maximum for each building**.
2. **Find the right maximum for each building**.
3. **Calculate the trapped water using**:
    - `min(leftMax, rightMax) - height`.

```go
arr := []int{3, 1, 2, 4, 0, 1, 3, 2}

// Left maxima array
leftMax := []int{3, 3, 3, 4, 4, 4, 4, 4}

// Right maxima array
rightMax := []int{4, 4, 4, 4, 3, 3, 3, 2}

// Calculate trapped water for each building
trappedWater := 0
for i := range arr {
    trappedWater += max(0, min(leftMax[i], rightMax[i]) - arr[i])
}
```

---

This documentation outlines the key concepts and algorithms related to arrays, including how to solve various array-related problems effectively. Each algorithm is implemented in Go, with examples and step-by-step explanations to help you understand and apply them.