import "log"

func evalRPN(tokens []string) int {
	var stack []int
	var secondVal int
	var firstVal int
	for _, token := range tokens {
		if token == "+" || token == "-" || token == "/" || token == "*" {
			secondVal = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			firstVal = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}
		switch token {
		case "+":
			stack = append(stack, firstVal+secondVal)
		case "-":
			stack = append(stack, firstVal-secondVal)
		case "/":
			stack = append(stack, firstVal/secondVal)
		case "*":
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