package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}

type DoublyLinkedList struct {
	head *Node
	tail *Node
}

// insert in double linked list at beginning
func (dll *DoublyLinkedList) insertAtBeginning(data int) {
	newNode := &Node{data: data}
	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
		return
	}
	newNode.next = dll.head
	dll.head.prev = newNode
	dll.head = newNode
}
func (dll *DoublyLinkedList) insertAtEnd(data int) {
	newNode := &Node{data: data}
	current := dll.head
	for current.next != nil {
		current = current.next
	}
	newNode.prev = current
	current.next = newNode
	dll.tail = current
}

func (dll *DoublyLinkedList) insertAtMiddle(data int, key int) {
	newNode := &Node{data: data}
	current := dll.head
	if current == nil {
		fmt.Println("Your linked list is empty")
		return
	}

	for current.data != key {
		if current.next == nil {
			fmt.Println("Key not found in the linked list")
			return
		}
		current = current.next
	}
	newNode.prev = current
	newNode.next = current.next
	if current.next != nil {
		current.next.prev = newNode
	}
	current.next = newNode
}

func (dll *DoublyLinkedList) PrintData() {
	current := dll.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Printf("null")
}

func main() {
	dll := DoublyLinkedList{}

	// Insert some nodes into the doubly linked list
	dll.insertAtBeginning(1)
	dll.insertAtBeginning(2)
	dll.insertAtBeginning(3)
	dll.insertAtEnd(10)
	dll.insertAtMiddle(15, 2)

	dll.PrintData()
}
