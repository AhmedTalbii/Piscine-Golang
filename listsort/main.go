package main

import (
	"fmt"
)

type NodeI struct {
	Data int
	Next *NodeI
}

func ListPushBack(l *NodeI, data int) {
	n := &NodeI{Data: data}
	if l == nil {
		l = n
	}
	current := l
	for current != nil {
		current = current.Next
	}
	current = n
}

func ListSort(l *NodeI) *NodeI {
	current := l
	for current != nil && current.Next != nil {
		if current.Data > current.Next.Data {
			current.Data, current.Next.Data = current.Next.Data, current.Data
			current = l
		} else {
			current = current.Next
		}
	}
	return l
}

func PrintList(l *NodeI) {
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func listPushBack(l *NodeI, data int) *NodeI {
	n := &NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func main() {
	var link *NodeI

	link = listPushBack(link, 5)
	link = listPushBack(link, 4)
	link = listPushBack(link, 3)
	link = listPushBack(link, 2)
	link = listPushBack(link, 1)

	PrintList(ListSort(link))
}
