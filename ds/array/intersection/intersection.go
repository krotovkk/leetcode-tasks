package intersection

type Counter map[int]int

func (c Counter) fillByArray(nums []int) {
	for _, num := range nums {
		c[num]++
	}
}

func min(num1, num2 int) int {
	if num1 < num2 {
		return num1
	}

	return num2
}

func Intersect(nums1 []int, nums2 []int) []int {
	counter1 := Counter{}
	counter2 := Counter{}
	var result []int

	counter1.fillByArray(nums1)
	counter2.fillByArray(nums2)

	for num, count1 := range counter1 {
		if count2, ok := counter2[num]; ok {
			for i := 0; i < min(count1, count2); i++ {
				result = append(result, num)
			}
		}
	}

	return result
}
