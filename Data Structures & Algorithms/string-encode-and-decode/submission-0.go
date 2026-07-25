type Solution struct{}

func (s *Solution) Encode(strs []string) string {
        var encoded string
        for _, str := range strs {
                strlen := fmt.Sprintf("%03d", len(str))
                encoded += "==" + strlen + str
        }

        return encoded
}

func (s *Solution) Decode(encoded string) []string {
        decoded := []string{}
        i := 0
        for i < len(encoded) {
                fmt.Printf("Iterating for i=%d\n", i)
                if string(encoded[i]) + string(encoded[i+1]) == "==" {
                        if n, err := strconv.Atoi(encoded[i+2:i+5]); err == nil {
                                fmt.Println(encoded[i+5:i+5+n])
                                word := encoded[i+5:i+5+n]
                                decoded = append(decoded, word)
                                i += 5 + n
                        }
                }
        }

        return decoded
}
