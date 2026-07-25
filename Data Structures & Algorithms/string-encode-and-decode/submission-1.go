const (
        SEPARATOR = "=="
        STR_DECIMAL_LEN = 3
)

type Solution struct{}

func (s *Solution) Encode(strs []string) string {
        var encoded strings.Builder
        for _, str := range strs {
                strlen := fmt.Sprintf("%03d", len(str))
                s := "==" + strlen + str
                encoded.WriteString(s)
        }

        return encoded.String()
}

func (s *Solution) Decode(encoded string) []string {
        decoded := []string{}
        i := 0
        markLen := len(SEPARATOR) + STR_DECIMAL_LEN
        for i < len(encoded) {
                if string(encoded[i]) + string(encoded[i+1]) == SEPARATOR {
                        if n, err := strconv.Atoi(encoded[i+len(SEPARATOR):i+markLen]); err == nil {
                                word := encoded[i+markLen:i+markLen+n]
                                decoded = append(decoded, word)
                                i += markLen + n
                        }
                }
        }

        return decoded
}

