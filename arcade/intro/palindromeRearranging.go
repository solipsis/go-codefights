/*
Challenge: palindromeRearranging
https://codefights.com/arcade/intro/level-4/Xfeo7r9SBSpo3Wico/description

Given a string, find out if its characters can be rearranged to form a palindrome.

Example

For inputString = "aabb", the output should be
palindromeRearranging(inputString) = true.

We can rearrange "aabb" to make "abba", which is a palindrome.

Input/Output

[time limit] 4000ms (go)
[input] string inputString

A string consisting of lowercase English letters.

Guaranteed constraints:
1 ≤ inputString.length ≤ 50.

[output] boolean

true if the characters of the inputString can be rearranged to form a palindrome, false otherwise.
*/
func palindromeRearranging(s string) bool {

    m := make(map[rune]int)
    for _, r := range s {
        m[r]++
    }
    odd := 0
    for _, v := range m {
        if v % 2 == 1 {
            odd++
        }
    }
    return odd < 2
}
