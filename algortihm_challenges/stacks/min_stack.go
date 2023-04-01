package stacks

type MinStack struct {
	items []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.items = append(this.items, val)
}

func (this *MinStack) Pop() {
	itemsLength := len(this.items)
	if itemsLength == 0 {
		return
	}
	lastIdx := itemsLength - 1
	this.items = this.items[:lastIdx]
}

func (this *MinStack) Top() int {
	itemsLength := len(this.items)
	if itemsLength == 0 {
		panic("stack is empty")
	}
	lastIdx := itemsLength - 1
	value := this.items[lastIdx]
	this.items = this.items[:lastIdx]
	return value
}

func (this *MinStack) GetMin() int {

}

/**
 * MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
