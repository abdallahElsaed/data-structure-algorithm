package main

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}
type LinkedList struct {
	head *Node
}

//! create newNode and insert at end of it.
func (ll *LinkedList) insertAtEnd(data int) {
	// create new node and store his address in newNode variable
	newNode := &Node{data: data}

	// check if ll.head == nil
	if ll.head == nil {
		ll.head = newNode
		return
	}
	// create new next to point in current node to make check on it.
	current := ll.head
	// loop over the list to reach at end the list.
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

// ! insert at beginning of the list
func (ll *LinkedList) insertAtBeginning(data int) {
	newNode := &Node{data: data}
	newNode.next = ll.head
	ll.head = newNode
}

// ! insert at Middle of the list
func (ll *LinkedList) insertAtMiddle(data, key int) {
	newNode := &Node{data: data}
	current := ll.head
	for current != nil && current.data != key {
		current = current.next
	}
	if current == nil {
		fmt.Println("Key not found in the linked list")
	}
	newNode.next = current.next
	current.next = newNode
}
/** Delete node  
	* 1- delete first node 
	* 2- delete last node 
	* 3- delete nth node 
 */  
// ?  1- delete first node 
/**
	** linked list head sign to second node
	(in the latter you need way to delete from memory)
*/
func (ll *LinkedList) deleteAtBeginning(){
	headNode := ll.head 
	ll.head = headNode.next
}
// ?  2- delete last node 
/**
	** iterate over list which the second node next from end = null
	** make the second node next = null
	(in the latter you need way to delete from memory)
*/
func (ll *LinkedList) deleteAtEnd(){
	current := ll.head
	for  current.next.next != nil{
		current = current.next
	}
	current.next = nil
}

// ?  3- delete nth position 
/**
	** iterate over list which the node next data == key
	** make the second  next = null (delete next node)
	(in the latter you need way to delete from memory)
*/
func (ll *LinkedList) deleteAtMiddle( key int) {
	
	current := ll.head
	for current != nil && current.next.data != key {
		current = current.next
	}
	if current == nil {
		fmt.Println("Key not found in the linked list")
	}
	current.next = current.next.next
}

//! reverse node in list
/*
	* make 3 variable 
		- next -> to keep the address of the next node 
		- current -> to iterate over the list 
		- previous -> to save the address of the previous node and make current.next to equal the previous and make the reverse 
	** explain the steps:
	1- initialize three variables [current, next, previous]
	2- loop over the linked list when the current != null
	3- make the next var assigned to next node of the current
	5- reverse the next of the current to assign to the previous node
	6- the previous variable moves to the current node
	7- the current node moves to the next node
	8- after the loop finished the list is finished At this point make the head assign to the previous variable(node)
*/
func (ll *LinkedList) reverse(){
	var previous , next *Node 
	current := ll.head

	for current != nil {
		next = current.next 
		current.next = previous
		previous = current 
		current = next
	}

	ll.head = previous
}

// ! print linked list data
func (ll *LinkedList) PrintData() {
	current := ll.head
	for current != nil {
		fmt.Printf("%d -> ", current.data)
		current = current.next
	}
	fmt.Printf("null")

}

func main() {
	ll := LinkedList{}

	ll.insertAtEnd(2)
	ll.insertAtBeginning(4)
	ll.insertAtEnd(1)
	ll.insertAtEnd(3)
	ll.insertAtMiddle(10 , 1)
	// ll.deleteAtBeginning()
	// ll.deleteAtEnd()
	// ll.deleteAtMiddle(1)
	ll.reverse()

	fmt.Printf("Linked-list : ") // Linked-list : 4 -> 2 -> 1 -> 10 -> 3 -> null

	ll.PrintData()

}
