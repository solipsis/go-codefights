/**
Challenge: nextSquare
https://codefights.com/challenge/DCMvQpeqXo3EjPjG3

Given an integer n, return the smallest integer a, such that a is a square of some integer and is greater than n.

Example

For n = 5, the output should be
nextSquare(n) = 9.

Input/Output

[time limit] 4000ms (go)
[input] integer n

Guaranteed constraints:
1 ≤ n < 231.

[output] integer

The smallest square that is greater than n.
**/
import "math"

func nextSquare(n int) int {
    for {
        n++
        rt := int(math.Sqrt(float64(n)))
        if rt * rt == n { break }
    }
    return n
}
