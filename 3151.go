func isArraySpecial(nums []int) bool {
    // we look at the sum of each pair
    // if sum = 0 or 2 => false
    // sum = 1 => continue to next
    for i := 0; i < len(nums) - 1; i++ {
        sum := nums[i] % 2 + nums[i+1] % 2
        if sum == 0 || sum == 2 {
            return false
        }
    }
	// all pairs with different parity
    return true
}