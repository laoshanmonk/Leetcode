func isPalindrome(x int) bool {
    // base case1: x is negative => false
    if x < 0 {
        return false
    }
    // base case2: x is 0-9 => true
    if x < 10 {
        return true
    }
    // base case3: x mod 10 = 0 => false
    if x % 10 == 0 {
        return false
    }
    // general case
    // use integer array to store all the digits
    array := []int{}
    for x > 0 {
        // append last digit into array
        array = append(array, x % 10)
        x = x / 10
    }
    // now we loop through array to check palindrome
    start := 0
    end := len(array) - 1
    for start < end {
        if array[start] != array[end] {
            return false
        }
        start++
        end--
    }
    return true
}