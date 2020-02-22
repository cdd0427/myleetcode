package main

type Cashier struct {
	products []int
	prices   []int
	n        int
	discount int
	cur      int
}

func Constructor(n int, discount int, products []int, prices []int) Cashier {
	return Cashier{
		products: products,
		prices:   prices,
		n:        n,
		discount: discount,
		cur:      0,
	}
}

func (this *Cashier) findIndex(n int) int {
	for i := 0; i < len(this.products); i++ {
		if this.products[i] == n {
			return i
		}
	}
	return 0
}

func (this *Cashier) GetBill(product []int, amount []int) float64 {
	this.cur++
	flag := 0
	if this.cur == this.n {
		flag = 1
		this.cur = 0
	}
	res := 0.0
	for i := 0; i < len(product); i++ {
		index := this.findIndex(product[i])
		res += float64(this.prices[index] * amount[i])
	}
	res = res - float64(flag*this.discount)*res/100.0
	return res
}

/**
 * Your Cashier object will be instantiated and called as such:
 * obj := Constructor(n, discount, products, prices);
 * param_1 := obj.GetBill(product,amount);
 */
