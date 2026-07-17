func minEatingSpeed(piles []int, h int) int {
	maxPile := piles[0]
	for _, v := range piles {
		if v > maxPile {
			maxPile = v
		}
	}
    start, end := 1, maxPile
    ans := end

    for start <= end {
        mid := start + (end-start)/2

        hours := 0
        for _, pile := range piles {
            hours += (pile + mid - 1) / mid
        }

        if hours <= h {
            ans = mid        
            end = mid - 1 
        } else {
            start = mid + 1 
        }
    }

    return ans
}