package main

func twoSum(nums []int, target int) []int {
	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] != target {
				continue
			}
			return []int{i, j}
		}
	}
	return []int{}
}
