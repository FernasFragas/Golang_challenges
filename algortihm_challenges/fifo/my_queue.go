package fifo

type MyQueue struct {
	value *int
	next  *MyQueue
	head  *MyQueue
}

func Constructor() MyQueue {
	return MyQueue{nil, nil, nil}
}

func (this *MyQueue) Push(x int) {

}

func (this *MyQueue) Pop() int {
	value := *this.head.value
	this.head = this.head.next
	return value
}

func (this *MyQueue) Peek() int {
	value := *this.head.value
	return value
}

func (this *MyQueue) Empty() bool {
	return this.head.value == nil
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
