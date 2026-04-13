func isPalindrome(x int) bool {
    // Guard clause 1: Số âm chắc chắn không đối xứng.
    // Guard clause 2: Số tận cùng bằng 0 (ví dụ 10, 100) thì không thể đối xứng, 
    //                 ngoại trừ chính số 0.
    if x < 0 || (x%10 == 0 && x != 0) {
        return false
    }

    revertedNumber := 0
    // Vòng lặp "gập giấy": Dừng lại khi nửa bên phải (reverted) 
    // lớn hơn hoặc bằng nửa bên trái (x)
    for x > revertedNumber {
        revertedNumber = revertedNumber*10 + x%10
        x = x / 10
    }

    // Nếu số có độ dài chẵn: x == revertedNumber (vd: 1221 -> x=12, reverted=12)
    // Nếu số có độ dài lẻ: x == revertedNumber/10 (vd: 12321 -> x=12, reverted=123. Ta vứt số 3 ở giữa đi)
    return x == revertedNumber || x == revertedNumber/10
}