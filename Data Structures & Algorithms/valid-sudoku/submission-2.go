func isValidSudoku(board [][]byte) bool {

	cmap := make(map[string]bool)  // this will map each column in the unique pass
	inmap := make(map[string]bool) // this will map the ineer matrix (3x3) duplications

	for i := 0; i < len(board); i++ {
		rmap := make(map[byte]bool) // this will map each row in the unique pass
		for j := 0; j < len(board[i]); j++ {
			v := board[i][j]
			if string(v) == "." {
				continue
			}

			// rows duplications
			if _, ok := rmap[v]; ok {
				return false
			}
			rmap[v] = true

			// columns duplications
			jchar := strconv.Itoa(j)
			ckey := jchar + ":" + string(v)
			if _, ok := cmap[ckey]; ok {
				return false
			}
			cmap[ckey] = true

			// inner matrix duplications
			pkey := innerMatrixPrefixKey(i, j)
			inkey := pkey + ":" + string(v)
			if _, ok := inmap[inkey]; ok {
				return false
			}
			inmap[inkey] = true
		}

	}

	return true

}

func innerMatrixPrefixKey(i, j int) string {
	switch {
	case i <= 2 && j <= 2:
		return "00"
	case i > 2 && i <= 5 && j <= 2:
		return "10"
	case i > 5 && j <= 2:
		return "20"
	case i <= 2 && j > 2 && j <= 5:
		return "01"
	case i > 2 && i <= 5 && j > 2 && j <= 5:
		return "11"
	case i > 5 && j > 2 && j <= 5:
		return "21"
	case i <= 2 && j > 5:
		return "02"
	case i > 2 && i <= 5 && j > 5:
		return "12"
	case i > 5 && j < 5:
		return "22"
	default:
		return ""
	}
}
