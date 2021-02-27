package demo

import (
	"fmt"
	"testing"
)

func TestReverseList(t *testing.T) {

	l := &ListNode{Val: 1, Next: nil}
	l.Next = &ListNode{Val: 2, Next: nil}
	l.Next.Next = &ListNode{Val: 3, Next: nil}
	l.Next.Next.Next = &ListNode{Val: 4, Next: nil}
	l.Next.Next.Next.Next = &ListNode{Val: 5, Next: nil}

	list := reverseList(l)
	for list != nil {
		fmt.Println("Val:", list.Val)
		list = list.Next
	}
}
