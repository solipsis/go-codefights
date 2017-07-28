/**
Challenge: gold
https://codefights.com/challenge/ypaJW47zSBpJ3b7EJ


Return the number of ways to write the given n as the sum of two prime numbers.

Example
For n = 20, the output should be
gold(n) = 2.
There are two ways to represent 20 as the sum of two prime numbers:

3 + 17 = 20,
7 + 13 = 20.
Input/Output

[time limit] 4000ms (go)
[input] integer n

Guaranteed Constraints:
2 ≤ n ≤ 105.

[output] integer
**/
func gold(n int) int {
    // gen primes of to 10^5
    primes := []int{2,3}
    for x := 5; x <= n; x++ {
        if x % 2 == 0 { continue }
        if isPrime(x) {
            primes = append(primes, x)
        }
    }  
    // two pointers. 1 at start 1 at end
    l, r := 0, len(primes)-1
    ways := 0
    for l <= r {
        res := primes[l] + primes[r]
        if res == n {
            ways++
            l++      
        } else if res > n {
            r--
        } else {
            l++
        }
    }
    return ways
}

func isPrime(n int) bool {
    for i := 3; i*i <= n; i++ {
        if n % i == 0 {
            return false
        }
    }
    return true
}
