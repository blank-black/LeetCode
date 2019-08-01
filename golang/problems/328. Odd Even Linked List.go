package problems


 // Definition for singly-linked list.
 type ListNode struct {
	 Val  int
	 Next *ListNode
 }

func OddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	odd := head
	even := head.Next
	he := head.Next
	h := head.Next.Next
	for h != nil {
		odd.Next = h
		odd = h
		h = h.Next
		if h == nil {
			break
		}
		even.Next = h
		even = h
		h = h.Next
	}
	even.Next = nil
	odd.Next = he
	return head
}