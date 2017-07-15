/*
Challenge: CommonCharacterCount
https://codefights.com/arcade/intro/level-3/JKKuHJknZNj4YGL32/description

Given two strings, find the number of common characters between them.

Example

For s1 = "aabcc" and s2 = "adcaa", the output should be
commonCharacterCount(s1, s2) = 3.

Strings have 3 common characters - 2 "a"s and 1 "c".

Input/Output

[time limit] 4000ms (go)
[input] string s1

A string consisting of lowercase latin letters a-z.

Guaranteed constraints:
1 ≤ s1.length ≤ 15.

[input] string s2

A string consisting of lowercase latin letters a-z.

Guaranteed constraints:
1 ≤ s2.length ≤ 15.

[output] integer


*/

func commonCharacterCount(s1 string, s2 string) int {

    sum := 0
    m := make(map[rune]int) 
    for _, r := range s1 {
        m[r]++
    } 
    for _, r := range s2 {
        if m[r] > 0 {
            sum++
            m[r]--
        }
    }
    return sum
}
