package linked_list

import "testing"

// Test case for merging two sorted linked lists
func TestMergeTwoLists(t *testing.T) {
	// Create two sorted linked lists
	list1 := &ListNode{1, &ListNode{2, &ListNode{4, nil}}}
	list2 := &ListNode{1, &ListNode{3, &ListNode{4, nil}}}

	// Merge the lists
	result := mergeTwoLists(list1, list2)

	// Expected merged list: 1 -> 1 -> 2 -> 3 -> 4 -> 4
	expected := []int{1, 1, 2, 3, 4, 4}
	for i, val := range expected {
		if result.Val != val {
			t.Errorf("Expected %d, but got %d at index %d", val, result.Val, i)
		}
		result = result.Next
	}

	// Create two sorted linked lists
	list1 = &ListNode{1, nil}
	list2 = &ListNode{2, nil}

	// Merge the lists
	result = mergeTwoLists(list1, list2)

	// Expected merged list: 1 -> 1 -> 2 -> 3 -> 4 -> 4
	expected = []int{1, 2}
	for i, val := range expected {
		if result.Val != val {
			t.Errorf("Expected %d, but got %d at index %d", val, result.Val, i)
		}
		result = result.Next
	}
}
