package demo

//type ListNode struct {
//	Val int
//	Next *ListNode
//}
// 快慢指针

func getKthFromEnd(head *ListNode, k int) *ListNode {
	slow := head
	fast := head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast != nil {
		slow, fast = slow.Next, fast.Next
	}
	return slow
}
