package godatastructures

type DoublyLinkednode struct {
	Value    interface{}
	Previous *DoublyLinkednode
	Next     *DoublyLinkednode
}
