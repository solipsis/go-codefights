/**
Challenge: Add Two Digits
https://codefights.com/arcade/code-arcade/intro-gates/wAGdN6FMPkx7WBq66/description


You are given a two-digit integer n. Return the sum of its digits.

Example

For n = 29, the output should be
addTwoDigits(n) = 11.

Input/Output

[time limit] 4000ms (go)
[input] integer n

A positive two-digit integer.

Guaranteed constraints:
10 ≤ n ≤ 99.

[output] integer

The sum of the first and second digits of the input number.
**/
func addTwoDigits(n int) int {
    return n/10 + n%10
}
