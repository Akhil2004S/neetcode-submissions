// func twoSum(nums []int, target int) []int {
//     numbers := make(map[int]int)
//     for id, num := range nums {
//         numbers[num] = id
//     }

//     for id, num := range nums {
//         result := []int{}
//         requiredNum := target - num
//         if val, ok := numbers[requiredNum]; ok && id != val{
//             result = append(result, id, numbers[requiredNum])
//             return result
//         }
//     }
//     return []int{}
// }

func twoSum(nums []int, target int) []int {
    numbers := make(map[int]int)
    for id, num := range nums {
    numRequired := target - num
    if val, ok := numbers[numRequired]; ok {
        return []int{val, id}
    }
    numbers[num] = id
}
    return []int{0, 0}
}