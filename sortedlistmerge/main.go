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

func ListMerge(l1 *NodeI, l2 *NodeI) {
	node := l2
	for node != nil {
		listPushBack(l1, node.Data)
		node = node.Next
	}
}

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	ListMerge(n1, n2)
	ListSort(n1)
	return n1
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
	var link2 *NodeI

	link = listPushBack(link, 3)
	link = listPushBack(link, 5)
	link = listPushBack(link, 7)

	link2 = listPushBack(link2, -2)
	link2 = listPushBack(link2, 9)

	PrintList(SortedListMerge(link2, link))
}

// -2 -> 3 -> 5 -> 7 -> 9 -> <nil>
