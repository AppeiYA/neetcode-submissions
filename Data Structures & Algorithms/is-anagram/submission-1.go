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

    for k, v := range ms {
        if mt_v, ok := mt[k]; !ok || mt_v != v {
            return false
        }
    }

    return true 
}
