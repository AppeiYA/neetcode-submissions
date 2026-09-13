func twoSum(nums []int, target int) []int {
	for j := 1; j < len(nums); j++ {
		for i := 0; i < j; i++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{0,0}
}
