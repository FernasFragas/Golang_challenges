package linked_list

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		if list2 == nil {
			return nil
		}
		return list2
	}
	if list2 == nil {
		if list1 == nil {
			return nil
		}
		return list1
	}
	//verifies if one of the list is empty if so return the other one
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	}
	if list1 == nil && list2 == nil {
		return list1
	}

	head := &ListNode{}
	curr := head
	for list1 != nil && list2 != nil {
		// verifies which value is smaller and adds it to the list
		if list1.Val >= list2.Val { // case list1 is bigger
			curr.Next = list2
			list2 = list2.Next
			// verifies if the curr.next value is not nil, if it is, the next element will be the element of the other list
			if curr.Next != nil {
				curr = curr.Next
			} else {
				curr.Next = list1
			}
		} else { // case list2 is bigger
			curr.Next = list1
			list1 = list1.Next
			if curr.Next != nil {
				curr = curr.Next
			} else {
				curr.Next = list2
			}
		}
	}
	// verifies if there is some element that was not iterated if so, it adds it to the list
	if list1 != nil {
		curr.Next = list1
	} else if list2 != nil {
		curr.Next = list2
	}
	return head.Next
}
