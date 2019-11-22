package lintcode

func minCostII(costs [][]int) int {
	// write your code here
	const INT_MAX = int(^uint(0) >> 1)

	if len(costs) == 0 {
		return 0
	}

	n := len(costs)
	K := len(costs[0])
	var j1, j2 int
	dp := make([][]int, n+1)
	for k := range dp {
		dp[k] = make([]int, K)
	}
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	// init dp[0][j]=0
	for j := 0; j < K; j++ {
		dp[0][j] = 0
	}

	for i := 1; i <= n; i++ {
		// m1,m2 是最小值和次小值
		m1, m2 := INT_MAX, INT_MAX
		for j := 0; j < K; j++ {
			if dp[i-1][j] < m1 {
				m2 = m1
				j2 = j1
				m1 = dp[i-1][j]
				j1 = j
			} else {
				if dp[i-1][j] < m2 {
					m2 = dp[i-1][j]
					j2 = j
				}
			}
		}

		for j := 0; j < K; j++ {
			if j != j1 {
				dp[i][j] = dp[i-1][j1] + costs[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j2] + costs[i-1][j]
			}
		}
	}
	result := INT_MAX
	for j := 0; j < K; j++ {
		result = min(dp[n][j], result)
	}
	return result
}
