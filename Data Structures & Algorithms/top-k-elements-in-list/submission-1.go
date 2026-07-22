func topKFrequent(nums []int, k int) []int {
    freq := make(map[int]int)
    for _, v := range nums {
        freq[v]++
    }

    buckets := make([][]int, len(nums)+1)
    for key, count := range freq {
        buckets[count] = append(buckets[count], key)
    }

    result := []int{}
    for i := len(buckets) - 1; i >= 0 && len(result) < k; i-- {
        result = append(result, buckets[i]...)
    }

    return result[:k]
}