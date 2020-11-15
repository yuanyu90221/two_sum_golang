package two_sum_golang

func twoSum(nums []int, target int) []int {
	results := []int{}
	valueMap := make(map[int]int)
	for idx, val := range nums {
		value, ok := valueMap[target-nums[idx]]
		if ok && value != idx {
			if idx < value {
				return []int{idx, value}
			} else {
				return []int{value, idx}
			}
		} else {
			valueMap[val] = idx
		}
	}
	return results
}
