package fifo

type MyQueue struct {
	tail *Node
	head *Node
}

type Node struct {
	value int
	next  *Node
}

func Constructor() MyQueue {
	return MyQueue{nil, nil}
}

func (this *MyQueue) Push(x int) {
	if this.head == nil {
		node := Node{x, this.tail}
		this.head = &node
		return
	}
	if this.head.next == nil {
		node := Node{x, nil}
		this.head.next = &node
		this.tail = &node
		return
	}
	node := Node{x, nil}
	this.tail.next = &node
	this.tail = &node
}

func (this *MyQueue) Pop() int {
	value := this.head.value
	this.head = this.head.next
	return value
}

func (this *MyQueue) Peek() int {
	value := this.head.value
	return value
}

func (this *MyQueue) Empty() bool {
	return this.head == nil
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
