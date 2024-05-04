package main

import "fmt"

/*
	* Node object
		- data
		- next *node
	* linked list object
		- top *node
*/

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	top *Node
}

/*
	* create new node
	* if empty
		- make top == new node
	* if not
		- loop over list
		- current node next = new node
*/
func (ll *LinkedList) pushLL(data int) {
	newNode := &Node{data: data}

	if ll.top == nil {
		ll.top = newNode
	}

	current := ll.top
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (ll *LinkedList) printLL() {
	current := ll.top
	fmt.Print("Linked List: ")
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Printf("null")
}
