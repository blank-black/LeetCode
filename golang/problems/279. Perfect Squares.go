package problems

// dynamic
// time o(nlogn)	space o(n)

func numSquares(n int) int {
	if n <= 0 {
		return 0
	}
	arr := make([]int, n + 1)
	for i := 0; i <= n; i++ {
		arr[i] = i
	}
	for i := 1; i <= n; i++ {
		for j := 1; i >= j*j; j++ {
			arr[i] = min(arr[i], arr[i - j * j] + 1)
		}
	}
	return arr[n]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}