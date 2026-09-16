func longestConsecutive(nums []int) int {
	m := map[int]bool{}

	for _, num := range nums {
		m[num] = true
	}
	
	longest := 0

	for key := range m {
		if !m[key-1] {
			length := 1
			for {
				if m[key+length] {
					length++
				}else {
					break
				}
			}

			if length > longest {
				longest = length
			}
		}
	}

	return longest
}
