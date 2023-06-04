package main

//
//import (
//	"fmt"
//	"golang.org/x/exp/constraints"
//)
//
//func main_() {
//	ll := &LinkedList{}
//	xy1 := XYCoordinates[float64]{x: -1.497777, y: 6.9299923}
//	xy2 := XYCoordinates[float64]{x: -1.345234, y: 6.4939423}
//	xy3 := XYCoordinates[float64]{x: -1.398623, y: 6.2452524}
//	xy4 := XYCoordinates[float64]{x: -1.222222, y: 6.6789123}
//
//	ll.AddXY(xy1)
//	ll.AddXY(xy2)
//	ll.AddXY(xy3)
//	ll.AddXY(xy4)
//
//	//ll.InsertAtBeginning(xy4)
//	ll.InsertAtPosition(xy1, 2)
//	//ll.InsertAtPosition(xy2, 3)
//	//ll.Remove(xy1)
//
//	ll.DisplayXY()
//}
//
//type Number interface {
//	constraints.Integer | constraints.Float
//}
//type XYCoordinates[N Number] struct {
//	x N
//	y N
//}
//
//type node struct {
//	data XYCoordinates[float64]
//	next *node
//}
//
//type LinkedList struct {
//	head *node
//}
//
//func (ll *LinkedList) AddXY(data XYCoordinates[float64]) {
//	NewXYNode := &node{data: data, next: nil}
//
//	if ll.head == nil {
//		ll.head = NewXYNode
//	} else {
//		current := ll.head
//		for current.next != nil {
//			current = current.next
//		}
//		current.next = NewXYNode
//	}
//}
//
//func (ll *LinkedList) InsertAtBeginning(data XYCoordinates[float64]) {
//
//	if ll.head == nil {
//		return
//	}
//	newNode := &node{data: data}
//	newNode.next = ll.head
//	ll.head = newNode
//}
//
//func (ll *LinkedList) InsertAtPosition(data XYCoordinates[float64], position int) {
//	if position <= 0 {
//		ll.InsertAtBeginning(data)
//		return
//	}
//	newNode := &node{data: data}
//	current := ll.head
//	for count := 0; current.next != nil && count < position; count++ {
//		if count == position-1 {
//			newNode.next = current.next
//			current.next = newNode
//		}
//		current = current.next
//	}
//}
//
//func (ll *LinkedList) Remove(data XYCoordinates[float64]) {
//	//newNodeToRemove := &node{data: data}
//	if ll.head == nil {
//		return
//	}
//	if ll.head.data == data {
//		ll.head = ll.head.next
//	}
//	current := ll.head
//	for current.next != nil {
//		if current.next.data == data {
//			current.next = current.next.next
//			return
//		}
//		current = current.next
//	}
//
//}
//func (ll *LinkedList) SearchSuccess(data XYCoordinates[float64]) bool {
//	current := ll.head
//	for current.next != nil {
//		if current.next.data == data {
//			return true
//		}
//		current = current.next
//	}
//	return false
//}
//
//func (ll *LinkedList) DisplayXY() {
//	current := ll.head
//	for current != nil {
//		fmt.Printf("%v -> ", current.data)
//		current = current.next
//	}
//	fmt.Print("nil")
//}
//
////func (ll *LinkedList) String() string {
////	return fmt.Sprintf("<%v, %v>", ll.head.data.x, &ll.head.data.y)
////}
