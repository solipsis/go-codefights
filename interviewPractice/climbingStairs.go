/*
Challenge: climbingStairs
https://codefights.com/interview-practice/task/oJXTWuwEZiC6FTw3A/description

You are climbing a staircase that has n steps. You can take the steps either 1 or 2 at a time. Calculate how many distinct ways you can climb to the top of the staircase.

Example

For n = 1, the output should be
climbingStairs(n) = 1;

For n = 2, the output should be
climbingStairs(n) = 2.

You can either climb 2 steps at once or climb 1 step two times.

Input/Output

[time limit] 4000ms (go)
[input] integer n

Guaranteed constraints:
1 ≤ n ≤ 50.

[output] integer

It's guaranteed that the answer will fit in a 32-bit integer.
*/
func climbingStairs(n int) int {    
    a, b := 1, 2
    for i := 2; i < n; i++ {
        a,b = b, b+a
    }
    return b
}
