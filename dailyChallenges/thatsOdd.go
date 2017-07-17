/**
Challenge: thatsOdd
https://codefights.com/challenge/fZDobaa2BQncF7NXA

This is a reverse challenge enjoy!

Input/Output

[time limit] 4000ms (go)
[input] string str

Guaranteed constraints:
0 ≤ str.length < 1000

[output] boolean
**/
func thatsOdd(s string) bool {
    m := 0
    for i := 0; i < len(s); i++ {
        m = m + s[i] % 2
    }
    return m > 0
}
