func minSubArrayLen(target int, nums []int) int {
	length := math.MaxInt64
	l := 0
	currSum := 0

	for r := 0; r < len(nums); r++ {
		currSum = currSum + nums[r]
		for currSum >= target {
			length = findMin(length, r-l+1)
			currSum = currSum - nums[l]
			l++
		}
	}

	if length == math.MaxInt64{
		return 0
	}

	return length
}

func findMin(a, b int) int {
	if a < b {
		return a
	}

	return b
}