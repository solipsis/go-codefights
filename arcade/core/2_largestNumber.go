/**
Challenge: Largest Number
https://codefights.com/arcade/code-arcade/intro-gates/SZB5XypsMokGusDhX/description

Given an integer n, return the largest number that contains exactly n digits.

Example

For n = 2, the output should be
largestNumber(n) = 99.

Input/Output

[time limit] 4000ms (go)
[input] integer n

Guaranteed constraints:
1 ≤ n ≤ 7.

[output] integer

The largest integer of length n.
**/
func largestNumber(n int) int {
    s := 0
    for n > 0 {
        s = s*10 + 9
        n--
    }
    return s
}
