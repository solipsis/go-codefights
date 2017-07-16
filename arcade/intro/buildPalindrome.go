/*
Challenge: buildPalindrome
https://codefights.com/arcade/intro/level-10/ppZ9zSufpjyzAsSEx/description

Given a string, find the shortest possible string which can be achieved by adding characters to the end of initial string to make it a palindrome.

Example

For st = "abcdc", the output should be
buildPalindrome(st) = "abcdcba".

Input/Output

[time limit] 4000ms (go)
[input] string st

A string consisting of lowercase latin letters.

Guaranteed constraints:
3 ≤ st.length ≤ 10.

[output] string
*/
func buildPalindrome(s string) string {
    
    tail := reverse(s)
    for size := 1; size < len(s)+1; size++ {
        for i := 0; i <= len(s) - size; i++ { 
            temp := s + string(tail[i:i+size])
            if temp == reverse(temp) {
                return temp
            }
        }
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
