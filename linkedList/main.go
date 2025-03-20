package main

func main() {

	myList := llcreate()

	items := []int{1, 2, 3, 4, 5}

	for _, item := range items {
		myList.llappend(item)
	}

	myList.llprint()

	myList.llprepend(0)

	myList.llprint()

	myList.lldelete(0)

	myList.llprint()
}
