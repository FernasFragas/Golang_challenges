package fifo

import "testing"

func TestMyQueue(t *testing.T) {
	queue := Constructor()

	// Push elements into the queue
	queue = queue.Push(1)
	queue = queue.Push(2)
	queue = queue.Push(3)

	// Peek the front element
	result := queue.Peek()
	expected := 1
	if result != expected {
		t.Errorf("Peek: Expected %d, but got %d", expected, result)
	}

	// Pop the front element
	result, queue = queue.Pop()
	expected = 1
	if result != expected {
		t.Errorf("Pop: Expected %d, but got %d", expected, result)
	}

	// Check if the queue is empty
	resultBool := queue.Empty()
	expectedBool := false
	if result != expected {
		t.Errorf("Empty: Expected %t, but got %t", expectedBool, resultBool)
	}

	// Pop remaining elements
	result, queue = queue.Pop()
	expected = 2
	if result != expected {
		t.Errorf("Pop: Expected %d, but got %d", expected, result)
	}

	result, queue = queue.Pop()
	expected = 3
	if result != expected {
		t.Errorf("Pop: Expected %d, but got %d", expected, result)
	}

	// Check if the queue is empty
	resultBool = queue.Empty()
	expectedBool = true
	if result != expected {
		t.Errorf("Empty: Expected %t, but got %t", expectedBool, resultBool)
	}
}
