func isNumeric(char byte) bool {
	return char >= '0' && char <= '9'
}

func rowHasDuplicates(row []byte) bool {
	byteMap := map[byte]int{}

	for _, value := range row {
		byteMap[value]++

		if byteMap[value] > 1 && isNumeric(value) {
			return true
		}
	}

	return false
}

func columnHasDuplicates(board [][]byte, columnIndex int) bool {
	byteMap := map[byte]int{}

	for i := 0; i < len(board); i++ {
		value := board[i][columnIndex]
		byteMap[value]++

		if byteMap[value] > 1 && isNumeric(value) {
			return true
		}
	}

	return false
}

func boxHasDuplicates(board [][]byte, startRow, startCol int) bool {
	byteMap := map[byte]int{}

	for row := startRow; row < startRow+3; row++ {
		for col := startCol; col < startCol+3; col++ {
			value := board[row][col]

			if !isNumeric(value) {
				continue
			}

			byteMap[value]++

			if byteMap[value] > 1 {
				return true
			}
		}
	}

	return false
}


func isValidSudoku(board [][]byte) bool {
	// Check rows and columns
	for i := 0; i < 9; i++ {
		if rowHasDuplicates(board[i]) {
			return false
		}

		if columnHasDuplicates(board, i) {
			return false
		}
	}

	// Check 3x3 boxes
	for row := 0; row < 9; row += 3 {
		for col := 0; col < 9; col += 3 {
			if boxHasDuplicates(board, row, col) {
				return false
			}
		}
	}

	return true
}
