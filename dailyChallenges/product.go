/**
Challenge: product
https://codefights.com/challenge/MAKAXvQK7v95B3j9R

This is a reverse challenge, enjoy!

Input/Output

[time limit] 4000ms (go)
[input] string s

A string of lowercase English letters.

Guaranteed constraints:
1 ≤ s.length ≤ 20.

[output] integer
**/
func product(s string) int {
    x := len(s)/2
    return x*x
}
