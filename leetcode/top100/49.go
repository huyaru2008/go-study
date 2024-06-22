package main

import (
	"fmt"
	"sort"
)

/**
https://leetcode.cn/problems/group-anagrams/?envType=study-plan-v2&envId=top-100-liked
字母异位词分组
给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。

字母异位词 是由重新排列源单词的所有字母得到的一个新单词。


示例 1:

输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
示例 2:

输入: strs = [""]
输出: [[""]]
示例 3:

输入: strs = ["a"]
输出: [["a"]]

**/

func groupAnagrams(strs []string) [][]string {
	strMap := make(map[string][]string)
	for _, str := range strs {
		sortStr := sortString(str)
		strMap[sortStr] = append(strMap[sortStr], str)
	}

	result := make([][]string, 0)
	for _, str := range strMap {
		result = append(result, str)
	}
	return result
}

// sortString 对字符串中的字符进行排序并返回排序后的字符串
func sortString(s string) string {
	str := []rune(s)
	sort.Slice(str, func(i, j int) bool {
		return str[i] < str[j]
	})
	return string(str)
}

func main() {
	str := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	result := groupAnagrams(str)
	fmt.Printf("result: %v\n", result)
}
