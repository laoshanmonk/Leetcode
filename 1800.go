func maxAscendingSum(nums []int) int {
    max := 0
    acc := 0
    // loop through and accumualte the sum
    pre := -10000
    for _, num := range nums {
        if num > pre {
            acc = acc + num
        } else {
            if acc > max {
                max = acc
            }
            acc = num
        }
        pre = num
    }
    if acc > max {
        return acc
    }
    return max
}