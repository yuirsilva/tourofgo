package main

import (
	"fmt"
	"math/rand"
)

type Node[T interface{}] struct {
	data int
	next *Node[T]
}

type SinglyLinkedList[T interface{}] struct {
	els []Node[T]
}

func (l SinglyLinkedList[T]) getHead() Node[T] {
	return l.els[0]
}

func (l SinglyLinkedList[T]) getTail() Node[T] {
	return l.els[len(l.els)-1]
}

//	type SinglyLinkedList[T any] struct {
//		data T
//		next *SinglyLinkedList[T]
//	}

func main() {
	elements := make([]Node[interface{}], 4)

	list := SinglyLinkedList[interface{}]{
		els: elements,
	}

	for i := 0; i < len(list.els); i++ {
		list.els[i].data = rand.Intn(30)
		if i+1 == 4 {
			break
		}
		list.els[i].next = &(list.els[i+1])
	}
	fmt.Println(list.els)
	fmt.Println(*list.els[0].next)
}

// https://go.dev/tour/generics/2
