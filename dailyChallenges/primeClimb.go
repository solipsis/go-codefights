/**
Challenge: primeClimb
https://codefights.com/challenge/exviHuziJjMp9hk4v

Recently, one of John Conway's $1000 problems proved to be false.

The problem is called Climb to a Prime. Here's how it works: Suppose we have a positive integer n. Find its prime factorization, omitting any exponents of 1, and sort the prime factors in ascending order. Then we "bring down" the exponents, remove the multiplication sign, and form a new number. For example, suppose we have n = 60. Factorize 60 = 22 × 3 × 5, then we "bring down" the exponent 2, remove the multiplication sign, and obtain the result 2235.

This process can then be repeated on the new number, and so on. Usually, this process will make the number bigger and bigger, so repeatedly applying this process can be thought of as "climbing" up the stairs. If we hit a prime number, we'll be stuck at if forever. Continuing the example, 2235 = 3 × 5 × 149, so the next number is 35149, which is a prime, and the "climbing" process is stopped. The "Climb to a Prime" problem conjectures that, for any initial number, this process will always "climb" up to a prime and get stuck there.

Proving this conjecture to be true of false is not an easy task, as there are some initial numbers that we currently don't know whether they will get stuck or not. One such example is the initial value 20: The sequence begins with 20 → 225 → 3252 → 223271 → ……, and currently contains over 100 known non-repeated terms, all of which are composite.

But just a few days ago, a person named James Davis found this number: 13532385396179. If we factorize this number, we get 13532385396179 = 13 × 532 × 3853 × 96179. When we bring down the exponent and remove multiplication sign, we get the number itself. The fact that we are stuck at this composite number proves the conjecture false.

As you see, "climbing" all the way up will sometimes lead to long or even unresolved chains. So in this challenge, you are only asked to do one step of "climbing". That is, given the number n, do the procedure described above one time, and return the result as a string.

Example

For n = 60, the answer should be
primeClimb(n) = "2235".

For n = 20, the answer should be
primeClimb(n) = "225".

Input/Output

[time limit] 4000ms (go)
[input] integer n

Guaranteed constraints:
2 ≤ n < 231.

[output] string

The number after one step of "climbing", returned as a string.
**/
import "sort"

func primeClimb(n int) (s string) {
    
    m := make(map[int]int)
    z := make([]int, 0)
    for i := 2; i * i <= n; i++ {
        
        if n % i == 0 {
            z = append(z, i)
        }
        
        for n % i == 0 {
            m[i]++
            n /= i
        }
    }
    if n > 1 {
        m[n]++
        z = append(z, n)
    }
    
    sort.Ints(z)
    for _, v := range z {
        s += strconv.Itoa(v)
        if m[v] > 1 {
            s += strconv.Itoa(m[v])
        }
    }
    
    return 
}
