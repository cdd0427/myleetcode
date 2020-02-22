package main

import "fmt"

type queue []int

type MyStack struct {
	data queue
	size int
}

func (this *queue) push(x int) {
	(*this) = append((*this), x)
}

func (this *queue) pop() int {
	res := (*this)[0]
	(*this) = (*this)[1:]
	return res
}

/** Initialize your data structure here. */
func Constructor() MyStack {
	return MyStack{
		data: queue{},
		size: 0,
	}
}

/** Push element x onto stack. */
func (this *MyStack) Push(x int) {
	this.data.push(x)
	this.size++
	fmt.Println(this.data)
}

/** Removes the element on top of the stack and returns that element. */
func (this *MyStack) Pop() int {
	tmp := queue{}
	x := 0
	for {
		x = this.data.pop()
		if len(this.data) == 0 {
			break
		}
		tmp.push(x)
	}
	this.size--
	this.data = tmp
	fmt.Println(this.data)
	return x
}

/** Get the top element. */
func (this *MyStack) Top() int {
	for i := 0; i < this.size; i++ {
		if i == this.size-1 {
			return this.data[i]
		}
	}
	return 0
}

/** Returns whether the stack is empty. */
func (this *MyStack) Empty() bool {
	if this.size == 0 {
		return true
	}
	return false
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
