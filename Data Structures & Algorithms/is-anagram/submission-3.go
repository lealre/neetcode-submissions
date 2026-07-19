func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    wc := make(map[byte]int, len(s))
    for i,_ := range s {
        wc[s[i]]++
        wc[t[i]]--
    }

    for _,v := range wc {
        if v != 0 {
            return false
        }
    }

    return true

}
