package graphs

import (
	"reflect"
	"testing"
)

func TestSortedArrayToBST(t *testing.T) {
	// Input test case
	nums := []int{-10, -3, 0, 5, 9}

	// Expected output
	expected := &TreeNode{
		Val: 0,
		Left: &TreeNode{
			Val: -3,
			Left: &TreeNode{
				Val: -10,
			},
		},
		Right: &TreeNode{
			Val: 9,
			Left: &TreeNode{
				Val: 5,
			},
		},
	}

	// Convert sorted array to BST
	result := sortedArrayToBST(nums)

	// Compare the generated BST with the expected BST
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Generated BST does not match the expected BST")
	}
}
