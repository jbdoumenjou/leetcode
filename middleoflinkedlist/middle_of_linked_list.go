package middleoflinkedlist

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// Given the head of a singly linked list,
// return the middle node of the linked list.
//
// If there are two middle nodes, return the second middle node.
//
// Time complexity: O(n)
// Space complexity: O(n)
func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var nodes []*ListNode

	for head != nil {
		nodes = append(nodes, head)
		head = head.Next
	}

	return nodes[len(nodes)/2]
}

// Time complexity: O(n)
// Space complexity: O(1)
func middleNode_ptr(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	// 1
	// 1 2 3
	//   ^
	// 1 2 3 4
	//     ^
	// 1 2 3 4 5
	//     ^
	// 1 2 3 4 5 6
	//       ^
	middle, end := head, head
	for end != nil && end.Next != nil {
		middle = middle.Next
		end = end.Next.Next
	}

	return middle
}
