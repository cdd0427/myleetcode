package main

//需修改，复杂度太高
import "fmt"

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
	i := 0
	for ; i < this.size && x < this.min[i]; i++ {
	}
	if i == this.size {
		this.min = append(this.min, x)
		this.size++
		return
	}
	rear := append([]int{}, this.min[i:]...)
	this.min = append(this.min[:i], x)
	this.min = append(this.min, rear...)
	this.size++
}

func (this *MinStack) deleteMin(x int) {
	l, r, mid := 0, this.size-1, 0
	for l <= r {
		mid = (l + r) / 2
		if this.min[mid] == x {
			break
		}
		if this.min[mid] > x {
			l = mid + 1
		}
		if this.min[mid] < x {
			r = mid - 1
		}
	}
	this.min = append(this.min[:mid], this.min[mid+1:]...)
	return
}

func (this *MinStack) Pop() {
	this.deleteMin(this.data[this.size-1])
	this.data = this.data[:this.size-1]
	this.size--
}

func (this *MinStack) Top() int {
	return this.data[this.size-1]
}

func (this *MinStack) GetMin() int {
	return this.min[this.size-1]
}

func main() {
	s := Constructor()
	s.Push(-2)
	fmt.Println("push: ", s.data, " ", s.min, " ", s.size)
	s.Push(0)
	fmt.Println("push: ", s.data, " ", s.min, " ", s.size)
	s.Push(-3)
	fmt.Println("push: ", s.data, " ", s.min, " ", s.size)
	fmt.Println("min:", s.GetMin())
	s.Pop()
	fmt.Println("pop: ", s.data, " ", s.min, " ", s.size)
	fmt.Println("top:", s.Top())
	fmt.Println("min:", s.GetMin())
}
