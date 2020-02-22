package main

type stack []int

type MyQueue struct {
	data stack
	size int
}

func (this *stack) push(x int) {
	*this = append(*this, x)
}

func (this *stack) pop() int {
	res := (*this)[len(*this)-1]
	*this = (*this)[:len(*this)-1]
	return res
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{
		data: stack{},
		size: 0,
	}
}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.data.push(x)
}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	opt := stack{}
	x := 0
	for {
		x = this.data.pop()
		if len(this.data) == 0 {
			break
		}
		opt.push(x)
	}
	for len(opt) != 0 {
		this.data.push(this.data.pop())
	}
	this.size--
	return x
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	for i := this.size - 1; i >= 0; i-- {
		if i == 0 {
			return this.data[0]
		}
	}
	return 0
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	if this.size == 0 {
		return true
	}
	return false
}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
