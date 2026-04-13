func containsNearbyDuplicate(nums []int, k int) bool {
     window := make(map[int]bool) // Cur window of size <= k
    L := 0

    for R := 0; R < len(nums); R++ {
        if R-L > k {
            delete(window, nums[L])
            L++
        }
        if _, ok := window[nums[R]]; ok {
            return true
        }
        window[nums[R]] = true
    }
    return false
}
