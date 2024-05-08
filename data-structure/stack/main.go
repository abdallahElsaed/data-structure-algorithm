package main

import (
	"errors"
	"fmt"
)

type Stack struct {
	data []int
	top  int
}

func createStack(capacity int) *Stack {
	return &Stack{
		data: make([]int , 0 , capacity),
		top:  -1,
	}
}

/***
		- increase top 1
		- append the element in data array at top position
*/
func (s *Stack) push(value int) {
	s.top++
	s.data= append(s.data, value)
}

/*
*
check if the top == -1
*/
func (s *Stack) isEmpty() bool {
	return s.top == -1
}

/*
*
- Check if the data is empty
- If not,
  - take the top element value
  - remove the top element
  - decrease the top by 1
*/
func (s *Stack) pop() (int, error){
	if s.isEmpty() {
		// fmt.Println("You can not pop, your stack is empty")
		return 0 , errors.New("You can not pop")
		// errors.New("your stack is empty")
	}
	element := s.data[s.top]
	s.data = s.data[:s.top]
	s.top--
	return element , nil
}
func (s *Stack) getTop() {
	fmt.Printf("Top element is %d" ,s.data[s.top])
}

func (s *Stack) print() {
	for _, value := range s.data {
		fmt.Printf(" %d\n", value)
	}
}
func main() {
	// stack := createStack(3)
	// stack.push(1)
	// stack.push(2)
	// stack.push(3)
	// stack.push(4)

	// stack.pop()
	// stack.pop()
	// stack.pop()
	// stack.pop()
	// stack.getTop()
	// stack.print()

	llStack := LinkedList{} 
	llStack.pushLL(1)
	llStack.pushLL(2)
	llStack.pushLL(3)
	llStack.printLL()

}
