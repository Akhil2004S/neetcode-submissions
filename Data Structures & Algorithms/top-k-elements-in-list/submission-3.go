func topKFrequent(nums []int, k int) []int {
	numbers := make(map[int]int)
	maxVal := -1
	var result []int
	for _, num := range nums {
		numbers[num]++
	}

	for _, value := range numbers {
		if value > maxVal {
			maxVal = value
		}
	}

	topK := make([][]int, maxVal+1)
	for key, value := range numbers {
		topK[value] = append(topK[value], key)
	}

	fmt.Println(topK)
	i := 0
	j := 0
	for i < k && len(topK)-j-1 >= 0 {
		if topK[len(topK)-j-1] != nil {
			result = append(result, topK[len(topK)-j-1]...)
			i += len(topK[len(topK)-j-1])
		}
		j++
	}
	fmt.Println(result)
	return result
}