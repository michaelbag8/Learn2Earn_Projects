package main

// Check if num is valid in board[row][col]
func isSafe(b Board, row, col, num int) bool {
	for i := 0; i < 9; i++ {
		if b[row][i] == num || b[i][col] == num {
			return false
		}
	}

	startRow := (row / 3) * 3
	startCol := (col / 3) * 3

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b[startRow+i][startCol+j] == num {
				return false
			}
		}
	}

	return true
}

// Find an empty cell (row, col)
func findEmpty(b Board) (int, int, bool) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b[i][j] == 0 {
				return i, j, true
			}
		}
	}
	return -1, -1, false
}

// Count how many solutions exist
func solveCount(b *Board, count *int) {
	if *count > 1 {
		return
	}

	row, col, empty := findEmpty(*b)
	if !empty {
		*count++
		return
	}

	for num := 1; num <= 9; num++ {
		if isSafe(*b, row, col, num) {
			(*b)[row][col] = num
			solveCount(b, count)
			(*b)[row][col] = 0
		}
	}
}

// Solve board fully (only used after uniqueness check)
func solve(b *Board) bool {
	row, col, empty := findEmpty(*b)
	if !empty {
		return true
	}

	for num := 1; num <= 9; num++ {
		if isSafe(*b, row, col, num) {
			(*b)[row][col] = num
			if solve(b) {
				return true
			}
			(*b)[row][col] = 0
		}
	}
	return false
}

// Ensure exactly one solution, then solve board
func SolveUnique(input Board) (Board, bool) {
	// First count number of solutions
	check := input
	count := 0
	solveCount(&check, &count)

	if count != 1 {
		return Board{}, false
	}

	// Then solve fully
	solved := input
	if !solve(&solved) {
		return Board{}, false
	}

	return solved, true
}
