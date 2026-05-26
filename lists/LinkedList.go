package lists

import (
	"fmt"
)

type Node[T comparable] struct {
	Val  T
	Next *Node[T]
	Prev *Node[T]
}
type LinkedList[T comparable] struct {
	Head  *Node[T]
	Tail  *Node[T]
	Cur   *Node[T]
	Lenth int
}

func (l *LinkedList[T]) PushFront(v T) {
	newNode := &Node[T]{
		Val:  v,
		Next: nil,
		Prev: nil,
	}
	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
	} else {
		newNode.Next = l.Head
		l.Head.Prev = newNode
		l.Head = newNode
	}
	l.Lenth++
}
func (l *LinkedList[T]) PushBack(v T) {
	newNode := &Node[T]{
		Val:  v,
		Next: nil,
		Prev: nil,
	}
	if l.Tail == nil {
		l.Tail = newNode
		l.Head = newNode
	} else {
		newNode.Prev = l.Tail
		l.Tail.Next = newNode
		l.Tail = newNode
	}
	l.Lenth++
}
func (l *LinkedList[T]) InsertBefore(v T, targetVal T) {
	var targetNode *Node[T]

	current := l.Head
	for current != nil {
		if current.Val != targetVal {
			current = current.Next
			continue
		}
		targetNode = current
		break
	}
	if targetNode == nil {
		fmt.Print("There is no such node")
		return
	}
	var newNode = &Node[T]{Val: v}

	if targetNode.Prev == nil {
		l.Head = newNode
	} else {
		targetNode.Prev.Next = newNode
		newNode.Prev = targetNode.Prev
	}
	newNode.Next = targetNode
	targetNode.Prev = newNode
	l.Lenth++
}
func (l *LinkedList[T]) InsertAfter(v T, targetVal T) {
	var targetNode *Node[T]

	current := l.Head
	for current != nil {
		if current.Val != targetVal {
			current = current.Next
			continue
		}
		targetNode = current
		break
	}
	if targetNode == nil {
		fmt.Print("There is no such node")
		return
	}
	var newNode = &Node[T]{Val: v}
	if targetNode.Next == nil {
		l.Tail = newNode
	} else {
		newNode.Next = targetNode.Next
		targetNode.Next.Prev = newNode
	}

	newNode.Prev = targetNode
	targetNode.Next = newNode

	l.Lenth++
}

func (l *LinkedList[T]) Remove(val T) {
	cur_node := l.Head
	var targetNode *Node[T]

	for cur_node != nil {
		if cur_node.Val != val {
			cur_node = cur_node.Next
			continue
		} else {
			targetNode = cur_node
			break
		}
	}
	if targetNode == nil {
		fmt.Print("There is no such node")
		return
	}
	if targetNode.Prev != nil {
		targetNode.Prev.Next = targetNode.Next
	} else {
		l.Head = targetNode.Next
	}
	if targetNode.Next != nil {
		targetNode.Next.Prev = targetNode.Prev
	} else {
		l.Tail = targetNode.Prev
	}
	l.Lenth--
}
func (l *LinkedList[T]) Front() T {
	if l.Head == nil {
		var zero T
		return zero
	}
	return l.Head.Val

}
func (l *LinkedList[T]) Back() T {
	if l.Tail == nil {
		var zero T
		return zero
	}
	return l.Tail.Val
}
func (l *LinkedList[T]) Len() int {
	return l.Lenth
}
