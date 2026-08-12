func isValidSudoku(board [][]byte) bool {

        // this will map each column in the unique pass we want
        cmap := make(map[string]bool)
        // this will map the ineer matrix (3x3) duplications
        inmap := make(map[string]bool)
        for i:=0;i<len(board);i++ {
                rmap := make(map[byte]bool)
                for j:=0;j<len(board[i]);j++ {
                        v := board[i][j] 
                        if string(v) == "." {
                                continue
                        }

                        // hanldes duplication on rows
                        if _, ok := rmap[v]; ok {
                                return false
                        }
                        rmap[v] = true


                        // handles duplication on columns
                        jchar := strconv.Itoa(j)
                        key := jchar + ":" + string(v)
                        if _, ok := cmap[key]; ok {
                                return false
                        }
                        cmap[key] = true

                        // hanldes duplication on inner matrix
                        switch {
                        // handles (0x0) til (2x0)
                        case i <= 2 && j <= 2:
                                key := "00:" + string(v)
                                if _, ok := inmap[key]; ok {
                                        return false
                                }
                                inmap[key] = true
                        case i > 2 && i <=5 &&  j <= 2:
                                key := "10:" + string(v)
                                if _, ok := inmap[key]; ok {
                                        return false
                                }
                                inmap[key] = true
                        case i > 5 && j <= 2:
                                key := "20:" + string(v)
                                if _, ok := inmap[key]; ok {
                                        return false
                                }
                                inmap[key] = true
                        // hanldes (0X1) till (2x1)
                        case i <= 2 && j > 2 && j <= 5:
                                key := "01:" + string(v)
                                if _, ok := inmap[key]; ok {
                                        return false
                                }
                                inmap[key] = true
                        case i > 2 && i<= 5 && j > 2 && j<= 5:
                                key := "11:" + string(v)
                                if _, ok := inmap[key]; ok {
                                        return false
                                }
                                inmap[key] = true
                        case i > 5  && j > 2 && j<= 5:
                                key := "21:" + string(v)
                                if _, ok := inmap[key]; ok {
                                        return false
                                }
                                inmap[key] = true
                        // handles (0x2) till (2x2)
                        case i <= 2  && j> 5:
                                key := "02:" + string(v)
                                if _, ok := inmap[key]; ok {
                                        return false
                                }
                                inmap[key] = true
                        case i > 2 && i<= 5 && j> 5:
                                key := "12:" + string(v)
                                if _, ok := inmap[key]; ok {
                                        return false
                                }
                                inmap[key] = true
                        case i >  5 && j< 5:
                                key := "22:" + string(v)
                                if _, ok := inmap[key]; ok {
                                        return false
                                }
                                inmap[key] = true
                        }

                }
        }

        return true

}
