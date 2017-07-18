/**
Challenge: totalOr
https://codefights.com/challenge/NXE7XwgYZvCXRY6PW

Given a non-negative integer k, you task is to calculate the result of the bitwise OR operation of all the integers from 0 to k, inclusive.

Example

For k = 5, the output should be
totalOr = 7.

0 | 1 | 2 | 3 | 4 | 5 = 02 | 12 | 102 | 112 | 1002 | 1012 = 1112 = 7.

Input/Output

[time limit] 4000ms (go)
[input] integer k

Guaranteed constraints:
0 ≤ k < 231

[output] integer

Guaranteed constraints:
result < 231
**/
func totalOr(k int) (x int) {
    for k>0 {
        k /= 2
        x = x*2+1
    }
    return
   // return (1 << 32 -1) >> uint(32 - int(math.Log2(float64(k))+1)) & -1
}
