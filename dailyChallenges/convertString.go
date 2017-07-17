/**
Challenge: convertString
https://codefights.com/challenge/DF6j9AcQSCGtLonyM

You are given the string s. Your friend just asked you if it's possible to change the string from s to a string t by removing some characters from it. You're a pro at programming, so you decided to create a program to determine this.

Example

For s = "ceoydefthf5iyg5h5yts" and t = "codefights", the output should be
convertString(s, t) = true.
For s = "addbyca" and t = "abcd", the output should be
convertString(s, t) = false.
Input/Output

[time limit] 4000ms (go)
[input] string s

A string with alphanumeric characters.

Guaranteed constraints:
1 ≤ s.length ≤ 1000.

[input] string t

A string with alphanumeric characters.

Guaranteed constraints:
1 ≤ t.length ≤ 1000.

[output] boolean

Return true if it is possible to convert s to t, otherwise return false.
**/
func convertString(s,  t string) (x bool) {
    for n := range s {
        if s[n] == t[0] {
            t = t[1:]
        }
        if len(t) < 1 {return !x}
        
    }
    return
}
