package main

import "fmt"

/**
移动零
给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。

请注意 ，必须在不复制数组的情况下原地对数组进行操作。

示例 1:

输入: nums = [0,1,0,3,12]
输出: [1,3,12,0,0]
示例 2:

输入: nums = [0]
输出: [0]
**/
// moveZeroes 将所有的 0 移动到数组的末尾，同时保持非零元素的相对顺序
func moveZeroes(nums []int) {
	if len(nums) == 0 {
		return
	}

	// 指向下一个非零元素的位置
	nonZeroIndex := 0

	// 遍历数组，找到所有非零元素并移动到前面
	for _, num := range nums {
		if num != 0 {
			nums[nonZeroIndex] = num
			nonZeroIndex++
		}
	}

	// 将剩余的位置填充为零
	for i := nonZeroIndex; i < len(nums); i++ {
		nums[i] = 0
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Println(nums) // 输出: [1, 3, 12, 0, 0]
}
