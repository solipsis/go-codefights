/**
Challenge: imagine 
https://codefights.com/challenge/vwrjaSzc2egE3FiYf

This is a reverse challenge, enjoy!

Input/Output

[time limit] 4000ms (go)
[input] integer a

Guaranteed constraints:
-1000 ≤ a ≤ 1000.

[input] integer b

Guaranteed constraints:
-1000 ≤ b ≤ 1000.

[input] integer c

Guaranteed constraints:
-1000 ≤ c ≤ 1000.

[input] integer d

Guaranteed constraints:
-1000 ≤ d ≤ 1000.

[output] array.integer

An array of length 2.


**/
func imagine(a, b, c, d int) []int {

    return []int{a*c-b*d, b*c+a*d}
}
