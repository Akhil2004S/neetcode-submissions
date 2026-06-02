func hasDuplicate(nums []int) bool {
	numbers := make(map[int]bool)
    for _, num := range nums {
		if _, ok := numbers[num]; ok {
			return true
		}
		numbers[num] = true
	}
	return false
}
