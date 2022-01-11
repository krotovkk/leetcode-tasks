package main

import (
	"fmt"
	"github.com/krotovkk/leetcode-tasks/ds/array/intersection"
)

func main() {
	nums1 := []int{1, 2, -5, 1, 5, 2, -1}
	nums2 := []int{1, 2, 2, 2, 100, -1}
	//fmt.Println(contains_duplicate.ContainsDuplicate(nums))
	//fmt.Println(maximum_subarray.MaxSubArray(nums))
	fmt.Println(intersection.Intersect(nums1, nums2))
}
