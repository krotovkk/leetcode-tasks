package contains_duplicate

func ContainsDuplicate(nums []int) bool {
	numbersMap := map[int]int{}
	for _, num := range nums {
		if _, ok := numbersMap[num]; ok {
			return true
		} else {
			numbersMap[num] = 1
		}
	}

	return false
}
