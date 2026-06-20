func isValidSudoku(board [][]byte) bool {
	rowSeen := [9]map[byte]bool{}
	colSeen := [9]map[byte]bool{}
	boxSeen := [9]map[byte]bool{}

	for i := range 9 {
		rowSeen[i] = make(map[byte]bool)
	}
	for i := range 9 {
		colSeen[i] = make(map[byte]bool)
	}
	for i := range 9 {
		boxSeen[i] = make(map[byte]bool)
	}

	for r := range 9 {
		for c := range 9 {
			if board[r][c] == '.' {
				continue
			}
			currentVal := board[r][c]
			boxRow := r / 3
			boxCol := c / 3
			boxIndex := 3*boxRow + boxCol

			if _, ok := rowSeen[r][currentVal]; ok {
				return false
			}
			rowSeen[r][currentVal] = true

			if _, ok := colSeen[c][currentVal]; ok {
				return false
			}
			colSeen[c][currentVal] = true

			if _, ok := boxSeen[boxIndex][currentVal]; ok {
				return false
			}
			boxSeen[boxIndex][currentVal] = true
		}
	}
	return true
}
