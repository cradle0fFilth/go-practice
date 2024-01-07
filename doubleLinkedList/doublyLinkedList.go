package main

import "fmt"

type Node struct {
	Data int
	Prev *Node
	Next *Node
}

type DoublyLinkedList struct {
	Head *Node
	Tail *Node
}

func (list *DoublyLinkedList) Append(data int) {
	newNode := &Node{Data: data}
	if list.Head == nil {
		list.Head = newNode
		list.Tail = newNode
		return
	}
	list.Tail.Next = newNode
	newNode.Prev = list.Tail
	list.Tail = newNode
}

func (list *DoublyLinkedList) Display() {
	current := list.Head
	for current != nil {
		fmt.Print(current.Data, "")
		current = current.Next
	}
	fmt.Println()
}
func main() {
	dll := DoublyLinkedList{}
	dll.Append(1)
	dll.Append(2)
	dll.Append(3)

	dll.Display() // 输出: 1 2 3
}
