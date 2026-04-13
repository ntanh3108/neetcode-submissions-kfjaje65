func productExceptSelf(nums []int) []int {
    res := make([]int, len(nums))
    product := 1
    zeros := 0

    for _, num := range nums {
        if num == 0 {
            zeros += 1
            continue
        }
        product = product*num
    }

    if zeros == 1 {
        for key, num := range nums {
            if num == 0 {
                res[key] = product
            }
        }
    } else if zeros == 0 {
        for key, num := range nums {
            res[key] = product/num
        }
    }

    return res
}
