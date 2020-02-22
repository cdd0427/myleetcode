package main

type stack []int

type CQueue struct {
	data stack
	tmp  stack
}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) push(x int) {
	*s = append(*s, x)
}

func (s *stack) pop() int {
	x := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return x
}

func Constructor() CQueue {
	return CQueue{
		data: stack{},
		tmp:  stack{},
	}
}

func (this *CQueue) AppendTail(value int) {
	this.data.push(value)
}

func (this *CQueue) DeleteHead() int {
	if this.data.isEmpty() {
		return -1
	}
	for len(this.data) != 1 {
		this.tmp.push(this.data.pop())
	}
	x := this.data.pop()
	for !this.tmp.isEmpty() {
		this.data.push(this.tmp.pop())
	}
	return x
}
