/*
Challenge: 
https://codefights.com/arcade/intro/level-11/o2uq6eTuvk7atGadR/description

Given a string, return its encoding defined as follows:

First, the string is divided into the least possible number of disjoint substrings consisting of identical characters
for example, "aabbbc" is divided into ["aa", "bbb", "c"]
Next, each substring with length greater than one is replaced with a concatenation of its length and the repeating character
for example, substring "bbb" is replaced by "3b"
Finally, all the new strings are concatenated together in the same order and a new string is returned.
Example

For s = "aabbbc", the output should be
lineEncoding(s) = "2a3bc".

Input/Output

[time limit] 4000ms (go)
[input] string s

String consisting of lowercase English letters.

Guaranteed constraints:
4 ≤ s.length ≤ 15.

[output] string

Encoded version of s.
*/
func lineEncoding(s string) string {

    count := 0
    prev := rune(s[0])
    res := ""
    for n, r := range s {
        if n == len(s) {
            break
        }
        if (r == prev) {
            count++
        } else {
            if count > 1 {
                res += strconv.Itoa(count) + string(prev)
            } else {
                res += string(prev)
            }
            count = 1
            prev = r
        }
    }
    
    if count > 1 {
        res += strconv.Itoa(count) + string(prev)
    } else {
        res += string(prev)
    }
    
    return res
}
