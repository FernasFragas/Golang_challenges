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

	if list1.Next == nil {
		return list2
	} else if list2.Next == nil {
		return list1.Next
	}
	if list1.Next == nil && list2.Next == nil {
		return list1
	}

	var head = list1
	for list1.Next != nil || list2.Next != nil {
		if list1.Val > list2.Val {
			temp := list1.Next
			list1.Next = list2.Next
			list2.Next = temp
			if head.Val > list2.Val {
				head = list2
			}
		} else {
			list1 = list1.Next
			continue
		}
		list2 = list2.Next
	}
	return head
}
