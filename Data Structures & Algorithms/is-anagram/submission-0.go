func isAnagram(s string, t string) bool {

    
    sb := []byte(s)
    for i, v := range sb {
        minor := v
        for j := i+1; j < len(sb); j++ {
            if sb[j] < minor {
                minor = sb[j]
                sb[j], sb[i] = sb[i], sb[j]
            }
        }
    }

    tb := []byte(t)
    for i, v := range tb {
        minor := v
        for j := i+1; j < len(tb); j++ {
            if tb[j] < minor {
                minor = tb[j]
                tb[j], tb[i] = tb[i], tb[j]
            }
        }
    }



    if string(sb) == string(tb) {
        return true
    }


    return false
}
