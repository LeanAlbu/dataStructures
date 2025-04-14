package main

import (
	"binaryTree/tree"
	"fmt"
)

func main() {
	newTree := tree.TreeCreate()
	newTree.Insert(10)
	newTree.Insert(8)
	newTree.Insert(4)
	newTree.Insert(20)
	tree.Print(newTree.Root)

	myArr := tree.CreateArray(newTree.Root)

	fmt.Println(myArr)

	find := tree.Search(newTree.Root, 8)

	fmt.Println(find)
	fmt.Println(find.Parent)
}
