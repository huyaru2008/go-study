package main

import (
	"fmt"
	"sort"
)

/**

三数之和
给你一个整数数组 nums ，判断是否存在三元组 [nums[i], nums[j], nums[k]] 满足 i != j、i != k 且 j != k ，同时还满足 nums[i] + nums[j] + nums[k] == 0 。请

你返回所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。


示例 1：

输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
解释：
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0 。
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0 。
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0 。
不同的三元组是 [-1,0,1] 和 [-1,-1,2] 。
注意，输出的顺序和三元组的顺序并不重要。
示例 2：

输入：nums = [0,1,1]
输出：[]
解释：唯一可能的三元组和不为 0 。
示例 3：

输入：nums = [0,0,0]
输出：[[0,0,0]]
解释：唯一可能的三元组和为 0 。
**/

// 方法一：排序 + 双指针
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	sort.Ints(nums)
	var targetNums [][]int
	for i := 0; i < len(nums)-2; i++ {
		// 这里避免 i 的重复
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		start, end := i+1, len(nums)-1
		for start < end {
			// 避免 start 和 end 的重复数字
			if start > i+1 && nums[start] == nums[start-1] {
				start++
				continue
			}
			if end < len(nums)-1 && nums[end] == nums[end+1] {
				end--
				continue
			}
			sum := nums[i] + nums[start] + nums[end]
			if sum == 0 {
				targetNums = append(targetNums, []int{nums[i], nums[start], nums[end]})
				start++
				end--
			} else if sum > 0 {
				end--
			} else {
				start++
			}
		}
	}
	return targetNums
}

func main() {
	var nums = []int{-1, 0, 1, 2, -1, -4}
	var nums2 = []int{0, 0, 0}
	var nums3 = []int{0, 1, 1}
	fmt.Println(threeSum(nums))  // 输出: [[-1 -1 2] [-1 0 1]]
	fmt.Println(threeSum(nums2)) // 输出: [[0 0 0]]
	fmt.Println(threeSum(nums3)) // 输出: []

}
