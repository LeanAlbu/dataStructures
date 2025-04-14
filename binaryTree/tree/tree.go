package tree

import "fmt"

type Node struct {
	value  int
	left   *Node
	right  *Node
	Parent *Node
}

type Tree struct {
	Root *Node
	Qty  int
}

func TreeCreate() *Tree {
	return &Tree{Qty: 0, Root: nil}
}

func NodeCreate(data int) *Node {
	return &Node{value: data, right: nil, left: nil, Parent: nil}
}

func (tree *Tree) Insert(value int) {
	data := NodeCreate(value)
	if tree.Root == nil {
		tree.Root = data
		tree.Qty++
		return
	} else {
		aux := tree.Root
		for {
			if data.value == aux.value {
				return
			}
			if data.value > aux.value {
				if aux.right == nil {
					aux.right = data
					data.Parent = aux
					tree.Qty++
					return
				} else {
					aux = aux.right
				}
			}
			if data.value < aux.value {
				if aux.left == nil {
					aux.left = data
					data.Parent = aux
					tree.Qty++
					return
				} else {
					aux = aux.left
				}
			}
		}
	}
}

func CreateFromArray(data []int) {
	newTree := TreeCreate()
	for i := range data {
		newTree.Insert(i)
	}
}

func CreateArray(n *Node) []int {
	var myArr = []int{}
	if n != nil {
		myArr = append(myArr, n.value)
		myArr = append(myArr, CreateArray(n.left)...)
		myArr = append(myArr, CreateArray(n.right)...)
	}
	return myArr
}

func Print(root *Node) {
	if root == nil {
		return
	}
	fmt.Println(root.value)
	Print(root.left)
	Print(root.right)
}

func Search(node *Node,value int)*Node{
	if node == nil{
		return nil
	}
	if node.value == value{
		return node
	}
	if left := Search(node.left, value); left != nil{
		return left
	}
	return Search(node.right, value)
}

