/*
Challenge: Reverse Parenthesis
https://codefights.com/arcade/intro/level-3/3o6QFqgYSontKsyk4/description

You have a string s that consists of English letters, punctuation marks, whitespace characters, and brackets. It is guaranteed that the parentheses in s form a regular bracket sequence.

Your task is to reverse the strings contained in each pair of matching parentheses, starting from the innermost pair. The results string should not contain any parentheses.

Example

For string s = "a(bc)de", the output should be
reverseParentheses(s) = "acbde".

Input/Output

[time limit] 4000ms (go)
[input] string s

A string consisting of English letters, punctuation marks, whitespace characters and brackets. It is guaranteed that parentheses form a regular bracket sequence.

Constraints:
5 ≤ s.length ≤ 55.

[output] string

*/
import "strings"
func reverseParentheses(s string) string {

    for strings.Contains(s, "(") {
        l := strings.LastIndex(s, "(")
        r := 0
        for n, ru := range s[l:] {
            if ru == ')' {
                r = n+l
                break
            }
        }
        s = s[:l] + reverse(s[l+1:r]) + s[r+1:]
    }
    return s
}


func reverse(s string) string {
    
    r := []rune(s)
    for i, j := 0, len(s)-1; i < len(s)/2; i, j = i+1, j-1 {
        r[i], r[j] = r[j], r[i]
    }
    return string(r)
}
