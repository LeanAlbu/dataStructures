package main

import "fmt"

type Node struct {
	data  rune
	left  *Node
	right *Node
}

type Tree struct {
	qty  int
	root *Node
}

func main() {
	tree := TreeCreate()
	tree.root = NodeCreate('R')
	tree.root.left = NodeCreate('A')
	tree.root.right = NodeCreate('B')
	tree.root.left.left = NodeCreate('C')
	tree.root.left.right = NodeCreate('D')
	tree.root.right.left = NodeCreate('E')
	tree.root.right.right = NodeCreate('F')
	tree.root.right.right.left = NodeCreate('G')

	TreePrintDSA(tree.root)
}

func NodeCreate(data rune) *Node {
	return &Node{data: data, left: nil, right: nil}
}

func TreeCreate() *Tree {
	return &Tree{qty: 0, root: nil}
}

func TreePrintDSA(root *Node) {
	if root == nil {
		return
	}
	fmt.Printf("%c ", root.data)
	TreePrintDSA(root.left)
	TreePrintDSA(root.right)
}
