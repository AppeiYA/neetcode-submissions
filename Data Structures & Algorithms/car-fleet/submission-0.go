type car struct {
	pos int
	arrTime float64
}

func carFleet(target int, position []int, speed []int) int {
	n := len(position)
	if n == 0 {
		return 0
	}

	cars := make([]car, n)
	for k, v := range position {
		cars[k] = car{
			pos: v,
			arrTime: float64(target - v)/float64(speed[k]),
		}
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos > cars[j].pos
	})

	fleets := 0
	var maxTime float64 = 0

	for _, v := range cars {
		if v.arrTime > maxTime {
			maxTime = v.arrTime
			fleets++
		}
	}

	return fleets
}
