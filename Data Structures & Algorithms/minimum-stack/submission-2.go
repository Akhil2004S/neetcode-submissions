import(
	"log"
)

type MinStack struct {
	stack    []int
	minStack []int
}

func Constructor() MinStack {
	var stack []int
	var minStack []int
	return MinStack{
		stack,
		minStack,
	}
}

func (stack *MinStack) Push(val int) {
	stack.stack = append(stack.stack, val)
	if len(stack.minStack) == 0 || val <= stack.GetMin() {
		stack.minStack = append(stack.minStack, val)
	} else {
		stack.minStack = append(stack.minStack, stack.GetMin())
	}
}

func (stack *MinStack) Pop() {
	if len(stack.stack)-1 >= 0 {
		stack.stack = stack.stack[:len(stack.stack)-1]
		stack.minStack = stack.minStack[:len(stack.minStack)-1]
	} else {
		log.Fatal("Stack Underflow.")
	}
}

func (stack *MinStack) Top() int {
	top := stack.stack[len(stack.stack)-1]
	return top
}

func (stack *MinStack) GetMin() int {
	min := stack.minStack[len(stack.minStack)-1]
	return min
}