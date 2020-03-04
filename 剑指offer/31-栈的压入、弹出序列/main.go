package main

//按pushed进栈，按popped出栈,popped是否可以是pushed的弹出序列
//所有数字均不相等
//逐次取popped中的元素为待pop元素→x
//用pushed的顺序模拟入栈，遇到x后pop
type stack []int

func (s *stack) push(x int) {
	*s = append(*s, x)
}

func (s *stack) pop() {
	*s = (*s)[:len(*s)-1]
}

func (s *stack) top() int {
	return (*s)[len(*s)-1]
}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func validateStackSequences(pushed []int, popped []int) bool {
	popIndex := 0
	s := stack{}
	for i := 0; i <= len(pushed); i++ {
		for !s.isEmpty() && s.top() == popped[popIndex] {
			popIndex++
			s.pop()
		}
		if i == len(pushed) {
			break
		}
		s.push(pushed[i])
	}
	return popIndex == len(popped)
}
