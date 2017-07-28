/**
Challenge: raiseItMinLength
https://codefights.com/challenge/7LzAodAedJ46pqv7W

There was a recent reverse challenge called raiseIt, the algorithm of which takes an array of two digit numbers and finds the sum of the first digit raised to the power of the second digit. For example raiseIt([20, 42]) = 20 + 42 = 1 + 16 = 17.

In that task, I was curious what the shortest array would be that generated the largest possible output. In this new task, now I'm curious about what the length of the shortest array for any given number is.

Given an integer n, find the length of the shortest array arr such that raiseIt(arr) = n.

Example
For n = 24, the output should be
raiseItMinLength(n) = 2.
One of the possible shortest arrays is [23, 42]. 23 + 42 = 8 + 16 = 24.

[time limit] 4000ms (go)
[input] integer n

Guaranteed constraints:
0 ≤ n ≤ 105.

[output] integer

The length of the shortest array arr such that raiseIt(arr) = n.
**/
func raiseItMinLength(n int) int {

    if n == 0 {
        return 0
    }
    
    // very ugly :D
    m := make(map[int]bool)
    v := make([]int, 0)
    for i := 0; i < 9; i++ {
        x := 1
        for j := 0; j < 9; j++ {
            x *= i
            if (x <= 100000) {
                v = append(v, x)
                m[x] = 1 > 0
            }
        }
    }
    if m[n] {
        return 1
    }
    for i := 2; i < 5; i++ {
        t := make([]int, 0)
        for _, x := range v {
            for z := range m {
                t = append(t, z+x)
            }           
        }
        for _, j := range t {
            m[j] = 1 > 0
        }       
        if m[n] {
            return i
        }
    }
    return 5
}
