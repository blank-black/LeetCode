package problems

func gameOfLife(board [][]int)  {
	m := len(board)
	if m == 0 {
		return
	}
	n := len(board[0])
	if n == 0 {
		return
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			num := 0
			if i != 0 {
				if board[i - 1][j] > 0 {
					num ++
				}
				if j != 0 && board[i - 1][j - 1] > 0 {
					num ++
				}
				if j != n - 1 && board[i - 1][j + 1] > 0 {
					num ++
				}
			}
			if i != m - 1 {
				if board[i + 1][j] > 0 {
					num ++
				}
				if j != 0 && board[i + 1][j - 1] > 0 {
					num ++
				}
				if j != n - 1 && board[i + 1][j + 1] > 0 {
					num ++
				}
			}
			if j != 0 && board[i][j - 1] > 0 {
				num ++
			}
			if j != n - 1 && board[i][j + 1] > 0 {
				num ++
			}
			if board[i][j] == 1 && (num < 2 || num > 3) {
				board[i][j] = 2
			}
			if board[i][j] == 0 && num == 3 {
				board[i][j] = -1
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if board[i][j] == -1 {
				board[i][j] = 1
			}
			if board[i][j] == 2 {
				board[i][j] = 0
			}
		}
	}
}