func trap(height []int) int {
	if len(height) == 0 {
		return 0
	}

	left, right := 0, len(height)-1
	maxLeft, maxRight := height[0], height[len(height)-1]
	totalWater := 0

	for left < right {
		if maxLeft < maxRight {
			left++
			if height[left] > maxLeft {
				maxLeft = height[left]
			} else {
				totalWater = totalWater + (maxLeft - height[left])
			}
		} else {
			right--
			if height[right] > maxRight {
				maxRight = height[right]
			} else {
				totalWater = totalWater + (maxRight - height[right])
			}
		}
	}

	return totalWater
}