package n0002

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(vals ...int) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	for _, v := range vals {
		curr.Next = &ListNode{Val: v}
		curr = curr.Next
	}
	return dummy.Next
}
