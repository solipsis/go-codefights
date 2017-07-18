/**
Challenge: canBuildRectangle
https://codefights.com/tournaments/ZS9S4h3HFKsuWRMnD/D

You have n sticks with the lengths 1, 2, 3, ..., n, respectively. You like to build different shapes with the sticks, and these sticks are very dear to you so of course you don't want to break them. You came up with an idea to build a rectangle from your sticks, so you want to check whether it's possible to build a rectangle with positive sides using all of your sticks. Given the number of sticks as an integer n, determine whether it's possible to do so without breaking any of them.

Example

For n = 12, the output should be
canBuildRectangle(n) = true.

For example, we can build the following rectangle:

The first pair of parallel sides with length 10 can be built from sticks 4, 6 and 10;
The second pair of parallel sides with length 29 can be built from sticks 1, 5, 11, 12 and 2, 3, 7, 8, 9.
For n = 13, the output should be
canBuildRectangle(n) = false.

Input/Output

[time limit] 4000ms (go)
[input] integer n

A positive integer.

Guaranteed constraints:
1 ≤ n ≤ 109.

[output] boolean

Return true if you can build a rectangle using sticks with lengths 1, 2, 3, ..., n, otherwise return false.
**/
func canBuildRectangle(n int) bool {
    
    x := (n*(n+1))/2
    return n >= 7 && x % 2 == 0
}
