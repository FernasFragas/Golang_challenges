package stacks

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMinStack(t *testing.T) {
	// Create a new MinStack
	minStack := Constructor()

	// Test push method
	minStack.Push(1)
	minStack.Push(2)
	minStack.Push(3)

	// Test top method
	assert.Equal(t, 3, minStack.Top())

	// Test getMin method
	assert.Equal(t, 1, minStack.GetMin())

	// Test pop method
	minStack.Pop()
	assert.Equal(t, 2, minStack.Top())

	// Test getMin method after pop
	assert.Equal(t, 1, minStack.GetMin())
}
