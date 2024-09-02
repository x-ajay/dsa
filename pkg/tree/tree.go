package tree

type Tree struct {
	root *Node
}

type Node struct {
	data  int
	left  *Node
	right *Node
}

func New() *Tree {
	return &Tree{}
}

func 