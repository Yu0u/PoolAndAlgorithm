package demo

import (
	"fmt"
	"testing"
)

func TestMergeKLists(t *testing.T) {
	node := &ListNode{Val: 1, Next: nil}
	node.Next = &ListNode{Val: 4, Next: nil}
	node.Next.Next = &ListNode{Val: 5, Next: nil}
	var list []*ListNode
	list = append(list, node)
	node1 := &ListNode{Val: 1, Next: nil}
	node1.Next = &ListNode{Val: 3, Next: nil}
	node1.Next.Next = &ListNode{Val: 4, Next: nil}
	list = append(list, node1)
	node2 := &ListNode{Val: 2, Next: nil}
	node2.Next = &ListNode{Val: 6, Next: nil}
	list = append(list, node2)
	lists := mergeKLists(list)
	cur := lists
	for cur != nil {
		fmt.Println("val:", cur.Val)
		cur = cur.Next
	}
}
