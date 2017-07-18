/**
Challenge: meaningOfLife
https://codefights.com/challenge/iuePposGo5AYBn8w9

This is a reverse challenge, enjoy!

Input/Output

[time limit] 4000ms (go)
[input] string n

A string of alphanumeric characters.

Guaranteed constraints:
1 ≤ n.length ≤ 9.

[output] integer64

The answer is guaranteed to be less than 251.
**/
import "math"
import "strings"

func meaningOfLife(n string) int64 {
    
    s := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
    
    var tot int64
    exp := len(n)
    for _, r := range n {
        exp--
        tot += int64(strings.Index(s, string(r)) * int(math.Pow(42, float64(exp))))
    }
    
    
    return tot
}
