Certainly! Here's how you can implement stack operations using arrays and linked lists in Go with `panic` instead of returning errors when trying to pop or peek from an empty stack.

# Stack Implementation in Go

## Stack Using Array

### Overview

An array-based stack uses an array to store stack elements, with an integer tracking the top element's index. If you attempt to pop or peek from an empty stack, the program will panic.

### Stack Structure

```go
package main

import "fmt"

// Stack represents a stack using an array.
type Stack struct {
    items []int
}
```

### Operations

1. **Push**: Adds an element to the top of the stack.
2. **Pop**: Removes and returns the top element from the stack. Panics if the stack is empty.
3. **Peek**: Returns the top element without removing it. Panics if the stack is empty.
4. **IsEmpty**: Checks if the stack is empty.

```go
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
```

### Example Usage

```go
func main() {
    stack := &Stack{}

    stack.Push(10)
    stack.Push(20)
    stack.Push(30)

    fmt.Println(stack.Pop()) // Output: 30
    fmt.Println(stack.Peek()) // Output: 20
    fmt.Println(stack.IsEmpty()) // Output: false

    // Uncommenting the following line will cause a panic
    // stack.Pop()
    // stack.Pop()
    // stack.Pop() // This will panic: pop from empty stack
}
```

## Stack Using Linked List

### Overview

A linked list-based stack uses nodes, where each node contains a value and a pointer to the next node. The top of the stack is the most recently added node. Attempting to pop or peek from an empty stack will cause the program to panic.

### Stack Structure

```go
package main

import "fmt"

// Node represents an element in the stack (linked list).
type Node struct {
    data int
    next *Node
}

// Stack represents the stack using a linked list.
type Stack struct {
    top *Node
}
```

### Operations

1. **Push**: Adds an element to the top of the stack.
2. **Pop**: Removes and returns the top element from the stack. Panics if the stack is empty.
3. **Peek**: Returns the top element without removing it. Panics if the stack is empty.
4. **IsEmpty**: Checks if the stack is empty.

```go
// Push adds an element to the top of the stack.
func (s *Stack) Push(data int) {
    newNode := &Node{data: data}
    newNode.next = s.top
    s.top = newNode
}

// Pop removes and returns the top element of the stack. Panics if the stack is empty.
func (s *Stack) Pop() int {
    if s.IsEmpty() {
        panic("pop from empty stack")
    }
    poppedNode := s.top
    s.top = s.top.next
    return poppedNode.data
}

// Peek returns the top element of the stack without removing it. Panics if the stack is empty.
func (s *Stack) Peek() int {
    if s.IsEmpty() {
        panic("peek from empty stack")
    }
    return s.top.data
}

// IsEmpty checks if the stack is empty.
func (s *Stack) IsEmpty() bool {
    return s.top == nil
}
```

### Example Usage

```go
func main() {
    stack := &Stack{}

    stack.Push(10)
    stack.Push(20)
    stack.Push(30)

    fmt.Println(stack.Pop()) // Output: 30
    fmt.Println(stack.Peek()) // Output: 20
    fmt.Println(stack.IsEmpty()) // Output: false

    // Uncommenting the following lines will cause a panic
    // stack.Pop()
    // stack.Pop()
    // stack.Pop() // This will panic: pop from empty stack
}
```

## Summary

- **Stack Using Array**: Uses an array to store stack elements. If the stack is empty when attempting to pop or peek, the program will panic.
- **Stack Using Linked List**: Uses a linked list to implement a dynamic stack. Panics occur on pop or peek operations when the stack is empty.

In both implementations, using `panic` ensures that the code will immediately stop if you attempt to perform invalid operations (like popping from an empty stack), which can help catch bugs early during development.