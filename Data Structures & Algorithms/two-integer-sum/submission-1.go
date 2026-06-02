func twoSum(nums []int, target int) []int {
    numbers := make(map[int]int)
    for id, num := range nums {
        numbers[num] = id
    }

    for id, num := range nums {
        result := []int{}
        requiredNum := target - num
        if val, ok := numbers[requiredNum]; ok && id != val{
            result = append(result, id, numbers[requiredNum])
            return result
        }
    }
    return []int{}
}
