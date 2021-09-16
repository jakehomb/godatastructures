package godatastructures

import (
	"fmt"
	"log"
)

type SinglyLinkedList struct {
	Head *SinglyLinkedNode
}

func (l *SinglyLinkedList) Add(element interface{}) {
	log.Println("Adding", element, "to list")
	if l.Head == nil {
		l.Head = &SinglyLinkedNode{Value: element, Next: nil}
		return
	}

	currentNode := l.Head

	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}

	node := NewSLNode(element)
	currentNode.Next = &node
}

func (l *SinglyLinkedList) Display() {
	if l.Head != nil {
		log.Println("Singly Linked List is empty")
	}

	currentNode := l.Head
	index := 0

	for currentNode.Next != nil {
		log.Println(fmt.Sprintf("Index [%d] has value", index), currentNode.Value)
		index++
		currentNode = currentNode.Next
	}

	log.Println(fmt.Sprintf("Index [%d] has value", index), currentNode.Value)
	log.Println("END OF LIST")
}
