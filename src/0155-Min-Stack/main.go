package main

type MinStack struct {
	data []int
	min  []int
	size int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		data: []int{},
		min:  []int{},
		size: 0,
	}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	if this.size == 0 {
		this.min = append(this.min, x)
	} else if x < this.min[this.size-1] {
		this.min = append(this.min, x)
	} else {
		this.min = append(this.min, this.min[this.size-1])
	}
	this.size++
}

func (this *MinStack) Pop() {
	this.data = this.data[:this.size-1]
	this.min = this.min[:this.size-1]
	this.size--
}

func (this *MinStack) Top() int {
	return this.data[this.size-1]
}

func (this *MinStack) Min() int {
	return this.min[this.size-1]
}
