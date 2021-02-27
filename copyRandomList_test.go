package demo

import (
	"fmt"
	"testing"
)

func TestCopyRandomList(t *testing.T) {
	h := &Node{Val: 7, Next: nil, Random: nil}
	h.Next = &Node{Val: 13, Next: nil, Random: h}
	h.Next.Next = &Node{Val: 11, Next: nil, Random: nil}
	h.Next.Next.Next = &Node{Val: 10, Next: nil, Random: nil}
	h.Next.Next.Next.Next = &Node{Val: 1, Next: nil, Random: h}
	h.Next.Next.Random = h.Next.Next.Next.Next
	h.Next.Next.Next.Random = h.Next.Next
	list := CopyRandomList(h)
	for list != nil {
		fmt.Println("[", list.Val, " ", list.Random, "]")
		list = list.Next
	}
}
