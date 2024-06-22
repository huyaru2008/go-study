package main

import (
	"fmt"
	"sort"
)

/**
最长连续序列
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。

请你设计并实现时间复杂度为 O(n) 的算法解决此问题。



示例 1：

输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。
示例 2：

输入：nums = [0,3,7,2,5,8,4,6,0,1]
输出：9

**/

func longestConsecutive(nums []int) int {
	// 先排序
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	if len(nums) == 0 {
		return 0
	}
	numlen := 1
	// for循环便利，定义临时变量记录连续数字长度
	num := 1
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] == nums[i] {
			continue
		}

		if nums[i+1]-nums[i] == 1 {
			num++
		} else {
			num = 1
		}

		if numlen < num {
			numlen = num
		}

	}
	return numlen

}

func longestConsecutive2(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	// 基于map实现
	numMap := make(map[int]struct{})
	for _, num := range nums {
		numMap[num] = struct{}{}
	}

	longnum := 1
	for num, _ := range numMap {
		_, ok := numMap[num-1]
		if !ok {
			startNum := num
			r := 1
			for {
				if _, numOk := numMap[startNum+1]; numOk {
					startNum++
					r++
				} else {
					break
				}

			}

			if longnum < r {
				longnum = r
			}
		}
	}

	return longnum
}

func main() {
	var nums = []int{1, 2, 0, 1}
	fmt.Println(longestConsecutive(nums))
	fmt.Println(longestConsecutive2(nums))
}
