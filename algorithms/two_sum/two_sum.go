package main

func twoSum(nums []int, target int) []int {
	occurrences := map[int]int{}

	for i, num := range nums {
		occurrences[num] = i
	}

	for i1 := 0; i1 < len(nums) - 1; i1++ {
		rem := target - nums[i1]
		if i2, ok := occurrences[rem]; ok && i1 != i2 {
			return []int{i1, i2}
		}
	}

	return []int{}
}