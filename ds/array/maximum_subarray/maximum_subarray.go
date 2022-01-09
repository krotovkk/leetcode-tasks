package maximum_subarray

import "math"

func MaxSubArray(nums []int) int {
	current := 0
	max := nums[0]
	for _, num := range nums {
		current += num
		max = int(math.Max(float64(current), float64(max)))
		if current < 0 {
			current = 0
			continue
		}
	}

	return max
}
