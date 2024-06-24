package main

import "fmt"

/**
盛最多水的容器
给定一个长度为 n 的整数数组 height 。有 n 条垂线，第 i 条线的两个端点是 (i, 0) 和 (i, height[i]) 。

找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。

返回容器可以储存的最大水量。

说明：你不能倾斜容器

输入：[1,8,6,2,5,4,8,3,7]
输出：49
解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
示例 2：

输入：height = [1,1]
输出：1
**/

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxWater := 0

	for left < right {
		// Calculate the area
		width := right - left
		high := min(height[left], height[right])
		currentArea := width * high
		// Update the maximum area found
		if currentArea > maxWater {
			maxWater = currentArea
		}
		// Move the pointer pointing to the shorter line
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxWater
}

// Helper function to find the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Test cases
	height1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	height2 := []int{1, 1}

	fmt.Println("The maximum area for height1 is:", maxArea(height1)) // Output: 49
	fmt.Println("The maximum area for height2 is:", maxArea(height2)) // Output: 1
}
