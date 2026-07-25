func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }
    ms,mt := make(map[rune]int), make(map[rune]int)

    for _, v := range s {
        ms[v]++
    }
    for _, v := range t {
        mt[v]++
    }

    for _, v := range s {
        if ms[v] != mt[v] {
            return false
        }
    }

    return true 
}
