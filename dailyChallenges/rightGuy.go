/**
Challenge: rightGuy
https://codefights.com/challenge/5W5Q5fNMiauqSgA6c

This is a reverse challenge. Enjoy!

Input/Output

[time limit] 4000ms (go)
[input] integer k

Constraints:
0 ≤ k < 231.

[output] integer
**/
func rightGuy(k int) (x int) {

    if k > 0 {
        x++
        for k & 1 < 1 {
            k >>= 1
            x++
        }
    }
    return
}
