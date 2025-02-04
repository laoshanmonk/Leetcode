func longestCommonPrefix(strs []string) string {
    // assume first one is common index
    commonPrefix := strs[0]
    commonPrefixLength := len(commonPrefix)
    // loop through to find longest
    for _, str := range strs {
        if commonPrefixLength == 0 {
            return ""
        }
        if len(str) < commonPrefixLength {
            commonPrefix = str[:common(str, commonPrefix)]
        }
        commonPrefix = str[:common(commonPrefix, str)]
        commonPrefixLength = len(commonPrefix)
    }
    return commonPrefix
}

// helper function
// return the length of common prefix
func common(str1 string, str2 string) int {
    byteSlice1 := []byte(str1)
    byteSlice2 := []byte(str2)
    index := 0
    for i, v := range byteSlice1 {
        if v != byteSlice2[i] {
            return index
        }
        index++
    }
    return index
}