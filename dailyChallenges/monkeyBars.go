/*
Challenge: monkeyBars
https://codefights.com/challenge/vQ3EJBWeGDHEPb8aK

You are at the International Monkey Bars Climbing Competition 2017! Each contestant must climb various stretches of monkey bars during each round. The rules are that you may not more move than 3 bars at a time. Find the number of distinct ways to climb a given section.

Example
For n = 4, the output should be
monkeyBars(n) = 7.
There are 7 distinct ways to climb the bars - 3 + 1, 1 + 3, 2 + 1 + 1, 1 + 2 + 1, 1 + 1 + 2, 1 + 1 + 1 + 1, 2 + 2.

Input/Output

[time limit] 4000ms (go)
[input] integer n

Number of bars.

Guaranteed constraints:
1 ≤ n ≤ 60.

[output] integer64

Number of distinct ways to climb.
*/
func monkeyBars(n int) int64 {
    
    memo := make([]int64, 61)
    memo[1] = 1;
    memo[2] = 2;
    memo[3] = 4;
    return helper(n, memo)
}

func helper(n int, memo []int64) int64 {
    if memo[n] > 0 {
        return memo[n]
    }
    
    v := helper(n-1, memo) + helper(n-2, memo) + helper(n-3, memo)
    memo[n] = v
    return v
}

// Super golfed solution
var m =[61]int64 {0, 1,2,4}
func monkeyBars(n int) int64 {
    f := monkeyBars
   
    if m[n] > 0 {
        return m[n]
    }
    m[n] = f(n-1) + f(n-2) + f(n-3)
    
    return m[n]
}
