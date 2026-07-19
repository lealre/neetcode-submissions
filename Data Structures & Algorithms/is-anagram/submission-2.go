func isAnagram(s string, t string) bool {
    if len(s) != len(t) {
        return false
    }

    wcs := make(map[byte]int, len(s))
    wct := make(map[byte]int, len(s))
    for i,_ := range s {
        wcs[s[i]]++
        wct[t[i]]++
    }

    for k,v := range wcs {
        if v != wct[k] {
            return false
        }
    }

    return true

}
