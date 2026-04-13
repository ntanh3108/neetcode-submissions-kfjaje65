func numOfSubarrays(arr []int, k int, threshold int) int {
	currentSum := 0
	l := 0
	targetSum := k * threshold
	count := 0

	for r := 0; r < len(arr); r++ {
		currentSum += arr[r]
		if r-l+1 == k {
			if currentSum >= targetSum{
				count++
			}
			currentSum -= arr[l]
			l++
		}
	}

	return count
}