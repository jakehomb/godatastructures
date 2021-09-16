package godatastructures

type SinglyLinkedNode struct {
	Value interface{}
	Next  *SinglyLinkedNode
}

func NewSLNode(value interface{}) SinglyLinkedNode {
	return SinglyLinkedNode{Value: value, Next: nil}
}
