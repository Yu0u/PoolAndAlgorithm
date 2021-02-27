package demo

import (
	"fmt"
	"testing"
)

func TestReversePrint(t *testing.T) {
	l := &ListNode{Val: 1, Next: nil}
	l.Next = &ListNode{Val: 2, Next: nil}
	l.Next.Next = &ListNode{Val: 3, Next: nil}
	l.Next.Next.Next = &ListNode{Val: 4, Next: nil}
	l.Next.Next.Next.Next = &ListNode{Val: 5, Next: nil}

	ints := reversePrint(l)
	fmt.Println(ints)
}
