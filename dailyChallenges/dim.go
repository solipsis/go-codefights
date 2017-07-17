/**
Challenge: dim
https://codefights.com/challenge/pngP3vJ74YdtyBaWb

This is a reverse challenge, enjoy!

Input/Output

[time limit] 4000ms (go)
[input] integer64 n

Guaranteed constraints:
0 ≤ n ≤ 1015.

[output] integer
**/
import "math"
func dim(n int64) int {
    if n == 0 {
        return 1
    }
    return int(math.Log2(float64(n)))+1
}
