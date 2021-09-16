package main

import (
	datastructures "github.com/jakehomb/DataStructures"
)

func main() {
	sllist := datastructures.SinglyLinkedList{}

	sllist.Add("Hello")
	sllist.Add(2)
	sllist.Add(3)

	sllist.Display()
}
