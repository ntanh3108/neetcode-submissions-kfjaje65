func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	maxHeight := 0

	for l < r {
		currHeight := findMin(height[l], height[r]) * (r - l)
		if currHeight > maxHeight {
			maxHeight = currHeight
		}

		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return maxHeight
}

func findMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}