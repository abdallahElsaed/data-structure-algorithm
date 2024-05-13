package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type QueueLinkedList struct {
	end   *Node
	front *Node
}

func (ll *QueueLinkedList) enqueueLl(item int) {
	node := &Node{
		data: item,
	}
	if ll.isEmptyLl() {
		ll.front = node
		ll.end = node
		return
	}
	ll.end.next = node
	ll.end = node
}
func (ll *QueueLinkedList) dequeueLl() {
	if ll.isEmptyLl() {
		fmt.Println("Your queue is empty, You can't remove")
	}
	if ll.front == ll.end {
		ll.end = nil
		ll.front = nil
		return
	}
	ll.front = ll.front.next

}
func (ll *QueueLinkedList) isEmptyLl() bool {
	if ll.front == nil {
		return true
	}
	return false
}
func (ll *QueueLinkedList) frontLl() {
	fmt.Printf("The front is: %d", ll.front.data )
}
func (ll *QueueLinkedList) printLl() {
	if ll.isEmptyLl() {
		fmt.Println("Your queue is empty")
	}
	current := ll.front 
	for current != ll.end.next { // or != nil
		fmt.Printf("%d <- " ,current.data)
		current = current.next
	}
	fmt.Printf("null")

}