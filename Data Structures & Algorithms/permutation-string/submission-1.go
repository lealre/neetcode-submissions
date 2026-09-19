import "slices"
func checkInclusion(s1 string, s2 string) bool {

        left := 0
        rigth := len(s1)-1
        s1sorted := []byte(s1)
        slices.Sort(s1sorted)
        s1macth := string(s1sorted)

        fmt.Printf("To macth sorted: %s\n", s1macth)

        for rigth < len(s2) {
                sub :=[]byte(s2[left:rigth+1]) 

                slices.Sort(sub)

                fmt.Printf("Comparing aginst substring: %v\n", string(sub))

                if string(sub) == s1macth {
                        return true
                }

                left++
                rigth ++
        }


        return false
}
