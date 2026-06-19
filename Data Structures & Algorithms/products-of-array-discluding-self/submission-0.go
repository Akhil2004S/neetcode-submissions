func productExceptSelf(nums []int) []int {
	output := make([]int, len(nums))
	left := make([]int, len(nums))
	right := make([]int, len(nums))
	for i := range nums {
		if i == 0 {
			left[i] = 1
		} else {
			left[i] = left[i-1] * nums[i-1]
		}
	}

	i := len(nums) - 1
	for i != -1 {
		if i == len(nums)-1 {
			right[i] = 1
		} else {
			right[i] = right[i+1] * nums[i+1]
		}
		i--
	}

	for i := range nums {
		output[i] = left[i] * right[i]
	}
	fmt.Println(output)
	return output
}
