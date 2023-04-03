package linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

// reverseList iterates through the original list and updates the next pointers of each node to point to the previous node
func reverseList(head *ListNode) *ListNode {
	// to keep track of the previously processed node
	var previousNode *ListNode
	for head != nil {
		tempNode := head.Next    //to avoid losing the next node of the current node
		head.Next = previousNode // reverses the direction of the arrow
		previousNode = head
		head = tempNode // move iteration to the next node
	}
	return previousNode
}
