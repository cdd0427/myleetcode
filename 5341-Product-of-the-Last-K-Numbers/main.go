package main

import "fmt"

type ProductOfNumbers struct {
	data    []int
	product []int
	sum     int
	zero    int
}

func Constructor() ProductOfNumbers {
	return ProductOfNumbers{
		data:    []int{},
		product: []int{},
		sum:     1,
	}
}

func (this *ProductOfNumbers) Add(num int) {
	this.data = append(this.data, num)
	if num == 0 {
		this.zero = len(this.data) - 1
		this.product = []int{}
		this.sum = 1
		return
	}
	this.sum *= num
	if len(this.product) == 0 {
		this.product = append(this.product, num)
		return
	}
	this.product = append(this.product, num*this.product[len(this.product)-1])
}

func (this *ProductOfNumbers) GetProduct(k int) int {
	if len(this.product)-k == 0 {
		return this.sum
	}
	if len(this.product)-k < 0 {
		return 0
	}
	return this.sum / this.product[len(this.product)-k-1]
}

func main() {
	obj := Constructor()
	obj.Add(3)
	fmt.Println("add：", obj.data, obj.product, obj.sum, " ", obj.zero)
	obj.Add(0)
	fmt.Println("add：", obj.data, obj.product, obj.sum, " ", obj.zero)
	obj.Add(2)
	fmt.Println("add：", obj.data, obj.product, obj.sum, " ", obj.zero)
	obj.Add(5)
	fmt.Println("add：", obj.data, obj.product, obj.sum, " ", obj.zero)
	obj.Add(4)
	fmt.Println("add：", obj.data, obj.product, obj.sum, " ", obj.zero)
	fmt.Println("get2: ", obj.GetProduct(2))
	fmt.Println("get3: ", obj.GetProduct(3))
	fmt.Println("get4: ", obj.GetProduct(4))
	obj.Add(8)
	fmt.Println("add：", obj.data, obj.product, obj.sum, " ", obj.zero)
	fmt.Println("get2: ", obj.GetProduct(2))
}

/**
 * Your ProductOfNumbers object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(num);
 * param_2 := obj.GetProduct(k);
 */
