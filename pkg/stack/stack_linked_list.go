package stack

// Node represents an element in the stack (linked list).
type Node struct {
	data int
	next *Node
}

// Stack represents the stack using a linked list.
type StackLL struct {
	top *Node
}

func NewLL() *StackLL {
	return &StackLL{}
}

// Push adds an element to the top of the stack.
func (s *StackLL) Push(data int) {
	newNode := &Node{data: data}
	newNode.next = s.top
	s.top = newNode
}

// Pop removes and returns the top element of the stack. Panics if the stack is empty.
func (s *StackLL) Pop() int {
	if s.IsEmpty() {
		panic("pop from empty stack")
	}
	poppedNode := s.top
	s.top = s.top.next
	return poppedNode.data
}

// Peek returns the top element of the stack without removing it. Panics if the stack is empty.
func (s *StackLL) Peek() int {
	if s.IsEmpty() {
		panic("peek from empty stack")
	}
	return s.top.data
}

// IsEmpty checks if the stack is empty.
func (s *StackLL) IsEmpty() bool {
	return s.top == nil
}
