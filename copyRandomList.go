package demo

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

// 两次遍历
func CopyRandomList(head *Node) *Node {
	maps := make(map[*Node]*Node)
	cur := head
	for cur != nil {
		temp := &Node{
			Val: cur.Val,
		}
		maps[cur] = temp
		cur = cur.Next
	}
	cur = head
	for cur != nil {
		maps[cur].Next = maps[cur.Next]
		maps[cur].Random = maps[cur.Random]
		cur = cur.Next
	}
	return maps[head]
}
