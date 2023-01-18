package main

import "fmt"

/*
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出 和为目标值 target  的那 两个 整数，并返回它们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。

你可以按任意顺序返回答案。

 

示例 1：

输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
示例 2：

输入：nums = [3,2,4], target = 6
输出：[1,2]
示例 3：

输入：nums = [3,3], target = 6

来源：力扣（LeetCode）
链接：https://leetcode.cn/problems/two-sum
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

func twoSum(nums []int, target int) []int {
	numMap := make(map[int][]int, 0)

	for index, num := range nums {
		if  _, ok := numMap[num]; !ok {
			numMap[num] = make([]int, 0)
		}
		numMap[num] = append(numMap[num], index)
	}

	//fmt.Printf("%+v \r\n", numMap)

	for index, num := range nums {
		if jIndexs, ok := numMap[target - num]; ok {
			for _, jIndex := range jIndexs {
				if jIndex == index {
					continue
				}
				return []int {index, jIndex}
			}
		}
	}

	return nil
}


func main() {
	//nums := []int {2, 7, 11, 15}
	//target := 9

	//nums := []int {3, 2, 4}
	//target := 6

	nums := []int {3, 3}
	target := 6
	fmt.Println(twoSum(nums, target))
}
