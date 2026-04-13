func characterReplacement(s string, k int) int {
    l := 0
    maxFrequent := 0
    countChar := make(map[byte]int)
    maxLength := 0

    for r := 0; r < len(s); r++ {
        countChar[s[r]]++
        
        // Cập nhật maxFrequent
        if countChar[s[r]] > maxFrequent {
            maxFrequent = countChar[s[r]]
        }

        // Kiểm tra tính hợp lệ của cửa sổ
        length := r - l + 1
        if length - maxFrequent > k {
            countChar[s[l]]-- // Nhớ: Phải trừ trước!
            l++               // Rồi mới tăng l
        }

        // Cập nhật kỷ lục
        // Lúc này r - l + 1 luôn là kích thước cửa sổ hiện tại (dù L có bị dịch hay không)
        currentWindowSize := r - l + 1
        if currentWindowSize > maxLength {
            maxLength = currentWindowSize
        }
    }

    return maxLength
}