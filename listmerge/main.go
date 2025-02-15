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
		l.Tail = n
		l.Head = n
		return
	}
	l.Tail.Next = n
	l.Tail = n
}

func ListMerge(l1 *List, l2 *List) {
	node2 := l2.Head

	for node2 != nil {
		ListPushBack(l1, node2.Data)
		node2 = node2.Next
	}
}

func PrintList(l *List) {
	it := l.Head
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func main() {
	link := &List{}
	link2 := &List{}

	ListPushBack(link, "a")
	ListPushBack(link, "b")
	ListPushBack(link, "c")
	ListPushBack(link, "d")
	fmt.Println("-----first List------")
	PrintList(link)

	ListPushBack(link2, "e")
	ListPushBack(link2, "f")
	ListPushBack(link2, "g")
	ListPushBack(link2, "h")
	fmt.Println("-----second List------")
	PrintList(link2)

	fmt.Println("-----Merged List-----")
	ListMerge(link, link2)
	PrintList(link)
}
