import(
	"log"
)

type MinStack struct {
	stack []int
}

func Constructor() MinStack {
	var stack []int
	return MinStack{
		stack,
	}
}

func (stack *MinStack) Push(val int) {
	stack.stack = append(stack.stack, val)
}

func (stack *MinStack) Pop() {
	if len(stack.stack)-1 >= 0 {
		stack.stack = stack.stack[:len(stack.stack)-1]
	} else {
		log.Fatal("Stack Underflow.")
	}
}

func (stack *MinStack) Top() int {
	top := stack.stack[len(stack.stack)-1]
	return top
}

func (stack *MinStack) GetMin() int {
	min := math.MaxInt64
	for _, element := range stack.stack {
		if element < min {
			min = element
		}
	}
	return min
}