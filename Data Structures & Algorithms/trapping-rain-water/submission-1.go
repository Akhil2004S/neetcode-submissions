func trap(height []int) int {
	left := 0
	right := len(height) - 1
	maxLeft := height[left]
	maxRight := height[right]
	var totalWaterTrapped int

	for left < right {
		for maxLeft <= maxRight && left < right {
			left++
			waterTrapped := maxLeft - height[left]
			if waterTrapped > 0 {
				totalWaterTrapped += waterTrapped
			}
			if height[left] > maxLeft {
				maxLeft = height[left]
			}
		}

		for maxRight < maxLeft && left < right {
			right--
			waterTrapped := maxRight - height[right]
			if waterTrapped > 0 {
				totalWaterTrapped += waterTrapped
			}
			if height[right] > maxRight {
				maxRight = height[right]
			}
		}
	}
	return totalWaterTrapped
}
