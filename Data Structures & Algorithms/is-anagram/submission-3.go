import "maps"
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	ms := make(map[byte]int)
	mt := make(map[byte]int)

	for i := 0; i < len(s); i++ {
		ms[s[i]]++
		mt[t[i]]++
	}

	if !maps.Equal(ms, mt) {
		return false
	}

	// for key, vs := range ms {
	// 	if vt, ok := mt[key]; !ok || vs != vt {
	// 		return false
	// 	}
	// }

	return true
}
