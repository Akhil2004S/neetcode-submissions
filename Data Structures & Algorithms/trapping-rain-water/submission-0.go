func trap(height []int) int {
	var left []int
	right := make([]int, len(height))
	maxLeft := -1
	maxRight := -1

	for _, h := range height {
		if h > maxLeft {
			maxLeft = h
		}
		left = append(left, maxLeft)
	}

	for i := len(height) - 1; i >= 0; i-- {
		if height[i] > maxRight {
			maxRight = height[i]
		}
		right[i] = maxRight
	}

	var totalWaterTrapped int
	for i, h := range height {
		length := min(left[i], right[i])
		waterTrapped := length - h
		if waterTrapped > 0 {
			totalWaterTrapped += waterTrapped
		}
	}
	fmt.Println(totalWaterTrapped)
	return totalWaterTrapped
}
