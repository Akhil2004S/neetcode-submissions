func carFleet(target int, position []int, speed []int) int {
	var time []float64
	indices := make([]int, len(position))

	for i := range indices {
		indices[i] = i
	}

	sort.Slice(indices, func(i, j int) bool {
		return position[indices[i]] > position[indices[j]]
	})
	for _, idx := range indices {
		currentTime := (float64(target-position[idx]) / float64(speed[idx]))
		if len(time) == 0 || currentTime > time[len(time)-1] {
			time = append(time, currentTime)
			continue
		}
	}
	return len(time)
}
