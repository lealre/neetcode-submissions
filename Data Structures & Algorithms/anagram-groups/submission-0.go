func groupAnagrams(strs []string) [][]string {

    anagram := make(map[string][]string)
    for _, w := range strs {
        b := []byte(w)
        sort.Slice(b, func(i,j int) bool { return b[i] < b[j]})
        sorted := string(b)
        anagram[sorted] = append(anagram[sorted],w)
    }

    output := make([][]string, 0, len(anagram))
    for _,v := range anagram {
        output = append(output,v)
    }

    return output
}
