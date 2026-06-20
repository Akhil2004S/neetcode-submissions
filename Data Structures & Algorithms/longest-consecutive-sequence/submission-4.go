func longestConsecutive(nums []int) int {
	var count int
	count = 1
	var largestSequence int
	var done bool
	numbers := make(map[int]struct{})
	for _, num := range nums {
		numbers[num] = struct{}{}
	}

	for _, num := range nums {
		count = 1
		done = false
		currentNum := num
		_, ok := numbers[num-1]
		if ok {
			continue
		}
		for !done {
			if _, ok = numbers[currentNum+1]; ok {
				count++
				currentNum = currentNum + 1
			} else {
				done = true
			}
		}
		if count > largestSequence {
			largestSequence = count
		}
	}
	return largestSequence
}
