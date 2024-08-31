package linkedlist

import "fmt"

type LinkedList struct {
	Head *Node
}

type Node struct {
	Data int
	Next *Node
}

func New(data int) *LinkedList {
	return &LinkedList{
		Head: &Node{
			Data: data,
		},
	}
}

func (ll *LinkedList) Insert(data int) *LinkedList {
	if ll == nil {
		return New(data)
	}

	head := ll.Head
	for head.Next != nil {
		head = head.Next
	}

	node := &Node{Data: data}
	head.Next = node
	return ll
}

func (ll *LinkedList) InsertAt(idx int, data int) *LinkedList {
	if ll == nil {
		return New(data)
	}

	node := &Node{Data: data}
	if idx == 0 {
		curr := ll.Head
		ll.Head = node
		node.Next = curr
		return ll
	}

	curr := ll.Head
	for idx-1 != 0 {
		curr = curr.Next
		idx--
	}

	node.Next = curr.Next
	curr.Next = node

	return ll
}

func (ll *LinkedList) DeleteAt(idx int) *LinkedList {
	curr := ll.Head
	if idx == 0 {
		ll.Head = ll.Head.Next
		return ll
	}

	for idx-1 != 0 {
		curr = curr.Next
		idx--
	}
	curr.Next = curr.Next.Next
	return ll
}

func (ll *LinkedList) Len() int {
	curr := ll.Head
	var count int
	for curr != nil {
		count++
		curr = curr.Next
	}
	return count
}

func (ll *LinkedList) Print() *LinkedList {
	curr := ll.Head
	for curr != nil {
		fmt.Printf("%d->", curr.Data)
		curr = curr.Next
	}
	fmt.Printf("NILL\n")
	return ll
}

func (ll *LinkedList) Reverse() *LinkedList {
	curr := ll.Head
	var prev *Node
	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	ll.Head = prev

	return ll
}

func fromNodeReverse(head *Node) *Node {
	if head == nil {
		return nil
	}

	var prev *Node
	curr := head
	for curr != nil {
		temp := curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
	}
	return prev
}

func (ll *LinkedList) IsPalindrome() bool {
	head := ll.Head

	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	lside := head
	rside := fromNodeReverse(slow)

	for lside != nil && rside != nil {
		if lside.Data != rside.Data {
			return false
		}
		lside = lside.Next
		rside = rside.Next
	}
	return true
}

func (ll *LinkedList) IsCycle() bool {
	if ll == nil {
		return false
	}
	slow := ll.Head
	fast := ll.Head
	for fast.Next != nil && fast.Next.Next != nil {
		if fast == slow {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return true
}
