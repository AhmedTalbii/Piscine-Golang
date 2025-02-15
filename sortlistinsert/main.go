package main

import (
	"fmt"
)

type NodeI struct {
	Data int
	Next *NodeI
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

func SortListInsert(l *NodeI, data_ref int) *NodeI {
	listPushBack(l, data_ref)
	ListSort(l)
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

	link = listPushBack(link, 1)
	link = listPushBack(link, 4)
	link = listPushBack(link, 9)

	PrintList(link)

	link = SortListInsert(link, -2)
	link = SortListInsert(link, 2)
	PrintList(link)
}
