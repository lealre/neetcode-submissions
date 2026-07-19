func isAnagram(s string, t string) bool {
    return naiveSort(s) == naiveSort(t)
}

func naiveSort(str string) string {
    bstr := []byte(str)
    for i, v := range bstr {
        minor := v
        minorIndex := i
        for j := i+1; j < len(bstr); j++ {
            if bstr[j] < minor {
                minor = bstr[j]
                minorIndex = j
            }
        }
        if minorIndex != i {
            bstr[minorIndex], bstr[i] = bstr[i], bstr[minorIndex]
        }
    }

    return string(bstr)
}

