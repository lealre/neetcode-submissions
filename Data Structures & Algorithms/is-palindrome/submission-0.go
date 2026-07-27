func isPalindrome(s string) bool {

	s = strings.ToLower(s)
	i := 0
	j := len(s) -1 // safe to use len, as there are no chars with more than one byte
	for j > 0 {
		for i < len(s)-1 && s[i] < '0' || s[i] > '9' && s[i] < 'a' || s[i] > 'z' {
			i++
		}

		for j > 0 && s[j] < '0' || s[j] > '9' && s[j] < 'a' || s[j] > 'z' {
			j--
		}

		if j > 0 && i < len(s)-1 && s[i] != s[j] {
			return false
		}

		i++
		j--
	}

	return true
}
