func productExceptSelf(nums []int) []int {
	// logic says we take product of everything before first, then take product of everything after.
	//[1, 2, 4, 6]
	res := make([]int, len(nums)) // [0,0,0,0] 
	res[0] = 1 // => [1,0,0,0]
	res[1] = nums[0]

	for i := 2; i < len(nums); i++ {
		product := nums[i-1] * res[i-1]
		res[i] = product
	}

	right := 1

	for j := len(res) - 2; j >= 0; j-- {
		right *= nums[j+1]
		res[j] *= right
	}

	return res
}
