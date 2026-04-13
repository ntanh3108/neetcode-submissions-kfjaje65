func twoSum(numbers []int, target int) []int {
    left, right := 0, len(numbers)-1

    // O(n) loop using two-pointer technique
    for left < right {
        sum := numbers[left] + numbers[right]

        if sum == target {
            return []int{left + 1, right + 1}  // return 1-based indices
        } else if sum < target {
            left++  // move left to increase the sum
        } else {
            right--  // move right to decrease the sum
        }
    }

    // unreachable since one solution is guaranteed
    return []int{}
}