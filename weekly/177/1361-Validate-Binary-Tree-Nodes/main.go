package main

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	visited := make([]bool, n+1)
	for i := 0; i <= n; i++ {
		if leftChild[i] != -1 {
			if visited[leftChild[i]] != false {
				return false
			} else {
				visited[leftChild[i]] = true
			}
		}
		if rightChild[i] != -1 {
			if visited[rightChild[i]] != false {
				return false
			} else {
				visited[rightChild[i]] = true
			}
		}
	}
	count := 0
	for i := 0; i <= n; i++ {
		if visited[i] == false {
			count++
		}
		if count > 1 {
			return false
		}
	}
	if count == 0 {
		return false
	}
	return true
}
