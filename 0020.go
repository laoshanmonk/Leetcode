func isValid(s string) bool {
    // use stack to pop and push open bracket
    stack := make([]rune, 0)
    // '(' = 40, ')' = 41, '{' = 123, '}' = 125, '[' = 91, ']' = 93
    for _, b := range s {
        // if b is open bracket, we push it to the stack
        if b == 40 || b == 91 || b == 123 {
            stack = append(stack, b)
        } else {
            // case where nothing to be pop
            if len(stack) == 0 {
                return false
            }
            // if b is close bracket, we pop if it match, otherwise return false
            openBracket := stack[len(stack)-1]
            if b == 41 && openBracket == 40 || b == 93 && openBracket == 91 || b == 125 && openBracket == 123 {
                stack = stack[:len(stack)-1]
            } else {
                return false
            }
        }
    }
    return len(stack) == 0
}