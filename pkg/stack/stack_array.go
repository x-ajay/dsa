package stack

// Stack represents a stack using an array.
type Stack struct {
	items []int
}

func New() *Stack {
	return &Stack{items: make([]int, 0)}
}

// Push adds an element to the top of the stack.
func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

// Pop removes and returns the top element of the stack. Panics if the stack is empty.
func (s *Stack) Pop() int {
	if s.IsEmpty() {
		panic("pop from empty stack")
	}
	lastIndex := len(s.items) - 1
	item := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return item
}

// Peek returns the top element of the stack without removing it. Panics if the stack is empty.
func (s *Stack) Peek() int {
	if s.IsEmpty() {
		panic("peek from empty stack")
	}
	return s.items[len(s.items)-1]
}

// IsEmpty checks if the stack is empty.
func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}
