func twoSum(numbers []int, target int) []int {
	var output []int
	left := 0
	right := len(numbers) - 1

	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			output = append(output, left+1, right+1)
			break
		}
		if sum > target {
			right--
		} else if sum < target {
			left++
		}
	}
	return output
}
