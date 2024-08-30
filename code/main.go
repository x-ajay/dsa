package main

import (
	"fmt"

	"github.com/x-ajay/dsa/pkg/linkedlist"
)

func main() {

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

}
