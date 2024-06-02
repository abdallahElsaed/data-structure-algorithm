package main

import "fmt"

type queue struct {
	data []int
}

func (q *queue) enqueue(item int) {
	q.data  = append(q.data, item)
}
func (q *queue) dequeue() {
	q.data  = q.data[:len(q.data)-1]
}

func (q *queue) peek()  {
	fmt.Printf("the peek is: %d \n", q.data[0])
}

func (q *queue) isEmpty() bool {
	if len(q.data) == 0 {
		return true
	}
	return false
}

func (q *queue) print()  {
	if q.isEmpty()  {
		fmt.Println("the queue is empty")
	}else{
		for  _ , item :=range q.data{
			fmt.Println(item)
		}
	}
}

func main() {
	// queue := &queue{
	// 	data: make([]int, 0),
	// }
//! implement queue with array.
	// queue.enqueue(1)
	// queue.enqueue(3)
	// queue.enqueue(5)
	// queue.dequeue()
	// queue.dequeue()
	// queue.dequeue()
	// queue.peek()
	// queue.isEmpty()
	// queue.print()

//! implement queue with linked list.
llQueue := &QueueLinkedList{
		front: nil,
		end: nil,
	}

	llQueue.enqueueLl(1)
	llQueue.enqueueLl(2)
	llQueue.enqueueLl(3)
	// llQueue.dequeueLl()

	llQueue.frontLl()
	// llQueue.printLl()


}
/* Note in go 
	* You can make slice by make() 	  
	* You can push new element to array by append()
	* You must reassign to queue data slice when make any update like this 
		- q.data  = append(q.data, item)
*/