package problems
import "math"
var M, N int
var Pos [4][2]int
func LongestIncreasingPath(matrix [][]int) int {
	M = len(matrix)
	if M == 0 {
		return 0
	}
	N = len(matrix[0])
	if N == 0 {
		return 0
	}
	Pos = [4][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	nmatrix := make([][]int, M)
	for i := 0; i < M; i++ {
		nmatrix[i] = make([]int, N)
	}
	m := 0
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			helper329(matrix, nmatrix, i, j, math.MinInt32)
		}
	}
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			m = max(m, nmatrix[i][j])
		}
	}
	return m
}

func helper329(matrix [][]int, nmatrix [][]int, m, n, v int) int {
	if m >= M || m < 0 || n >= N || n < 0 || v >= matrix[m][n] {
		return 0
	}
	if nmatrix[m][n] > 0 {
		return nmatrix[m][n]
	}
	mm := 1
	for i := 0; i < 4; i++ {
		mm = max(mm, helper329(matrix, nmatrix, m + Pos[i][0], n + Pos[i][1], matrix[m][n]) + 1)
	}
	nmatrix[m][n] = mm
	return mm
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}