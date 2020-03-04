package main

func isMatch(s string, p string) bool {
	if len(s) == 0 && len(p) == 0 {
		return true
	}
	row, col := len(s)+1, len(p)+1
	dp := make([][]bool, row)
	for i := 0; i <= row; i++ {
		opt := make([]bool, col)
		dp[i] = opt
	}
	dp[0][0] = true
	for i := 1; i < col; i++ {
		if p[i-1] == '*' && dp[0][i-2] {
			dp[0][i] = true
		}
	}
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			ss, pp := s[i-1], p[j-1]
			if ss == pp {
				dp[i][j] = dp[i-1][j-1]
			} else {
				if pp == '.' {
					dp[i][j] = dp[i-1][j-1]
				} else if pp == '*' {
					if j >= 2 {
						opt := p[j-2]
						if opt == ss || opt == '.' {
							dp[i][j] = dp[i-1][j] || dp[i][j-1]
						}
						dp[i][j] = dp[i][j] || dp[i][j-2]
					}
				} else {
					dp[i][j] = false
				}
			}
		}
	}
	return dp[row-1][col-1]
}
