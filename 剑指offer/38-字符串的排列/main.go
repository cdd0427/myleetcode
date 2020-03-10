package main

type stack struct {
	data    []byte
	maxSize int
}

func (s *stack) pop() {
	s.data = s.data[:len(s.data)-1]
}

func (s *stack) push(x byte) {
	s.data = append(s.data, x)
}

func (s *stack) isFull() bool {
	return len(s.data) == s.maxSize
}

func permutation(s string) []string {
	stack := stack{
		data:    make([]byte, 0),
		maxSize: len(s),
	}
	visited := make([]bool, len(s))
	for i := 0; i < len(s); i++ {
		if visited[i] == false {
			stack.push(s[i])
		}
	}
	return nil
}
