# Linked List in Go

A **Linked List** is a data structure that consists of a sequence of elements, where each element points to the next element in the sequence. Unlike arrays, linked lists do not require contiguous memory locations. This makes linked lists ideal for dynamic memory allocation, where the size of the data structure can change dynamically during runtime.

Each element in a linked list is called a **node**, and each node contains:
1. **Data**: The value or data the node holds.
2. **Pointer**: A reference to the next node in the list.

## Linked List Structure

```go
package main

import "fmt"

// Node represents an element in the linked list.
type Node struct {
    data int
    next *Node
}

// LinkedList represents the linked list itself.
type LinkedList struct {
    head *Node
}
```

### Text Diagram

Here's a simple diagram of how nodes are linked together in a linked list:

```
head -> [ data | next ] -> [ data | next ] -> [ data | next ] -> nil
```

Where `head` is a pointer to the first node, and each node points to the next node in the list, with the last node pointing to `nil` to indicate the end of the list.

## Functions for Linked List Operations

### 1. Insert a Node at the End

The `Insert` function adds a new node with a given value to the end of the linked list.

```go
// Insert adds a new node with the given data to the end of the list.
func (ll *LinkedList) Insert(data int) {
    newNode := &Node{data: data}
    if ll.head == nil {
        ll.head = newNode
        return
    }

    current := ll.head
    for current.next != nil {
        current = current.next
    }
    current.next = newNode
}
```

**Explanation**:
- If the list is empty (`head == nil`), the new node becomes the head.
- Otherwise, traverse to the end of the list and set the `next` pointer of the last node to the new node.

### 2. Insert a Node at a Specific Index

The `InsertAtIndex` function inserts a new node with a given value at a specified index in the linked list.

```go
// InsertAtIndex inserts a new node with the given data at the specified index.
func (ll *LinkedList) InsertAtIndex(data int, index int) error {
    newNode := &Node{data: data}
    if index == 0 {
        newNode.next = ll.head
        ll.head = newNode
        return nil
    }

    current := ll.head
    for i := 0; i < index-1; i++ {
        if current == nil {
            return fmt.Errorf("index out of range")
        }
        current = current.next
    }
    newNode.next = current.next
    current.next = newNode
    return nil
}
```

**Explanation**:
- If the index is `0`, the new node is inserted at the beginning.
- Otherwise, traverse to the node just before the specified index and adjust the pointers to insert the new node.

### 3. Delete a Node at a Specific Index

The `DeleteAtIndex` function removes the node at the specified index from the linked list.

```go
// DeleteAtIndex removes the node at the specified index.
func (ll *LinkedList) DeleteAtIndex(index int) error {
    if ll.head == nil {
        return fmt.Errorf("list is empty")
    }

    if index == 0 {
        ll.head = ll.head.next
        return nil
    }

    current := ll.head
    for i := 0; i < index-1; i++ {
        if current.next == nil {
            return fmt.Errorf("index out of range")
        }
        current = current.next
    }

    if current.next == nil {
        return fmt.Errorf("index out of range")
    }

    current.next = current.next.next
    return nil
}
```

**Explanation**:
- If the index is `0`, the head is moved to the next node.
- Otherwise, traverse to the node just before the specified index and adjust the pointers to remove the node.

### 4. Get the Length of the Linked List

The `Length` function returns the number of nodes in the linked list.

```go
// Length returns the number of nodes in the list.
func (ll *LinkedList) Length() int {
    length := 0
    current := ll.head
    for current != nil {
        length++
        current = current.next
    }
    return length
}
```

**Explanation**:
- Traverse the list, incrementing a counter for each node until the end of the list is reached.

### 5. Print the Linked List

The `Print` function prints all the elements in the linked list.

```go
// Print displays all the nodes in the list.
func (ll *LinkedList) Print() {
    current := ll.head
    for current != nil {
        fmt.Printf("%d -> ", current.data)
        current = current.next
    }
    fmt.Println("nil")
}
```

**Explanation**:
- Traverse the list and print the data of each node, followed by an arrow (`->`). Finally, print `nil` to indicate the end of the list.

### 6. Reverse the Linked List

The `Reverse` function reverses the linked list so that the last node becomes the first and vice versa.

```go
// Reverse reverses the linked list.
func (ll *LinkedList) Reverse() {
    var prev, next *Node
    current := ll.head

    for current != nil {
        next = current.next
        current.next = prev
        prev = current
        current = next
    }

    ll.head = prev
}
```

**Explanation**:
- The function iterates through the list, reversing the direction of each node's `next` pointer. Finally, the `head` is updated to the last node, which is now the first.

### Example Usage

```go
func main() {
    ll := LinkedList{}

    ll.Insert(10)
    ll.Insert(20)
    ll.Insert(30)
    ll.Print() // Output: 10 -> 20 -> 30 -> nil

    ll.InsertAtIndex(15, 1)
    ll.Print() // Output: 10 -> 15 -> 20 -> 30 -> nil

    ll.DeleteAtIndex(2)
    ll.Print() // Output: 10 -> 15 -> 30 -> nil

    fmt.Println("Length:", ll.Length()) // Output: Length: 3

    ll.Reverse()
    ll.Print() // Output: 30 -> 15 -> 10 -> nil
}
```

## Summary

- **Linked List**: A dynamic data structure consisting of nodes, each containing data and a pointer to the next node.
- **Insert**: Adds a node to the end of the list.
- **Insert at Index**: Adds a node at a specified position in the list.
- **Delete at Index**: Removes a node from a specified position in the list.
- **Length**: Returns the total number of nodes in the list.
- **Print**: Displays the elements of the list.
- **Reverse**: Reverses the order of the nodes in the list.

This structure is highly versatile and is useful in situations where you need dynamic memory allocation with efficient insertions and deletions.