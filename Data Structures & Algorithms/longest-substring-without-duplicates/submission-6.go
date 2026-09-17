func lengthOfLongestSubstring(s string) int {
        seen := make(map[byte]struct{})
        var left int
        var right int
        var lseq int
        for right < len(s) {
                for {
                    if _, ok := seen[s[right]]; !ok {
                        break
                  }
                        delete(seen, s[left])
                        left++
                }       
                
                seen[s[right]] = struct{}{}
                right++

                lseq = max(lseq, right - left)
        }

        return lseq
}
