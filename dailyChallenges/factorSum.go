/**
Challenge: factorSum
https://codefights.com/challenge/7WDDw8v7QKHaHN6bK

Consider the following operation - we take a positive integer n and replace it with the sum of its prime factors (if a prime number is presented multiple times in the factorization of n, then it's counted the same number of times in the sum). This operation is applied sequentially first to the given number, then to the first result, then to the second result and so on, until the result remains the same.

Given any number, find the final result of the operation.

Example

For n = 24, the output should be
factorSum(n) = 5.

24 -> (2 + 2 + 2 + 3) = 9 -> (3 + 3) = 6 -> (2 + 3) = 5 -> 5.
So the answer for n = 24 is 5.

Input/Output

[time limit] 4000ms (go)
[input] integer n

Guaranteed constraints:
2 ≤ n ≤ 200.

[output] integer
**/
func factorSum(n int) int {
    
    // keep summing the prime factors until the number doesn't change for 1 cycle
    p := -1
    for p != n {
        p = n
        s := 0
        for i := 2; i * i <= n; i++ {
            for n % i == 0 {
                n /= i
                s += i
            }
        }
        if n > 1 {
            s += n
        }
        n = s
    }
    return p
}
