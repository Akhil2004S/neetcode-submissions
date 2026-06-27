import "log"

func evalRPN(tokens []string) int {
	var stack []int
	for _, token := range tokens {
		switch token {
		case "+":
			secondVal := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			firstVal := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, firstVal+secondVal)
		case "-":
			secondVal := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			firstVal := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, firstVal-secondVal)
		case "/":
			secondVal := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			firstVal := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, firstVal/secondVal)
		case "*":
			secondVal := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			firstVal := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, firstVal*secondVal)
		default:
			element, err := strconv.Atoi(token)
			if err != nil {
				log.Fatal(err)
			}
			stack = append(stack, element)
		}
	}
	return stack[len(stack)-1]
}