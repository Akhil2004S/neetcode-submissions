func dailyTemperatures(temperatures []int) []int {
	var output []int
	output = make([]int, len(temperatures))
	var stack []int

	for day, temperature := range temperatures {
		for len(stack) != 0 {
			top := stack[len(stack)-1]
			if temperatures[top] < temperature {
				output[top] = day - top
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, day)
				break
			}
		}
		if len(stack) == 0 {
			stack = append(stack, day)
		}
	}
	return output
}
