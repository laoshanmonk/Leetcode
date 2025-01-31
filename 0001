func twoSum(nums []int, target int) []int {
    // We use hashmap to store (possibleValue, possibleIndex) pairs
    // where possibleValue = target - currentNumber, possibleIndex = currentIndex
    hashmap := make(map[int]int)

    // loop through nums array
    for i, v := range nums {
        // if we found possibleIndex, return result
        if possibleIndex, ok := hashmap[v]; ok {
            return []int{possibleIndex, i}
        } else {
            possibleValue := target - v
            hashmap[possibleValue] = i 
        }
    }
    return nil
}
