func maxSubarraySumCircular(nums []int) int {
    totalSum := 0
    globalMax := nums[0]
    globalMin := nums[0]
    currMax, currMin := 0, 0

    for i := 0; i < len(nums); i++ {
        currMax = findMax(currMax + nums[i], nums[i])
        currMin = findMin(currMin + nums[i], nums[i])
        globalMax = findMax(globalMax, currMax)
        globalMin = findMin(globalMin, currMin)
        totalSum += nums[i]
    }

    if globalMax > 0 {
        return findMax(globalMax, totalSum - globalMin)
    }

    return globalMax
}

func findMin (a, b int) int {
    if a < b {
        return a
    }

    return b
}

func findMax (a, b int) int {
    if a < b {
        return b
    }

    return a
}