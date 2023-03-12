package randoms

func MergeKList(lists []*ListNode) *ListNode {

	n := len(lists)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return lists[0]
	}

	merge := func(left, right *ListNode) *ListNode {
		res := &ListNode{}
		cur := res
		for left != nil && right != nil {
			if left.Val < right.Val {
				cur.Next = left
				left = left.Next
			} else {
				cur.Next = right
				right = right.Next
			}
			cur = cur.Next
		}
		if left != nil {
			cur.Next = left
		}
		if right != nil {
			cur.Next = right
		}
		return res.Next
	}

	mid := n / 2
	left := MergeKList(lists[:mid])
	right := MergeKList(lists[mid:])

	return merge(left, right)
}
