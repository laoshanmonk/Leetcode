func strStr(haystack string, needle string) int {
    sliceLength := len(needle)
    // loop through haystack with needle length
    for i := 0; i < len(haystack) - sliceLength + 1; i++ {
        toCheck := haystack[i:i+sliceLength]
        if toCheck == needle {
            return i
        }
    }
    return -1
}