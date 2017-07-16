/*
Challenge: isBeautifulString
https://codefights.com/arcade/intro/level-10/PHSQhLEw3K2CmhhXE/description

A string is said to be beautiful if b occurs in it no more times than a; c occurs in it no more times than b; etc.

Given a string, check whether it is beautiful.

Example

For inputString = "bbbaacdafe", the output should be
isBeautifulString(inputString) = true;
For inputString = "aabbb", the output should be
isBeautifulString(inputString) = false;
For inputString = "bbc", the output should be
isBeautifulString(inputString) = false.
Input/Output

[time limit] 4000ms (go)
[input] string inputString

A string of lowercase letters.

Guaranteed constraints:
3 ≤ inputString.length ≤ 50.

[output] boolean
*/
func isBeautifulString(inputString string) bool {
    
    m := make(map[rune]int)
    for _, r := range inputString {
        m[r]++
    }
    for i := 1; i < 26; i++ {
        r := rune(i + 97)
        r1 := rune(i + 96)
        if (m[r] > m[r1]) {
            return false
        }
    }
    
    return true
}
