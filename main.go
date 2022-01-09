package main

import (
	"fmt"
	"github.com/krotovkk/leetcode-tasks/ds/array/contains_duplicate"
	"github.com/krotovkk/leetcode-tasks/ds/array/maximum_subarray"
)

func main() {
	nums := []int{1, 2, -5, 1, 5, 2, -1}
	fmt.Println(contains_duplicate.ContainsDuplicate(nums))
	fmt.Println(maximum_subarray.MaxSubArray(nums))
}
