func threeSum(nums []int) [][]int {
	for i := range nums {
		for j := i; j > 0 && nums[j-1] > nums[j]; j-- {
			nums[j], nums[j-1] = nums[j-1], nums[j]
		}
	}

	var triplets [][]int
	for i, num := range nums {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		arrayToLook := nums[i+1:]
		target := -num
		left := 0
		right := len(arrayToLook) - 1
		for left < right {
			sum := arrayToLook[left] + arrayToLook[right]
			if sum == target {
				triplet := []int{num, arrayToLook[left], arrayToLook[right]}
				triplets = append(triplets, triplet)
				left++
				right--

				for left < right && arrayToLook[left] == arrayToLook[left-1] {
					left++
				}
				for left < right && arrayToLook[right] == arrayToLook[right+1] {
					right--
				}
			}
			if sum < target {
				left++
			} else if sum > target {
				right--
			}
		}
	}
	return triplets
}
