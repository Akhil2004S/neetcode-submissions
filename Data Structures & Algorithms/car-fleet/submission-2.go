func carFleet(target int, position []int, speed []int) int {
	var time []float64
	positionToSpeed := make(map[int]int)
	for i, pos := range position {
		positionToSpeed[pos] = speed[i]
	}

	sort.Slice(position, func(i, j int) bool {
		return position[i] > position[j]
	})
	for _, pos := range position {
		currentTime := (float64(target-pos) / float64(positionToSpeed[pos]))
		if len(time) == 0 || currentTime > time[len(time)-1] {
			time = append(time, currentTime)
			continue
		}
	}

	return len(time)
}
