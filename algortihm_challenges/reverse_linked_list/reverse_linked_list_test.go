package reverse_linked_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReverseList(t *testing.T) {
	// Create linked list
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}}

	reversed := reverseList(head)

	assert.Equal(t, 4, reversed.Val)
	assert.Equal(t, 3, reversed.Next.Val)
	assert.Equal(t, 2, reversed.Next.Next.Val)
	assert.Equal(t, 1, reversed.Next.Next.Next.Val)
}
