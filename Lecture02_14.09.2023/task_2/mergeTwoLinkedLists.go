package task_2

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/merge-two-sorted-lists/submissions/1055536253/
// Time Complexity: O(m+n)
// Space Complexity: O(1)
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	ret := &ListNode{}
	proc := ret

	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			proc.Next = list1
			list1 = list1.Next
		} else {
			proc.Next = list2
			list2 = list2.Next
		}
		proc = proc.Next
	}

	if list1 != nil {
		proc.Next = list1
	}
	if list2 != nil {
		proc.Next = list2
	}

	return ret.Next
}
