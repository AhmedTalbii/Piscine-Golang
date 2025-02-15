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
	l.Tail.Next = n
	l.Tail = n
}

func ListReverse(l *List) {
	if l == nil || l.Head == nil {
		return
	}
	arr := []interface{}{}
	node := l.Head
	for node != nil {
		arr = append(arr, node.Data)
		node = node.Next
	}
	nodeI := l.Head

	for i := len(arr) - 1; i >= 0; i-- {
		nodeI.Data = arr[i]
		nodeI = nodeI.Next
	}
}

func main() {
	link := &List{}

	ListPushBack(link, 1)
	ListPushBack(link, 2)
	ListPushBack(link, 3)
	ListPushBack(link, 4)
	ListPushBack(link, 5)
	ListPushBack(link, 6)

	ListReverse(link)

	it := link.Head

	for it != nil {
		fmt.Println(it.Data)
		it = it.Next
	}

	fmt.Println("Tail", link.Tail)
	fmt.Println("Head", link.Head)
}
