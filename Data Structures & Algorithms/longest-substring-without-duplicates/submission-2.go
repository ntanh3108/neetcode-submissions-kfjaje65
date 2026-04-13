func lengthOfLongestSubstring(s string) int {
	if len(s) == 0{
		return 0
	}

	max := 1
	l := 0
	window := make(map[byte]int)
	for r := 0; r < len(s); r++ {
		if pos, ok := window[s[r]]; ok && pos >= l {
			l = pos + 1
		}

		window[s[r]] = r
		length := r-l+1
		if length > max {
			max = length
		}
	}

	return max
}