func groupAnagrams(strs []string) [][]string {
	equal := func(a, b string) bool {
		ra := []rune(a)
		rb := []rune(b)

		if len(ra) != len(rb) {
			return false
		}

		count := make(map[rune]int)

		for _, r := range ra {
			count[r]++
		}

		for _, r := range rb {
			count[r]--
			if count[r] < 0 {
				return false
			}
		}

		return true
	}

	var anagrams [][]string

	var build func([]string)

	build = func(list []string) {
		if len(list) == 0 {
			return
		}

		pivot := list[0]

		group := []string{pivot}
		var remaining []string

		for i := 1; i < len(list); i++ {
			if equal(pivot, list[i]) {
				group = append(group, list[i])
			} else {
				remaining = append(remaining, list[i])
			}
		}

		anagrams = append(anagrams, group)

		build(remaining)
	}

	build(strs)

	return anagrams
}