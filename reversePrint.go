package demo

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reversePrint(head *ListNode) []int {
	if head == nil {
		return nil
	}

	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}

	for i, j := 0, len(result)-1; i < j; {
		result[i], result[j] = result[j], result[i]
		i++
		j--
	}
	return result
}
