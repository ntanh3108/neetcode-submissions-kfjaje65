func isPalindrome(s string) bool {
	
	var chars []string

	for i := 0; i < len(s); i++ {
		if (s[i] >= 'a' && s[i] <= 'z') || (s[i] >= 'A' && s[i] <= 'Z') || (s[i] >= '0' && s[i] <= '9') {
			chars = append(chars, strings.ToLower(string(s[i]))) 
		}
	}

	l := 0
	r := len(chars) - 1

	for l < r {
		if chars[l] == chars[r]{
			l++
			r--
		} else {
			return false
		}
	}

	return true
}
