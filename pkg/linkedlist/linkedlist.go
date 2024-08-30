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
	//     [v,ptr]->[v,ptr]->[v,ptr]->[v,nil]->
	//prev-nil   curr	 next
	// 			prev     curr    next
	//                   prev     curr  next

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

/*
	fmt.Println("Creating 10->20->30")
	ll := linkedlist.New(10).
		Insert(20).
		Insert(30).
		Print()

	fmt.Println("Inserting 9 at 0 and 11 at 2")
	ll = ll.InsertAt(0, 9).
		InsertAt(2, 11).
		Print()
	fmt.Println("Deleting at idx 4 ")
	ll = ll.DeleteAt(4).
		Print()
	fmt.Println("Reversing")
	ll = ll.Reverse().
		Print()
	fmt.Println("Length", ll.Len())
*/
