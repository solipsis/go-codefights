/**
Challenge: texting
https://codefights.com/challenge/dNKLPGwENpZPYbvDp

This is a reverse challenge, enjoy!

Input/Output

[time limit] 4000ms (go)
[input] string s

Guaranteed constraints:
0 ≤ s.length < 200.

[output] string
**/
import "strings"
func texting(s string) (z string) {
    m := map[int]string{2:"abc", 3:"def", 4:"ghi", 5:"jkl",6:"mno",7:"pqrs",8:"tuv", 9:"wxyz", 0:" "}
    o := strconv.Itoa
    
    // for each rune in string
    for i := 0; i < len(s); i++ {     
        c := 1
        // count how many in a row
        for i < len(s)-1 && s[i] == s[i+1] {
            i++
            c++
        }
        // append to result with different case if more than 1 in a row
        for k, v := range m {
            x := strings.IndexByte(v, s[i])
            if x >= 0 {
                if c > 1 {
                    z += o(c)
                }
                z += o(k) + o(x+1)
            }
        }
    } 
    return
}
