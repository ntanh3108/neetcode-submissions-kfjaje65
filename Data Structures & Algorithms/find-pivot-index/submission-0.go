func pivotIndex(nums []int) int {
    totalSum := 0
    for _, num := range nums {
        totalSum += num
    }

    leftSum := 0
    for i, num := range nums {
        // Tính tổng bên phải tại vị trí i hiện tại
        rightSum := totalSum - leftSum - num

        // Nếu tổng trái và phải bằng nhau, ta tìm thấy pivot
        if leftSum == rightSum {
            return i
        }

        // Cập nhật tổng bên trái cho vòng lặp tiếp theo
        leftSum += num
    }

    // Nếu duyệt hết mảng mà không thấy, trả về -1
    return -1
}