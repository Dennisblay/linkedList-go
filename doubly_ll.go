package main

import "fmt"

// Node represents a node in the doubly linked list
type Node struct {
	data interface{}
	prev *Node
	next *Node
}

// DoublyLinkedList represents a doubly linked list
type DoublyLinkedList struct {
	head *Node
	tail *Node
}

// InsertAtBeginning inserts a new node with the given data at the beginning of the doubly linked list
func (dll *DoublyLinkedList) InsertAtBeginning(data interface{}) {
	newNode := &Node{data: data}

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		newNode.next = dll.head
		dll.head.prev = newNode
		dll.head = newNode
	}
}

// InsertAtEnd inserts a new node with the given data at the end of the doubly linked list
func (dll *DoublyLinkedList) InsertAtEnd(data interface{}) {
	newNode := &Node{data: data}

	if dll.head == nil {
		dll.head = newNode
		dll.tail = newNode
	} else {
		newNode.prev = dll.tail
		dll.tail.next = newNode
		dll.tail = newNode
	}
}

// DisplayForward prints the elements of the doubly linked list in forward order
func (dll *DoublyLinkedList) DisplayForward() {
	current := dll.head
	for current != nil {
		fmt.Printf("%v ", current.data)
		current = current.next
	}
	fmt.Println()
}

// DisplayBackward prints the elements of the doubly linked list in backward order
func (dll *DoublyLinkedList) DisplayBackward() {
	current := dll.tail
	for current != nil {
		fmt.Printf("%v ", current.data)
		current = current.prev
	}
	fmt.Println()
}

func main() {
	dll := DoublyLinkedList{}
	// Inserting elements at the beginning of the list
	dll.InsertAtBeginning(3)
	dll.InsertAtBeginning(2)
	dll.InsertAtBeginning(1)

	// Inserting elements at the end of the list
	dll.InsertAtEnd(4)
	dll.InsertAtEnd(5)
	dll.InsertAtEnd(6)

	// Displaying the doubly linked list in forward and backward order
	fmt.Println("Forward order:")
	dll.DisplayForward()

	fmt.Println("Backward order:")
	dll.DisplayBackward()
}
