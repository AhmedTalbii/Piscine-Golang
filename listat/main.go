package main

import (
	"fmt"
)

type NodeL struct {
	Data interface{}
	Next *NodeL
}

type List struct {
	Head *NodeL
	Tail *NodeL
}

func ListPushBack(l *List, data interface{}) {
	n := &NodeL{Data: data}
	if l.Tail == nil {
		l.Head = n
		l.Tail = n
	}
	// create new node
	l.Tail.Next = n
	// add val to the new node
	l.Tail = n
}

func ListAt(l *NodeL, pos int) *NodeL {
	res := 0
	for l != nil {
		// fmt.Println(l)
		if res == pos {
			// fmt.Println(l)
			return l
		}
		l = l.Next
		res++
	}
	return nil
}

func main() {
	link := &List{}

	ListPushBack(link, "hello")
	ListPushBack(link, "how are")
	ListPushBack(link, "you")
	ListPushBack(link, 1)

	fmt.Println(ListAt(link.Head, 3).Data)
	fmt.Println(ListAt(link.Head, 1).Data)
	fmt.Println(ListAt(link.Head, 7))
}
